package main

import (
	"github.com/devnandito/ginweb/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/modules/show", handlers.HandlerShowModule)
	router.Run(":8080")
}