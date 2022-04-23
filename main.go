package main

import (
	"DressedApi/Config"
	"DressedApi/Services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type dress = Services.Dress

func main() {
	config := Config.NewConfig()
	db := Services.NewDatabase(config)
	ds := Services.NewDressService(&db)

	router := gin.Default()
	router.GET("/dresses", func(c *gin.Context) {
		dresses, err := ds.GetDresses()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, dresses)
	})

	router.Run("localhost:8080")
}
