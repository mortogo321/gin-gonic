package main

import (
	"gin-gonic/controllers"
	"gin-gonic/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("", controllers.Home)
	r.Run()
}
