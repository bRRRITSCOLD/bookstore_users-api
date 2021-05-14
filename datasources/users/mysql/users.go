package users_mysql_db

import (
	errors_utils "bookstore_users-api/utils/errors"
	"os"
	"strings"

	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERS_MYSQL_DB_USERNAME = "USERS_MYSQL_DB_USERNAME"
	USERS_MYSQL_DB_PASSWORD = "USERS_MYSQL_DB_PASSWORD"
	USERS_MYSQL_DB_HOST     = "USERS_MYSQL_DB_HOST"
	USERS_MYSQL_DB_DATABASE = "USERS_MYSQL_DB_DATABASE"
)

var (
	Client *sqlx.DB

	username = os.Getenv(USERS_MYSQL_DB_USERNAME)
	password = os.Getenv(USERS_MYSQL_DB_PASSWORD)
	host     = os.Getenv(USERS_MYSQL_DB_HOST)
	database = os.Getenv(USERS_MYSQL_DB_DATABASE)
)

func init() {
	datasourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		username,
		password,
		host,
		database,
	)

	var openErr error
	Client, openErr = sqlx.Open("mysql", datasourceName)
	errors_utils.PanicOnError(openErr)

	pingErr := Client.Ping()
	errors_utils.PanicOnError(pingErr)

	Client.Mapper = reflectx.NewMapperFunc("mysql", strings.ToLower)

	// mysql.SetLogger()
	log.Println("database successfully configured")
}
