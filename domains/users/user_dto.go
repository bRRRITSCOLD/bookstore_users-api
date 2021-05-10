package users

import (
	datesUtils "bookstore_users-api/utils/dates"
	errorsUtils "bookstore_users-api/utils/errors"
	"fmt"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) GetByUserID() *errorsUtils.APIError {
	result := usersDB[user.UserID]
	if result == nil {
		return errorsUtils.NewNotFoundAPIError(
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

func (user *User) Save() *errorsUtils.APIError {
	foundUser := usersDB[user.UserID]
	if foundUser != nil {
		if foundUser.Email == user.Email {
			return errorsUtils.NewNotFoundAPIError(
				fmt.Sprintf("email %s already registered", user.Email),
			)
		}
		return errorsUtils.NewNotFoundAPIError(
			fmt.Sprintf("user %d already exists", user.UserID),
		)
	}

	user.DateCreated = datesUtils.GetNowString()

	usersDB[user.UserID] = user

	return nil
}
