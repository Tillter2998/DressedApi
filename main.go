package main

import (
	"DressedApi/Config"
	"DressedApi/Services"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config := Config.NewConfig()
	db := Services.NewDatabase(config)
	ds := Services.NewDressService(&db)

	router := gin.Default()

	// TODO: Add 405 Method not allowed handling
	// TODO: Add Authentication

	router.GET("/dresses", func(c *gin.Context) {
		dresses, err := ds.GetDresses()
		if err != nil {
			c.JSON(500, err)
		}
		c.JSON(http.StatusOK, dresses)
	})
	router.GET("/dresses/:id", func(c *gin.Context) {
		dress, err := ds.GetDress(c.Param("id"))
		if err != nil {
			c.JSON(404, err)
		}
		c.JSON(http.StatusOK, dress)
	})
	router.POST("/dresses/", func(c *gin.Context) {
		var dress Services.Dress
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, err)
		}

		json.Unmarshal(body, &dress)

		response, err := ds.AddDress(&dress)
		if err != nil {
			c.JSON(500, err)
		}

		c.JSON(http.StatusOK, response)
	})
	router.PUT("/dresses/", func(c *gin.Context) {
		var dress Services.Dress
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, err)
		}

		json.Unmarshal(body, &dress)

		response, err := ds.UpdateDress(&dress)
		if err != nil {
			c.JSON(500, err)
		}

		c.JSON(http.StatusOK, response)
	})
	router.DELETE("/dresses/:id", func(c *gin.Context) {
		response, err := ds.DeleteDress(c.Param("id"))
		if err != nil {
			c.JSON(404, err)
		}

		c.JSON(http.StatusOK, response)
	})

	router.Run("localhost:8080")
}
