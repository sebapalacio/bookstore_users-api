package app

import "github.com/gin-gonic/gin"

var (
	//we use gin-gonic as web frameworks for microservices.
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
