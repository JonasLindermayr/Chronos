package handlers

import (
	"github.com/JonasLindermayr/Chronos/handlers/controller"
	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {
	controller.GetUser(c)
}