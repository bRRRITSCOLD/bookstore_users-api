package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApp() {
	initRoutes()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
