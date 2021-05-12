package users_domain

import (
	users_mysql_db "bookstore_users-api/datasources/users/mysql"
	dates_utils "bookstore_users-api/utils/dates"
	errors_utils "bookstore_users-api/utils/errors"
	"fmt"
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
	foundUser := usersDB[user.UserID]
	if foundUser != nil {
		if foundUser.Email == user.Email {
			return errors_utils.NewNotFoundAPIError(
				fmt.Sprintf("email %s already registered", user.Email),
			)
		}
		return errors_utils.NewNotFoundAPIError(
			fmt.Sprintf("user %d already exists", user.UserID),
		)
	}

	user.DateCreated = dates_utils.GetNowString()

	usersDB[user.UserID] = user

	return nil
}
