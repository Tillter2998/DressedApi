package main

import (
	"DressedApi/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dress = Models.Dress

func main() {
	router := gin.Default()
	router.GET("/dresses", getDresses)

	router.Run("localhost:8080")
}

func getDresses(c *gin.Context) {
	// DbCall would go here, hard coded for now
	var dresses = []dress{
		{Id: primitive.NewObjectID(), Price: 10.99, Name: "Dress 1", Description: "First dress"},
		{Id: primitive.NewObjectID(), Price: 9.99, Name: "Dress 2", Description: "Second Dress"},
	}

	c.IndentedJSON(http.StatusOK, dresses)
}
