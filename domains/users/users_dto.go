package users_domain

import (
	errorsUtils "bookstore_users-api/utils/errors"
	"time"
)

type User struct {
	UserID      int64     `json:"userId"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"dateCreated"`
}

func NewUser(user *User) *User {
	return nil
}

func (user *User) Validate() *errorsUtils.APIError {
	if user.Email == "" {
		return errorsUtils.NewBadRequestAPIError("invalid email address")
	}
	return nil
}
