package users_domain

import (
	errorsUtils "bookstore_users-api/utils/errors"
	"time"
)

const (
	USER_ACTIVE_STATUS   = "active"
	USER_INACTIVE_STATUS = "inactive"
)

type User struct {
	UserID      int64     `json:"userId" mysql:"id"`
	FirstName   string    `json:"firstName" mysql:"firstName"`
	LastName    string    `json:"lastName" mysql:"lastName"`
	Email       string    `json:"email" mysql:"email"`
	DateCreated time.Time `json:"dateCreated" mysql:"dateCreated"`
	Status      string    `json:"status" mysql:"status"`
	Password    string    `json:"password" mysql:"password"`
}

type Users []User

func NewUser(user *User) *User {
	return nil
}

func (user *User) Validate() *errorsUtils.APIError {
	if user.Email == "" {
		return errorsUtils.NewBadRequestAPIError("invalid email address")
	}

	if user.Password == "" {
		return errorsUtils.NewBadRequestAPIError("invalid password")
	}
	return nil
}
