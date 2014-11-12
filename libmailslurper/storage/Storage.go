package storage

import (
	"log"

	"github.com/adampresley/golangdb"
)

func ConnectToStorage(connectionInfo *golangdb.DatabaseConnection) error {
	return connection.Connect("main")
}

func DisconnectFromStorage() {
	golangdb.Db["main"].Close()
}

