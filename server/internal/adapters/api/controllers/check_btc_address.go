package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"server/internal/core/models"
)

func (u *HTTPHandler) CheckBTCAddress(c *gin.Context) {
	// Decode JSON payload
	var data models.CoinremitterRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Check if the required fields are present
	if data.ApiKey == "" || data.Password == "" || data.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// Create a request payload
	requestBody, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Make a POST request to the external endpoint
	response, err := http.Post(models.CoinremitterEndpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error making HTTP request"})
		return
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response body"})
		return
	}

	// Parse the response JSON
	var responseData models.CoinremitterResponse
	err = json.Unmarshal(responseBody, &responseData)
	log.Println(responseData, "rD", string(responseBody), "Rb")
	if err != nil {
		fmt.Println(err.Error(), "ERROR ...")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding JSON response"})
		return
	}

	// Return the result
	c.JSON(http.StatusOK, gin.H{"btc_address": data.Address, "data": responseData})
}
