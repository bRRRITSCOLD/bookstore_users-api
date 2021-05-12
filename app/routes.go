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
	router.GET("/users/:userId", users_controllers.GetUser)
	// router.GET("/users/search", controllers.SearchUsers)
}
