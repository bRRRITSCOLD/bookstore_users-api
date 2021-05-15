package errors_utils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	NO_ROWS = "sql: no rows in result set"
)

func ParseMySQLError(err error) *APIError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), NO_ROWS) {
			return NewNotFoundAPIError("mysql: no rows found")
		}

		return NewInternalServerAPIError("mysql: error identifying mysql")
	}

	switch sqlErr.Number {
	case 1062:
		return NewBadRequestAPIError("mysql: invalid data - unique index")
	}

	return NewInternalServerAPIError(fmt.Sprintf("error parsing mysql error: %s", sqlErr.Error()))
}
