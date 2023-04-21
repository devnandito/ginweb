package handlers

import (
	"net/http"

	"github.com/devnandito/ginweb/models"
	"github.com/devnandito/ginweb/utils"
	"github.com/gin-gonic/gin"
)

var mod models.Module

func HandlerShowModule(c *gin.Context) {
	objects, err := mod.ShowModuleGorm()
	utils.CheckError(err)
	headers := [3]string{"ID", "Description", "Action"}
	c.HTML(http.StatusOK, "modules/show.html", gin.H{
		"Title": "Modules",
		"Objects": objects,
		"Headers": headers,
	})
}