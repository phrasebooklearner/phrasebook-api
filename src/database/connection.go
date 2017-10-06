package database

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBConnection(credits string) *sql.DB {
	driver, dsn := splitCredits(credits)

	connection, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	return connection
}

func splitCredits(credits string) (driver string, dsn string) {
	split := strings.Split(credits, "://")
	if len(split) == 2 {
		driver = split[0]
		dsn = split[1]
	}

	return driver, dsn
}
