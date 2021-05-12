package users_mysql_db

import (
	errors_utils "bookstore_users-api/utils/errors"
	"os"

	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERS_MYSQL_DB_USERNAME = "USERS_MYSQL_DB_USERNAME"
	USERS_MYSQL_DB_PASSWORD = "USERS_MYSQL_DB_PASSWORD"
	USERS_MYSQL_DB_HOST     = "USERS_MYSQL_DB_HOST"
	USERS_MYSQL_DB_DATABASE = "USERS_MYSQL_DB_DATABASE"
)

var (
	Client *sql.DB

	username = os.Getenv(USERS_MYSQL_DB_USERNAME)
	password = os.Getenv(USERS_MYSQL_DB_PASSWORD)
	host     = os.Getenv(USERS_MYSQL_DB_HOST)
	database = os.Getenv(USERS_MYSQL_DB_DATABASE)
)

func init() {
	datasourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		database,
	)

	var openErr error
	Client, openErr = sql.Open("mysql", datasourceName)
	errors_utils.PanicOnError(openErr)

	pingErr := Client.Ping()
	errors_utils.PanicOnError(pingErr)

	// mysql.SetLogger()
	log.Println("database successfully configured")
}
