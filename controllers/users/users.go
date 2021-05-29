package users_controllers

import (
	users_domain "bookstore_users-api/domains/users"
	users_service "bookstore_users-api/services/users"
	errors_utils "bookstore_users-api/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func parseUserIDFromRequestPath(requestPathUserID string) (int64, *errors_utils.APIError) {
	userId, userIdErr := strconv.ParseInt(requestPathUserID, 10, 64)
	if userIdErr != nil {
		return 0, errors_utils.NewBadRequestAPIError("invalid user id")
	}
	return userId, nil
}

func CreateUser(c *gin.Context) {
	var user users_domain.User

	if shouldBindJSONErr := c.ShouldBindJSON(&user); shouldBindJSONErr != nil {
		apiError := errors_utils.NewBadRequestAPIError("invalid json body")
		c.JSON(apiError.Status, apiError)
		return
	}

	createUserResult, createUserErr := users_service.UsersService.CreateUser(user)
	if createUserErr != nil {
		c.JSON(createUserErr.Status, createUserErr)
		return
	}

	c.JSON(http.StatusCreated, createUserResult.Marshal(c.GetHeader("X-Public") == "true"))
}

func GetUserByUserID(c *gin.Context) {
	userId, userIdErr := parseUserIDFromRequestPath(c.Param("userId"))
	if userIdErr != nil {
		c.JSON(userIdErr.Status, userIdErr)
		return
	}

	getUserResult, getUserErr := users_service.UsersService.GetUserByUserID(userId)
	if getUserErr != nil {
		c.JSON(getUserErr.Status, getUserErr)
		return
	}

	c.JSON(http.StatusOK, getUserResult.Marshal(c.GetHeader("X-Public") == "true"))
}

func LoginUser(c *gin.Context) {
	var userLoginRequest users_domain.UserLoginRequest

	if shouldBindJSONErr := c.ShouldBindJSON(&userLoginRequest); shouldBindJSONErr != nil {
		apiError := errors_utils.NewBadRequestAPIError("invalid json body")
		c.JSON(apiError.Status, apiError)
		return
	}

	user, loginUserErr := users_service.UsersService.LoginUser(userLoginRequest)
	if loginUserErr != nil {
		c.JSON(loginUserErr.Status, loginUserErr)
		return
	}

	c.JSON(http.StatusCreated, user.Marshal(c.GetHeader("X-Public") == "true"))
}

func PutUserByUserID(c *gin.Context) {
	userId, userIdErr := parseUserIDFromRequestPath(c.Param("userId"))
	if userIdErr != nil {
		c.JSON(userIdErr.Status, userIdErr)
		return
	}

	var user users_domain.User

	if shouldBindJSONErr := c.ShouldBindJSON(&user); shouldBindJSONErr != nil {
		apiError := errors_utils.NewBadRequestAPIError("invalid json body")
		c.JSON(apiError.Status, apiError)
		return
	}

	user.UserID = userId

	updateUserResult, updateUserErr := users_service.UsersService.PutUserByUserID(user)
	if updateUserErr != nil {
		c.JSON(updateUserErr.Status, updateUserErr)
		return
	}

	c.JSON(http.StatusOK, updateUserResult.Marshal(c.GetHeader("X-Public") == "true"))
}

func PatchUserByUserID(c *gin.Context) {
	userId, userIdErr := parseUserIDFromRequestPath(c.Param("userId"))
	if userIdErr != nil {
		c.JSON(userIdErr.Status, userIdErr)
		return
	}

	var user users_domain.User

	if shouldBindJSONErr := c.ShouldBindJSON(&user); shouldBindJSONErr != nil {
		apiError := errors_utils.NewBadRequestAPIError("invalid json body")
		c.JSON(apiError.Status, apiError)
		return
	}

	user.UserID = userId

	updateUserResult, updateUserErr := users_service.UsersService.PatchUserByUserID(user)
	if updateUserErr != nil {
		c.JSON(updateUserErr.Status, updateUserErr)
		return
	}

	c.JSON(http.StatusOK, updateUserResult.Marshal(c.GetHeader("X-Public") == "true"))
}

func DeleteUserByUserID(c *gin.Context) {
	userId, userIdErr := parseUserIDFromRequestPath(c.Param("userId"))
	if userIdErr != nil {
		c.JSON(userIdErr.Status, userIdErr)
		return
	}

	deleteUserResult, deleteUserErr := users_service.UsersService.DeleteUserByUserID(userId)
	if deleteUserErr != nil {
		c.JSON(deleteUserErr.Status, deleteUserErr)
		return
	}

	// c.JSON(http.StatusCreated, deleteUserResult)
	c.JSON(http.StatusNoContent, map[string]bool{"delete": deleteUserResult})
}

func SearchUsers(c *gin.Context) {
	status := c.Query("status")

	users, err := users_service.UsersService.SearchUsers(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))
}
