package app

import (
	users_mysql_db "bookstore_users-api/datasources/users/mysql"
	"bookstore_users-api/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

func StartApp() {
	logger.Info("{}app::#StartApp::starting execution")

	if err := godotenv.Load(); err != nil {
		logger.Info("{}app::#StartApp::error loading .env")
		panic(err)
	}

	users_mysql_db.Init()

	initRoutes()

	logger.Info("{}app::#StartApp::starting server")
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

	logger.Info("{}app::#StartApp::started server")

	logger.Info("{}app::#StartApp::finishing execution")
}
