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
	"github.com/adampresley/mailslurper/libmailslurper/mode/attachment"
	"github.com/adampresley/mailslurper/libmailslurper/mode/mailitem"
)

/*
Creates a global connection handle in a map named "lib".
*/
func ConnectToStorage(connectionInfo *golangdb.DatabaseConnection) error {
	return connection.Connect("lib")
}

/*
Disconnects from the database storage
*/
func DisconnectFromStorage() {
	golangdb.Db["lib"].Close()
}

/*
Returns an attachment by ID
*/
func GetAttachment(id int) (attachment.Attachment, error) {
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

	result.Headers = attachment.AttachmentHeader{
		FileName: fileName,
		ContentType: contentType,
	}

	result.Contents = content
	return result
}

/*
Retrieves all stored mail items as an array of MailItem items.
Takes an ID as a filter. If ID == 0 then all records are returned.
*/
func GetMails(id int) ([]mailitem.MailItem, error) {
	result := make([]mailitem.MailItem, 0)
	attachments := make([]attachment.Attachment, 0)

	sql := `
		SELECT
			  mailitem.id AS mailItemId
			, mailitem.dateSent
			, mailitem.fromAddress
			, mailitem.toAddressList
			, mailitem.subject
			, mailitem.xmailer
			, attachment.id AS attachmentId
			, attachment.fileName

		FROM mailitem
			LEFT OUTER JOIN attachment ON mailitem.id=attachment.mailItemId

		WHERE 1=1
	`

	if id > 0 {
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
	currentMailItemId := 0
	newItemCreated := false

	for rows.Next() {
		var mailItemId int
		var dateSent string
		var fromAddress string
		var toAddressList string
		var subject string
		var xmailer string
		var attachmentId int
		var fileName string

		rows.Scan(&mailItemId, &dateSent, &fromAddress, &toAddressList, &subject, &xmailer, &attachmentId, &fileName)

		/*
		 * If this is our first iteration then we haven't looked at a
		 * current item yet
		 */
		if currentMailItemId == 0 {
			currentMailItemId = mailItemId
		}

		/*
		 * There will be multiple records per mail item if there are
		 * multiple attachments. As such make sure we are getting all
		 * the IDs first, and the mail item only once.
		 */
		if currentMailItemId > 0 && currentMailItemId == mailItemId {
			if attachmentId > 0 {
				attachments = append(attachments, attachment.Attachment{Id: attachmentId, Headers: attachment.AttachmentHeader{FileName: fileName}})
			}

			if !newItemCreated {
				newItemCreated = true

				newItem := mailitem.MailItem{
					Id:              mailItemId,
					DateSent:        dateSent,
					FromAddress:     fromAddress,
					ToAddresses:     strings.Split(toAddressList, "; "),
					Subject:         subject,
					XMailer:         xmailer,
					Attachments:     nil,
				}
			}
		} else {
			newItem.Attachments = attachments
			log.Printf("Retrieving mail item %d from %s with a subject of %s", mailItemId, fromAddress, subject)

			result = append(result, newItem)
			attachments = make([]attachment.Attachment, 0)

			if attachmentId > 0 {
				attachments = append(attachments, attachment.Attachment{Id: attachmentId, Headers: attachment.AttachmentHeader{FileName: fileName}})
			}

			newItemCreated = true

			newItem := mailitem.MailItem{
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
	return result
}

func storeAttachments(transaction *sql.Tx, attachments []*attachment.Attachment) error {
	for _, a := range attachments {
		statement, err = transaction.Prepare(`
			INSERT INTO attachment (
				  mailItemId
				, fileName
				, contentType
				, content
			) VALUES (
				  ?
				, ?
				, ?
				, ?
			)`
		)

		if err != nil {
			return fmt.Errorf("Error preparing insert attachment statement: %s", err)
		}

		_, err = statement.Exec(
			mailItemId,
			a.Headers.FileName,
			a.Headers.ContentType,
			a.Contents,
		)

		if err != nil {
			return fmt.Errorf("Error executing insert attachment in StoreMail: %s", err)
		}

		statement.Close()
	}

	return nil
}

func StoreMail(mailItem *mailitem.MailItem) (int, error) {
		/*
		 * Create a transaction and insert the new mail item
		 */
		transaction, err := golangdb.Db["lib"].Begin()
		if err != nil {
			return 0, fmt.Errorf("Error starting transaction in StoreMail: %s", err)
		}

		/*
		 * Insert the mail item
		 */
		statement, err := transaction.Prepare(`
			INSERT INTO mailitem (
				  dateSent
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
			)`
		)

		if err != nil {
			return 0, fmt.Errorf("Error preparing insert statement for mail item in StoreMail: %s", err)
		}

		result, err := statement.Exec(
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
			return 0, fmt.Errorf("Error executing insert for mail item in StoreMail: %s", err)
		}

		statement.Close()
		mailItemId, _ := result.LastInsertId()
		mailItem.Id = int(mailItemId)

		/*
		 * Insert attachments
		 */
		err = storeAttachments(transaction, mailItem.Attachments); err != nil {
			transaction.Rollback()
			return 0, fmt.Errorf("Unable to insert attachments in StoreMail: %s", err)
		}

		transaction.Commit()
		log.Printf("New mail item written to database.\n\n")

		return mailItemId, nil
}
