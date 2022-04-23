package main

import (
	"DressedApi/Config"
	"DressedApi/Services"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	router.GET("/dresses/:id", func(c *gin.Context) {
		dress, err := ds.GetDress(c.Param("id"))
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, dress)
	})
	router.POST("/dresses/new-dress", func(c *gin.Context) {
		var dress Services.Dress
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(body, &dress)

		response, err := ds.AddDress(&dress)

		c.JSON(http.StatusOK, response)
	})

	router.Run("localhost:8080")
}
