package app

import (
	ping_controllers "bookstore_users-api/controllers/ping"
	users_controllers "bookstore_users-api/controllers/users"
)

func initRoutes() {
	// ping
	router.GET("/ping", ping_controllers.Ping)

	// users
	router.POST("/users", users_controllers.CreateUser)
	router.GET("/users/search", users_controllers.SearchUsers)
	router.POST("/users/login", users_controllers.LoginUser)
	router.GET("/users/:userId", users_controllers.GetUserByUserID)
	router.PUT("/users/:userId", users_controllers.PutUserByUserID)
	router.PATCH("/users/:userId", users_controllers.PatchUserByUserID)
	router.DELETE("/users/:userId", users_controllers.DeleteUserByUserID)
}
