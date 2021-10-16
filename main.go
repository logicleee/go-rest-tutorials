package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album record
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice, i.e. sample data
var albums = []album{
	{ID: "1", Title: "Somethin' Else", Artist: "Cannonball Adderley", Price: 19.58},
	{ID: "2", Title: "Blue Train", Artist: "John Coltrane", Price: 19.58},
	{ID: "3", Title: "Milestones", Artist: "Miles Davis", Price: 19.58},
	{ID: "4", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 19.54},
	{ID: "5", Title: "Piano Starts Here", Artist: "Art Tatum", Price: 19.33},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums responder
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// this is how to translate (bind) JSON from  request
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add new album to slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// whish I knew how to do array.map or something
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
