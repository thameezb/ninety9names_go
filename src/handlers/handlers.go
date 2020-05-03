package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thameezb/ninety9names/src/service"
	"net/http"
	"strconv"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetAll())
}

func GetName(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	name, err := service.GetName(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, name)
}
