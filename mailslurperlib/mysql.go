// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package mailslurperlib

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(host string, port string, database string, userName string, password string) (*sql.DB, error) {
	/*
	 * Create the connection
	 */
	log.Println("Connecting to MySQL database")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?autocommit=true", userName, password, host, port, database))
	if err != nil {
		return nil, err
	}

	temp := os.Getenv("max_connections")
	if temp == "" {
		temp = "151"
	}

	maxConnections, _ := strconv.Atoi(temp)

	db.SetMaxIdleConns(maxConnections)
	db.SetMaxOpenConns(maxConnections)

	return db, nil
}

func CreateMySQLDatabase(db *sql.DB) error {
	log.Println("Creating tables...")

	var err error

	sql := `
		CREATE TABLE IF NOT EXISTS mailitem (
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
	`

	_, err = db.Exec(sql)
	if err != nil {
		return err
	}

	sql = `
		CREATE TABLE IF NOT EXISTS attachment (
			id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
			mailItemId INT,
			fileName VARCHAR(255),
			contentType VARCHAR(50),
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
