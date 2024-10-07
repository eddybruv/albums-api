package main

import (
	"eddybruv/albums-api/db"
	"eddybruv/albums-api/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load env file
	if envErr := godotenv.Load(); envErr != nil {
		fmt.Println("Error loading .env file")
	}

	//initialize db
	db.InitDB()

	//gin router
	router := gin.Default()

	//routes
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:artist", handlers.GetAlbumsByArtist)
	router.POST("/albums", handlers.PostAlbum)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
