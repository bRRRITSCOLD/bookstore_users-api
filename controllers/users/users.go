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

	createUserResult, createUserErr := users_service.CreateUser(user)
	if createUserErr != nil {
		c.JSON(createUserErr.Status, createUserErr)
		return
	}

	c.JSON(http.StatusCreated, createUserResult)
}

func GetUser(c *gin.Context) {
	userId, userIdErr := parseUserIDFromRequestPath(c.Param("userId"))
	if userIdErr != nil {
		c.JSON(userIdErr.Status, userIdErr)
		return
	}

	getUserResult, getUserErr := users_service.GetUser(userId)
	if getUserErr != nil {
		c.JSON(getUserErr.Status, getUserErr)
		return
	}

	c.JSON(http.StatusOK, getUserResult)
}

func PutUser(c *gin.Context) {
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

	updateUserResult, updateUserErr := users_service.PutUser(user)
	if updateUserErr != nil {
		c.JSON(updateUserErr.Status, updateUserErr)
		return
	}

	c.JSON(http.StatusOK, updateUserResult)
}

func PatchUser(c *gin.Context) {
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

	updateUserResult, updateUserErr := users_service.PatchUser(user)
	if updateUserErr != nil {
		c.JSON(updateUserErr.Status, updateUserErr)
		return
	}

	c.JSON(http.StatusOK, updateUserResult)
}

func DeleteUser(c *gin.Context) {
	userId, userIdErr := parseUserIDFromRequestPath(c.Param("userId"))
	if userIdErr != nil {
		c.JSON(userIdErr.Status, userIdErr)
		return
	}

	deleteUserResult, deleteUserErr := users_service.DeleteUser(userId)
	if deleteUserErr != nil {
		c.JSON(deleteUserErr.Status, deleteUserErr)
		return
	}

	// c.JSON(http.StatusCreated, deleteUserResult)
	c.JSON(http.StatusNoContent, map[string]bool{"delete": deleteUserResult})
}

// func SearchUsers(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me")
// }
