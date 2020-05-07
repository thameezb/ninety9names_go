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
	GetAllNames(c *gin.Context)
	GetSpecificName(c *gin.Context)
	GetRandomName(c *gin.Context)
}

type Handler struct {
	Service service.Interface
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func (h *Handler) GetAllNames(c *gin.Context) {
	names, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "names.html", gin.H{
		"names": names,
	})
}

func (h *Handler) GetSpecificName(c *gin.Context) {
	h.getName(c.Param("id"), c)
}

func (h *Handler) GetRandomName(c *gin.Context) {
	h.getName("", c)
}

func (h *Handler) getName(id string, c *gin.Context) {
	name, err := h.Service.GetName(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusFound, "name.html", gin.H{
		"name": name,
	})
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
