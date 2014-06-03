// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package smtp

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectSqlite() (*sql.DB, error) {
	os.Remove("./mail.db")

	/*
	 * Create the connection
	 */
	log.Println("Connecting to SQLITE3 database 'mail.db'")

	db, err := sql.Open("sqlite3", "./mail.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateSqlliteDatabase(db *sql.DB) error {
	log.Println("Creating tables...")

	var err error

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

	_, err = db.Exec(sql)
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

	_, err = db.Exec(sql)
	if err != nil {
		return err
	}

	log.Println("Created tables successfully.")
	return nil
}
