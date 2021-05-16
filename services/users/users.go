package users

import (
	users_domain "bookstore_users-api/domains/users"
	crypto_utils "bookstore_users-api/utils/crypto"
	dates_utils "bookstore_users-api/utils/dates"
	errors_utils "bookstore_users-api/utils/errors"
)

func CreateUser(user users_domain.User) (*users_domain.User, *errors_utils.APIError) {
	if validateUserErr := user.Validate(); validateUserErr != nil {
		return nil, validateUserErr
	}

	user.DateCreated = dates_utils.GetNow()
	user.Status = "active"
	user.Password = crypto_utils.MD5Hash(user.Password)

	if saveUserErr := user.Save(); saveUserErr != nil {
		return nil, saveUserErr
	}

	return &user, nil
}

func GetUserByUserID(userId int64) (*users_domain.User, *errors_utils.APIError) {
	user := users_domain.User{UserID: userId}

	if getUserErr := user.GetByUserID(); getUserErr != nil {
		return nil, getUserErr
	}
	return &user, nil
}

func PutUserByUserID(user users_domain.User) (*users_domain.User, *errors_utils.APIError) {
	if validateUserErr := user.Validate(); validateUserErr != nil {
		return nil, validateUserErr
	}

	currentUser, getUserErr := GetUserByUserID(user.UserID)
	if getUserErr != nil {
		return nil, getUserErr
	}

	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.Email = user.Email

	currentUser.PutByUserID()

	return currentUser, nil
}

func PatchUserByUserID(user users_domain.User) (*users_domain.User, *errors_utils.APIError) {
	currentUser, getUserErr := GetUserByUserID(user.UserID)
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
	if user.Status != "" {
		currentUser.Status = user.Status
	}
	if user.Password != "" {
		currentUser.Password = user.Password
	}

	currentUser.PutByUserID()

	return currentUser, nil
}

func DeleteUserByUserID(userId int64) (bool, *errors_utils.APIError) {
	user := users_domain.User{UserID: userId}

	if deleteUserErr := user.DeleteByUserID(); deleteUserErr != nil {
		return false, deleteUserErr
	}

	return true, nil
}

func SearchUsers(status string) (users_domain.Users, *errors_utils.APIError) {
	dao := &users_domain.User{}
	return dao.GetUsersByStatus(status)
}
