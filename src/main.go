package main

import (
	"html/template"
	"strings"

	"github.com/devnandito/ginweb/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	router.Static("/assets", "./assets")
	// router.LoadHTMLGlob("templates/*.tmpl")
	router.LoadHTMLFiles("./templates/modules/show.tmpl")
	router.GET("/modules/show", handlers.HandlerShowModule)
	router.Run(":8080")
}