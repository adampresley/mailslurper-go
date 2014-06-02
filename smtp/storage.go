// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package smtp

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adampresley/mailslurper/admin/model"
	_ "github.com/mattn/go-sqlite3"
)

/*
Structure for holding a persistent database connection.
*/
type MailStorage struct {
	Db *sql.DB
}

// Global variable for our server's database connection
var Storage MailStorage

/*
Open a connection to a SQLite database. This will attempt to delete any
existing database file and create a new one with a blank table for holding
mail data.
*/
func (ms *MailStorage) Connect(filename string) error {
	os.Remove(filename)

	/*
	 * Create the connection
	 */
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return err
	}

	ms.Db = db

	/*
	 * Create the mailitem table.
	 */
	sql := `
		CREATE TABLE mailitem (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			dateSent TEXT,
			fromAddress TEXT,
			toAddressList TEXT,
			subject TEXT,
			xmailer TEXT,
			body TEXT,
			contentType TEXT,
			boundary TEXT
		);
	`

	_, err = ms.Db.Exec(sql)
	if err != nil {
		return err
	}

	sql = `
		CREATE TABLE attachment (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			mailItemId INTEGER,
			fileName TEXT,
			contentType TEXT,
			content TEXT
		);
	`

	_, err = ms.Db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

/*
Close a SQLite database connection.
*/
func (ms *MailStorage) Disconnect() {
	ms.Db.Close()
}

/*
Listens for messages on a channel for mail messages to be written
to disk. This channel takes in MailItemStruct mail items.
*/
func (ms *MailStorage) StartWriteListener(dbWriteChannel chan MailItemStruct) {
	for {
		mailItem := <-dbWriteChannel

		/*
		 * Create a transaction and insert the new mail item
		 */
		transaction, err := ms.Db.Begin()
		if err != nil {
			panic(fmt.Sprintf("Error starting insert transaction: %s", err))
		}

		/*
		 * Insert the mail item
		 */
		statement, err := transaction.Prepare("INSERT INTO mailitem (dateSent, fromAddress, toAddressList, subject, xmailer, body, contentType, boundary) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(fmt.Sprintf("Error preparing insert statement: %s", err))
		}

		//defer statement.Close()

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
			panic(fmt.Sprintf("Error executing insert statement: %s", err))
		}

		statement.Close()
		mailItemId, _ := result.LastInsertId()

		/*
		 * Insert attachments
		 */
		for index := 0; index < len(mailItem.Attachments); index++ {
			statement, err = transaction.Prepare("INSERT INTO attachment (mailItemId, fileName, contentType, content) VALUES (?, ?, ?, ?)")
			if err != nil {
				panic(fmt.Sprintf("Error preparing insert attachment statement: %s", err))
			}

			_, err = statement.Exec(
				mailItemId,
				mailItem.Attachments[index].Headers.FileName,
				mailItem.Attachments[index].Headers.ContentType,
				mailItem.Attachments[index].Contents,
			)

			if err != nil {
				panic(fmt.Sprintf("Error executing insert attachment statement: %s", err))
			}

			statement.Close()
		}

		transaction.Commit()
		log.Printf("New mail item written to database.\n\n")
	}
}

/*
Retrieves all stored mail items as an array of MailItemStruct items.
*/
func (ms *MailStorage) GetMails() []model.JSONMailItem {
	result := make([]model.JSONMailItem, 0)

	rows, err := ms.Db.Query(`
		SELECT
			  mailitem.dateSent
			, mailitem.fromAddress
			, mailitem.toAddressList
			, mailitem.subject
			, mailitem.xmailer
			, mailitem.body
			, mailitem.contentType
			, COUNT(attachment.id) AS attachmentCount
		FROM mailitem
			LEFT OUTER JOIN attachment ON mailitem.id=attachment.mailItemId
		GROUP BY mailitem.id
		ORDER BY mailitem.dateSent DESC
	`)

	if err != nil {
		panic("Error running query to get mail items")
	}

	defer rows.Close()

	for rows.Next() {
		var dateSent string
		var fromAddress string
		var toAddressList string
		var subject string
		var xmailer string
		var body string
		var contentType string
		var attachmentCount int

		rows.Scan(&dateSent, &fromAddress, &toAddressList, &subject, &xmailer, &body, &contentType, &attachmentCount)

		newItem := model.JSONMailItem{
			DateSent:        dateSent,
			FromAddress:     fromAddress,
			ToAddresses:     strings.Split(toAddressList, "; "),
			Subject:         subject,
			XMailer:         xmailer,
			Body:            body,
			ContentType:     contentType,
			AttachmentCount: attachmentCount,
		}

		result = append(result, newItem)
	}

	return result
}
