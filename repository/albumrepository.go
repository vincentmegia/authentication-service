package repository

import (
	"authentication-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: "56.99"},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: "17.99"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: "39.99"},
}

var users = []models.User{
	{Firstname: "john", Lastname: "doe", Username: "johndoe", Password: "johndoe"},
	{Firstname: "john1", Lastname: "doe1", Username: "johndoe1", Password: "johndoe1"},
	{Firstname: "john2", Lastname: "doe2", Username: "johndoe2", Password: "johndoe2"},
	{Firstname: "john3", Lastname: "doe3", Username: "johndoe3", Password: "johndoe3"},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
