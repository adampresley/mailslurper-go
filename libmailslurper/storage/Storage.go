// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package storage

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/adampresley/golangdb"
	"github.com/adampresley/mailslurper/libmailslurper/model/attachment"
	"github.com/adampresley/mailslurper/libmailslurper/model/mailitem"

	"github.com/nu7hatch/gouuid"
)

/*
Creates a global connection handle in a map named "lib".
*/
func ConnectToStorage(connectionInfo *golangdb.DatabaseConnection) error {
	var err error

	err = connectionInfo.Connect("lib")
	if err != nil {
		return err
	}

	switch connectionInfo.Engine {
	case golangdb.SQLITE:
		CreateSqlliteDatabase()
	}

	return nil
}

/*
Disconnects from the database storage
*/
func DisconnectFromStorage() {
	golangdb.Db["lib"].Close()
}

/*
Generate a UUID ID for database records.
*/
func GenerateId() string {
	id, _ := uuid.NewV4()
	return id.String()
}

/*
Returns an attachment by ID
*/
func GetAttachment(id string) (attachment.Attachment, error) {
	result := attachment.Attachment{}

	rows, err := golangdb.Db["lib"].Query(`
		SELECT
			  fileName TEXT
			, contentType TEXT
			, content TEXT
		FROM attachment
		WHERE
			id=?
	`, id)

	if err != nil {
		return result, fmt.Errorf("Error running query to get attachment")
	}

	defer rows.Close()
	rows.Next()

	var fileName string
	var contentType string
	var content string

	rows.Scan(&fileName, &contentType, &content)

	result.Headers = &attachment.AttachmentHeader{
		FileName: fileName,
		ContentType: contentType,
	}

	result.Contents = content
	return result, nil
}

/*
Retrieves all stored mail items as an array of MailItem items.
Takes an ID as a filter. If ID == "" then all records are returned.
*/
func GetMails(id string) ([]mailitem.MailItem, error) {
	result := make([]mailitem.MailItem, 0)
	attachments := make([]*attachment.Attachment, 0)

	sql := `
		SELECT
			  mailitem.id AS mailItemId
			, mailitem.dateSent
			, mailitem.fromAddress
			, mailitem.toAddressList
			, mailitem.subject
			, mailitem.xmailer
			, mailitem.body
			, mailitem.contentType
			, mailitem.boundary
			, attachment.id AS attachmentId
			, attachment.fileName

		FROM mailitem
			LEFT OUTER JOIN attachment ON mailitem.id=attachment.mailItemId

		WHERE 1=1`

	if id != "" {
		sql = sql + " AND mailitem.id=? "
	}

	sql = sql + `ORDER BY mailitem.dateSent DESC`

	rows, err := golangdb.Db["lib"].Query(sql, id)

	if err != nil {
		return result, fmt.Errorf("Error running query to get mail items: %s", err)
	}

	/*
	 * Loop over the result, grouping by mail item ID, and add
	 * to the resulting array. There will be multiple mail items
	 * because of the join to attachments.
	 */
	currentMailItemId := ""
	newItemCreated := false
	newItem := mailitem.MailItem{}

	for rows.Next() {
		var mailItemId string
		var dateSent string
		var fromAddress string
		var toAddressList string
		var subject string
		var xmailer string
		var body string
		var contentType string
		var boundary string
		var attachmentId string
		var fileName string

		rows.Scan(&mailItemId, &dateSent, &fromAddress, &toAddressList, &subject, &xmailer, &body, &contentType, &boundary, &attachmentId, &fileName)

		/*
		 * If this is our first iteration then we haven't looked at a
		 * current item yet
		 */
		if currentMailItemId == "" {
			currentMailItemId = mailItemId
		}

		/*
		 * There will be multiple records per mail item if there are
		 * multiple attachments. As such make sure we are getting all
		 * the IDs first, and the mail item only once.
		 */
		newItem = mailitem.MailItem{}

		if currentMailItemId != "" && currentMailItemId == mailItemId {
			if attachmentId != "" {
				attachments = append(attachments, &attachment.Attachment{Id: attachmentId, Headers: &attachment.AttachmentHeader{FileName: fileName}})
			}

			if !newItemCreated {
				newItemCreated = true

				newItem = mailitem.MailItem{
					Id:              mailItemId,
					DateSent:        dateSent,
					FromAddress:     fromAddress,
					ToAddresses:     strings.Split(toAddressList, "; "),
					Subject:         subject,
					XMailer:         xmailer,
					Body:            body,
					ContentType:     contentType,
					Boundary:        boundary,
					Attachments:     nil,
				}
			}
		} else {
			newItem.Attachments = attachments
			log.Printf("Retrieving mail item %d from %s with a subject of %s", mailItemId, fromAddress, subject)

			result = append(result, newItem)
			attachments = make([]*attachment.Attachment, 0)

			if attachmentId != "" {
				attachments = append(attachments, &attachment.Attachment{Id: attachmentId, Headers: &attachment.AttachmentHeader{FileName: fileName}})
			}

			newItemCreated = true

			newItem = mailitem.MailItem{
				Id:              mailItemId,
				DateSent:        dateSent,
				FromAddress:     fromAddress,
				ToAddresses:     strings.Split(toAddressList, "; "),
				Subject:         subject,
				XMailer:         xmailer,
				Attachments:     nil,
			}

			if currentMailItemId != mailItemId {
				currentMailItemId = mailItemId
			}
		}
	}

	newItem.Attachments = attachments
	result = append(result, newItem)

	rows.Close()
	return result, nil
}

