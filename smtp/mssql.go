// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package smtp

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func ConnectMSSQL(host string, port string, database string, userName string, password string) (*sql.DB, error) {
	/*
	 * Create the connection
	 */
	log.Println("Connecting to MSSQL database")

	db, err := sql.Open("mssql", fmt.Sprintf("Server=%s;Port=%s;User Id=%s;Password=%s;Database=%s", host, port, userName, password, database))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateMSSQLDatabase(db *sql.DB) error {
	log.Println("Creating tables...")

	var err error

	sql := `
		IF OBJECT_ID('mailitem', 'U') IS NOT NULL BEGIN
			CREATE TABLE mailitem (
				id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
				dateSent DATETIME,
				fromAddress VARCHAR(255),
				toAddressList TEXT,
				subject VARCHAR(512),
				xmailer VARCHAR(50),
				body TEXT,
				contentType VARCHAR(50),
				boundary VARCHAR(50)
			);
		END
	`

	_, err = db.Exec(sql)
	if err != nil {
		return err
	}

	sql = `
		IF OBJECT_ID('attachment', 'U') IS NOT NULL BEGIN
			CREATE TABLE attachment (
				id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
				mailItemId INT,
				fileName VARCHAR(255),
				contentType VARCHAR(50),
				content TEXT
			);
		END
	`

	_, err = db.Exec(sql)
	if err != nil {
		return err
	}

	log.Println("Created tables successfully.")
	return nil
}
