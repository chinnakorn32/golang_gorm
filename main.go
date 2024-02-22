package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	routes "github.com/itsmaheshkariya/gin-gorm-rest/routers"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}