func storeAttachments(mailItemId string, transaction *sql.Tx, attachments []*attachment.Attachment) error {
	for _, a := range attachments {
		attachmentId := GenerateId()

		statement, err := transaction.Prepare(`
			INSERT INTO attachment (
				  id
				, mailItemId
				, fileName
				, contentType
				, content
			) VALUES (
				  ?
				, ?
				, ?
				, ?
				, ?
			)
		`)

		if err != nil {
			return fmt.Errorf("Error preparing insert attachment statement: %s", err)
		}

		_, err = statement.Exec(
			attachmentId,
			mailItemId,
			a.Headers.FileName,
			a.Headers.ContentType,
			a.Contents,
		)

		if err != nil {
			return fmt.Errorf("Error executing insert attachment in StoreMail: %s", err)
		}

		statement.Close()
		a.Id = attachmentId
	}

	return nil
}

func StoreMail(mailItem *mailitem.MailItem) (string, error) {
		/*
		 * Create a transaction and insert the new mail item
		 */
		transaction, err := golangdb.Db["lib"].Begin()
		if err != nil {
			return "", fmt.Errorf("Error starting transaction in StoreMail: %s", err)
		}

		/*
		 * Insert the mail item
		 */
		statement, err := transaction.Prepare(`
			INSERT INTO mailitem (
				  id
				, dateSent
				, fromAddress
				, toAddressList
				, subject
				, xmailer
				, body
				, contentType
				, boundary
			) VALUES (
				  ?
				, ?
				, ?
				, ?
				, ?
				, ?
				, ?
				, ?
				, ?
			)
		`)

		if err != nil {
			return "", fmt.Errorf("Error preparing insert statement for mail item in StoreMail: %s", err)
		}

		mailItemId := GenerateId()

		_, err = statement.Exec(
			mailItemId,
			mailItem.DateSent,
			mailItem.FromAddress,
			strings.Join(mailItem.ToAddresses, "; "),
			mailItem.Subject,
			mailItem.XMailer,
			mailItem.Body,
			mailItem.ContentType,
			mailItem.Boundary,
		)

		if err != nil {
			transaction.Rollback()
			return "", fmt.Errorf("Error executing insert for mail item in StoreMail: %s", err)
		}

		statement.Close()
		mailItem.Id = mailItemId

		/*
		 * Insert attachments
		 */
		if err = storeAttachments(mailItemId, transaction, mailItem.Attachments); err != nil {
			transaction.Rollback()
			return "", fmt.Errorf("Unable to insert attachments in StoreMail: %s", err)
		}

		transaction.Commit()
		log.Printf("New mail item written to database.\n\n")

		return mailItemId, nil
}
