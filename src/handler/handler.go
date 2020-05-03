package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thameezb/ninety9names/src/service"
)

type Interface interface {
	Ping(c *gin.Context)
	GetAll(c *gin.Context)
	GetName(c *gin.Context)
}

type Handler struct {
	Service service.Interface
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func (h *Handler) GetAll(c *gin.Context) {
	names, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, names)
}

func (h *Handler) GetName(c *gin.Context) {
	name, err := h.Service.GetName(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, name)
}
