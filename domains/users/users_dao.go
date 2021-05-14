package users_domain

import (
	users_mysql_db "bookstore_users-api/datasources/users/mysql"
	dates_utils "bookstore_users-api/utils/dates"
	errors_utils "bookstore_users-api/utils/errors"
	"fmt"
	"strings"
)

const (
	USERS_MYSQL_DB_INSERT_USER_QUERY = "INSERT INTO users(firstName, lastName, email, dateCreated) VALUES(?, ?, ?, ?);"
	USERS_MYSQL_DB_SELECT_USER_QUERY = "SELECT * from users WHERE id=?;"
	USERS_MYSQL_DB_EMAIL_UNIQUE      = "email_UNIQUE"
	USERS_MYSQL_DB_NO_ROWS           = "sql: no rows in result set"
)

func (user *User) GetByUserID() *errors_utils.APIError {
	stmt, prepareErr := users_mysql_db.Client.Preparex(USERS_MYSQL_DB_SELECT_USER_QUERY)
	if prepareErr != nil {
		return errors_utils.NewInternalServerAPIError(prepareErr.Error())
	}
	var foundUser User
	queryRowResult := stmt.QueryRowx(user.UserID)
	if scanStructErr := queryRowResult.StructScan(&foundUser); scanStructErr != nil {
		if strings.Contains(scanStructErr.Error(), USERS_MYSQL_DB_NO_ROWS) {
			return errors_utils.NewNotFoundAPIError(fmt.Sprintf("user %d not found", user.UserID))
		}
		return errors_utils.NewInternalServerAPIError(scanStructErr.Error())
	}

	user.UserID = foundUser.UserID
	user.Email = foundUser.Email
	user.FirstName = foundUser.FirstName
	user.LastName = foundUser.LastName
	user.DateCreated = foundUser.DateCreated

	return nil
}

func (user *User) Save() *errors_utils.APIError {
	stmt, prepareErr := users_mysql_db.Client.Preparex(USERS_MYSQL_DB_INSERT_USER_QUERY)
	if prepareErr != nil {
		return errors_utils.NewInternalServerAPIError(prepareErr.Error())
	}

	defer stmt.Close()

	user.DateCreated = dates_utils.GetNow()

	insertResult, insertErr := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
	)
	if insertErr != nil {
		if strings.Contains(insertErr.Error(), USERS_MYSQL_DB_EMAIL_UNIQUE) {
			return errors_utils.NewBadRequestAPIError(fmt.Sprintf("email %s already exists", user.Email))
		}

		return errors_utils.NewInternalServerAPIError(insertErr.Error())
	}

	userId, lastInsertedIDErr := insertResult.LastInsertId()
	if lastInsertedIDErr != nil {
		return errors_utils.NewInternalServerAPIError(lastInsertedIDErr.Error())
	}

	user.UserID = userId

	return nil
}
