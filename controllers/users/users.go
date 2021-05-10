package users

import (
	usersDomain "bookstore_users-api/domains/users"
	usersService "bookstore_users-api/services/users"
	errorsUtils "bookstore_users-api/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user usersDomain.User

	if shouldBindJSONErr := c.ShouldBindJSON(&user); shouldBindJSONErr != nil {
		apiError := errorsUtils.NewBadRequestAPIError("invalid json body")
		c.JSON(apiError.Status, apiError)
		return
	}

	createUserResult, createUserErr := usersService.CreateUser(user)
	if createUserErr != nil {
		c.JSON(createUserErr.Status, createUserErr)
		return
	}

	c.JSON(http.StatusCreated, createUserResult)
}

func GetUser(c *gin.Context) {
	userId, userIdErr := strconv.ParseInt(c.Param("userId"), 10, 64)
	if userIdErr != nil {
		apiError := errorsUtils.NewBadRequestAPIError("invalid user id")
		c.JSON(apiError.Status, apiError)
		return
	}

	getUserResult, getUserErr := usersService.GetUser(userId)
	if getUserErr != nil {
		c.JSON(getUserErr.Status, getUserErr)
		return
	}

	c.JSON(http.StatusCreated, getUserResult)
}

// func SearchUsers(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me")
// }
