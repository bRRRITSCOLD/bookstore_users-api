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
	router.PUT("/users/:userId", users_controllers.PutUser)
	router.PATCH("/users/:userId", users_controllers.PatchUser)
	router.DELETE("/users/:userId", users_controllers.DeleteUser)
}
