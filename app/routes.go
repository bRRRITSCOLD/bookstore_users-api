package app

import (
	pingControllers "bookstore_users-api/controllers/ping"
	usersControllers "bookstore_users-api/controllers/users"
)

func initRoutes() {
	// ping
	router.GET("/ping", pingControllers.Ping)

	// users
	router.POST("/users", usersControllers.CreateUser)
	router.GET("/users/:userId", usersControllers.GetUser)
	// router.GET("/users/search", controllers.SearchUsers)
}
