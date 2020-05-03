package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thameezb/ninety9names/src/handlers"
	"github.com/thameezb/ninety9names/src/lib"
)

func readData() {
	err := lib.ConvertToCSV("orig.xlsx", "./input/input.csv", "Names")
	if err != nil {
		log.Print(err)
	}

	err = lib.SetNames("./input/input.csv")
	if err != nil {
		log.Print(err)
	}
}

func router(port string) {
	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/ping", handlers.Ping)
	router.GET("/bff/names", handlers.GetAll)
	router.GET("/bff/names/:id", handlers.GetName)

	router.Run(":" + port)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	readData()
	router(port)
}
