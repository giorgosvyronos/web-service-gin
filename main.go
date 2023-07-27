package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// album slice to seek record album data
var albums = []album{
    {ID: "0", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "1", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "2", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id",getAlbumByID)
    router.POST("/albums", postAlbums)
    router.POST("/albums/:id", modifyAlbumByID)
    router.DELETE("/albums/:id", deleteAlbumByID)

    router.Run("localhost:8080")
}


// getAlbums : responds with the list of all albums in JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}


// getAlbumByID locates the album whose ID value matches the id parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for album with ID value matching parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    } 
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


// postAlbums : adds an album to the albums from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind to the received JSON to newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }
    id := fmt.Sprint(len(albums))
    newAlbum.ID = id
    // Add newAlbum to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}


// modifyAlbumByID : modifies pre-existing album by ID from JSON received in the request body.
func modifyAlbumByID(c *gin.Context) {
    id := c.Param("id")
    var newAlbum album

    // Call BindJSON to bind to the received JSON to newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    for a := range albums {
        if albums[a].ID == id {
            albums[a].Price = newAlbum.Price
            albums[a].Title = newAlbum.Title
            albums[a].Artist = newAlbum.Artist
            c.IndentedJSON(http.StatusOK, albums)
            return
        }
    }
    c.IndentedJSON(http.StatusNotModified, newAlbum)
}


// deleteAlbumByID : deletes an album from the albums from JSON received in the request body.
func deleteAlbumByID(c *gin.Context) {
    var newAlbums []album
    found := false
    newIndex := 0
    id := c.Param("id")

    // Loop over the list of albums, looking for album with ID value matching parameter.
    for _, a := range albums {
        if a.ID != id {
            a.ID = fmt.Sprint(newIndex)
            newAlbums = append(newAlbums,a)
            newIndex += 1
        } else {
            found = true
        }
    } 
    if found {
        albums = newAlbums
        c.IndentedJSON(http.StatusOK, albums)
    } else { 
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album to be deleted not found"})
    }
}
