package app

import (
	"bookstore_users-api/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	logger.Info("{}app::#StartApp::starting execution")
	initRoutes()

	logger.Info("{}app::#StartApp::starting server")
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
	logger.Info("{}app::#StartApp::started server")

	logger.Info("{}app::#StartApp::finishing execution")
}
