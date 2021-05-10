package users

import (
	usersDomain "bookstore_users-api/domains/users"
	errorsUtils "bookstore_users-api/utils/errors"
)

func CreateUser(user usersDomain.User) (*usersDomain.User, *errorsUtils.APIError) {
	if validateUserErr := user.Validate(); validateUserErr != nil {
		return nil, validateUserErr
	}

	if saveUserErr := user.Save(); saveUserErr != nil {
		return nil, saveUserErr
	}

	return &user, nil
}

func GetUser(userId int64) (*usersDomain.User, *errorsUtils.APIError) {
	user := usersDomain.User{UserID: userId}
	if getUserErr := user.GetByUserID(); getUserErr != nil {
		return nil, getUserErr
	}
	return &user, nil
}

func FindUser() {

}
