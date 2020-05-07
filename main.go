package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thameezb/ninety9names/src/handler"
	"github.com/thameezb/ninety9names/src/repository"
	"github.com/thameezb/ninety9names/src/service"
)

func router(port string, h handler.Interface) {
	router := gin.Default()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("resources/templates/*")

	router.GET("/ping", h.Ping)

	router.GET("/bff/names", h.GetAll)
	router.GET("/bff/names/:id", h.GetName)

	router.GET("/", h.GetAllNames)
	router.GET("/names/:id", h.GetSpecificName)

	router.GET("/challenge/english", h.GetRandomEnglish)
	router.GET("/challenge/arabic", h.GetRandomArabic)

	router.Run(":" + port)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("$DATABASE_URL must be set")
	}

	db := repository.New(dbURL)
	s := &service.Service{db}
	h := &handler.Handler{s}

	router(port, h)
}
