package users

import (
	users_domain "bookstore_users-api/domains/users"

	crypto_utils "github.com/bRRRITSCOLD/bookstore_utils-go/crypto"
	dates_utils "github.com/bRRRITSCOLD/bookstore_utils-go/dates"
	errors_utils "github.com/bRRRITSCOLD/bookstore_utils-go/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	CreateUser(user users_domain.User) (*users_domain.User, errors_utils.APIError)
	GetUserByUserID(userId int64) (*users_domain.User, errors_utils.APIError)
	PutUserByUserID(user users_domain.User) (*users_domain.User, errors_utils.APIError)
	PatchUserByUserID(user users_domain.User) (*users_domain.User, errors_utils.APIError)
	DeleteUserByUserID(userId int64) (bool, errors_utils.APIError)
	SearchUsers(status string) (users_domain.Users, errors_utils.APIError)
	LoginUser(users_domain.UserLoginRequest) (*users_domain.User, errors_utils.APIError)
}

func (uS *usersService) CreateUser(user users_domain.User) (*users_domain.User, errors_utils.APIError) {
	if validateUserErr := user.Validate(); validateUserErr != nil {
		return nil, validateUserErr
	}

	user.DateCreated = dates_utils.GetNow()
	user.Status = users_domain.USER_ACTIVE_STATUS
	user.Password = crypto_utils.MD5Hash(user.Password)

	if saveUserErr := user.Save(); saveUserErr != nil {
		return nil, saveUserErr
	}

	return &user, nil
}

func (uS *usersService) GetUserByUserID(userId int64) (*users_domain.User, errors_utils.APIError) {
	user := users_domain.User{UserID: userId}

	if getUserErr := user.GetByUserID(); getUserErr != nil {
		return nil, getUserErr
	}
	return &user, nil
}

func (uS *usersService) LoginUser(ulr users_domain.UserLoginRequest) (*users_domain.User, errors_utils.APIError) {
	user := users_domain.User{Email: ulr.Email, Password: crypto_utils.MD5Hash(ulr.Password)}

	if GetUserByEmailAndPasswordErr := user.GetUserByEmailAndPassword(); GetUserByEmailAndPasswordErr != nil {
		return nil, GetUserByEmailAndPasswordErr
	}
	return &user, nil
}

func (uS *usersService) PutUserByUserID(user users_domain.User) (*users_domain.User, errors_utils.APIError) {
	if validateUserErr := user.Validate(); validateUserErr != nil {
		return nil, validateUserErr
	}

	currentUser, getUserErr := uS.GetUserByUserID(user.UserID)
	if getUserErr != nil {
		return nil, getUserErr
	}

	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.Email = user.Email

	currentUser.PutByUserID()

	return currentUser, nil
}

func (uS *usersService) PatchUserByUserID(user users_domain.User) (*users_domain.User, errors_utils.APIError) {
	currentUser, getUserErr := uS.GetUserByUserID(user.UserID)
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

func (uS *usersService) DeleteUserByUserID(userId int64) (bool, errors_utils.APIError) {
	user := users_domain.User{UserID: userId}

	if deleteUserErr := user.DeleteByUserID(); deleteUserErr != nil {
		return false, deleteUserErr
	}

	return true, nil
}

func (uS *usersService) SearchUsers(status string) (users_domain.Users, errors_utils.APIError) {
	dao := &users_domain.User{}
	return dao.GetUsersByStatus(status)
}
