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
	USERS_MYSQL_DB_EMAIL_UNIQUE      = "email_UNIQUE"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) GetByUserID() *errors_utils.APIError {
	pingErr := users_mysql_db.Client.Ping()
	errors_utils.PanicOnError(pingErr)

	result := usersDB[user.UserID]
	if result == nil {
		return errors_utils.NewNotFoundAPIError(
			fmt.Sprintf("user %d not found", user.UserID),
		)
	}

	user.UserID = result.UserID
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors_utils.APIError {
	stmt, prepareErr := users_mysql_db.Client.Prepare(USERS_MYSQL_DB_INSERT_USER_QUERY)
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
