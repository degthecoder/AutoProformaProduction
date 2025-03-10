package app

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb" // Import the MSSQL driver
)

var Db *sql.DB

func ConnectDb() {
	conStr := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		Settings.DbUser, Settings.DbPassword, Settings.DbHost, Settings.DbPort, Settings.DbName)
	db, err := sql.Open("sqlserver", conStr)

	if err != nil {
		panic("Could not connect to database.")
	}

	db.SetMaxIdleConns(0)
	Db = db
}

func DisconnectDb() {
	Db.Close()
}
