// Copyright 2013 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package data

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

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
		create table mailitem (
			dateSent text,
			fromAddress text,
			toAddressList text,
			subject text,
			xmailer text,
			body text,
			contentType text,
			boundary text
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

		statement, err := transaction.Prepare("insert into mailitem (dateSent, fromAddress, toAddressList, subject, xmailer, body, contentType, boundary) values (?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(fmt.Sprintf("Erorr preparing insert statement: %s", err))
		}

		defer statement.Close()

		_, err = statement.Exec(
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

		transaction.Commit()
		fmt.Printf("New mail item written to database.\n\n")
	}
}

/*
Retrieves all stored mail items as an array of MailItemStruct items.
*/
func (ms *MailStorage) GetMails() []MailItemStruct {
	result := make([]MailItemStruct, 0)

	rows, err := ms.Db.Query("SELECT dateSent, fromAddress, toAddressList, subject, xmailer, body, contentType, boundary FROM mailitem ORDER BY dateSent DESC")
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
		var boundary string

		rows.Scan(&dateSent, &fromAddress, &toAddressList, &subject, &xmailer, &body, &contentType, &boundary)

		newItem := MailItemStruct{
			DateSent:    dateSent,
			FromAddress: fromAddress,
			ToAddresses: strings.Split(toAddressList, "; "),
			Subject:     subject,
			XMailer:     xmailer,
			Body:        body,
			ContentType: contentType,
			Boundary:    boundary,
		}

		result = append(result, newItem)
	}

	return result
}
