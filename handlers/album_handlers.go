package handlers

import (
	"database/sql"
	"eddybruv/albums-api/db"
	"eddybruv/albums-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetAlbums(c *gin.Context) {
	var albums []models.Album

	db.InitDB()
	rows, err := db.DB.Query(`select id, title, artist, price from album`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}(rows)

	for rows.Next() {
		var tempAlbum models.Album
		err := rows.Scan(&tempAlbum.ID, &tempAlbum.Title, &tempAlbum.Artist, &tempAlbum.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		albums = append(albums, tempAlbum)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumsByArtist(c *gin.Context) {
	artist := strings.ToLower(c.Param("artist"))

	var albums []models.Album

	query := "Select distinct * from album where lower(artist) like ?"
	rows, err := db.DB.Query(query, "%"+artist+"%")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}(rows)

	for rows.Next() {
		var tempAlb models.Album
		err := rows.Scan(&tempAlb.ID, &tempAlb.Title, &tempAlb.Artist, &tempAlb.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		albums = append(albums, tempAlb)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	var album models.Album

	if err := c.BindJSON(&album); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	_, err := db.DB.Query("INSERT into album (title, artist, price) values (?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Album added", "album": album})
}

//if err := c.BindJSON(&newAlbum); err != nil {
//return
//}
