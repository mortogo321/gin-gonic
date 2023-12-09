package main

import (
	"gin-gonic/initializers"
	"gin-gonic/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Product{})
}
