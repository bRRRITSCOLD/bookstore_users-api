package users

import (
	users_domain "bookstore_users-api/domains/users"
	errors_utils "bookstore_users-api/utils/errors"
)

func CreateUser(user users_domain.User) (*users_domain.User, *errors_utils.APIError) {
	if validateUserErr := user.Validate(); validateUserErr != nil {
		return nil, validateUserErr
	}

	if saveUserErr := user.Save(); saveUserErr != nil {
		return nil, saveUserErr
	}

	return &user, nil
}

func GetUser(userId int64) (*users_domain.User, *errors_utils.APIError) {
	user := users_domain.User{UserID: userId}
	if getUserErr := user.GetByUserID(); getUserErr != nil {
		return nil, getUserErr
	}
	return &user, nil
}

func FindUser() {

}
