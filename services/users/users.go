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

func PutUser(user users_domain.User) (*users_domain.User, *errors_utils.APIError) {
	if validateUserErr := user.Validate(); validateUserErr != nil {
		return nil, validateUserErr
	}

	currentUser, getUserErr := GetUser(user.UserID)
	if getUserErr != nil {
		return nil, getUserErr
	}

	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.Email = user.Email

	currentUser.PutByUserID()

	return currentUser, nil
}

func PatchUser(user users_domain.User) (*users_domain.User, *errors_utils.APIError) {
	currentUser, getUserErr := GetUser(user.UserID)
	if getUserErr != nil {
		return nil, getUserErr
	}

	if user.FirstName != "" {
		currentUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		currentUser.LastName = user.LastName
	}
	if user.Email != "" {
		currentUser.Email = user.Email
	}

	currentUser.PutByUserID()

	return currentUser, nil
}

func FindUser() {

}
