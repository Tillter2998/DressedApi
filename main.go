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
	router.HandleMethodNotAllowed = true

	// Unsure if my approach to handling authentication is the best.
	// It might have to change down the line as I develop the front end

	// This will be changed to not be a hard coded user/pass
	auth := router.Group("/", gin.BasicAuth(gin.Accounts{
		"testUser": "testPass",
	}))

	auth.GET("/dresses", func(c *gin.Context) {
		dresses, err := ds.GetDresses()
		if err != nil {
			c.JSON(500, err)
		}
		c.JSON(http.StatusOK, dresses)
	})
	auth.GET("/dresses/:id", func(c *gin.Context) {
		dress, err := ds.GetDress(c.Param("id"))
		if err != nil {
			c.JSON(404, err)
		}
		c.JSON(http.StatusOK, dress)
	})
	auth.POST("/dresses/", func(c *gin.Context) {
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
	auth.PUT("/dresses/", func(c *gin.Context) {
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
	auth.DELETE("/dresses/:id", func(c *gin.Context) {
		response, err := ds.DeleteDress(c.Param("id"))
		if err != nil {
			c.JSON(404, err)
		}

		c.JSON(http.StatusOK, response)
	})

	//Unsure if this is the best way to handle HTTP/HTTPS, may change as development continues
	if config.ENVIRONMENT == "DEV" {
		router.Run("localhost:8080")
	} else {
		router.RunTLS(":8080", config.CERTFILE_LOCATION, config.KEYFILE_LOCATION)
	}
}
