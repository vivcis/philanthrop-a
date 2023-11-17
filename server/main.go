package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const coinremitterEndpoint = "https://coinremitter.com/api/v3/BTC/validate-address"

type CoinremitterRequest struct {
	ApiKey   string `json:"api_key"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type CoinremitterResponse struct {
	Flag   int    `json:"flag"`
	Msg    string `json:"msg"`
	Action string `json:"action"`
	Data   struct {
		Valid bool `json:"valid"`
	} `json:"data"`
}

func main() {
	router := gin.Default()

	// Define a route for the endpoint
	router.POST("/check-btc-address", CheckBTCAddress)

	fmt.Println("Server listening on :8080")
	router.Run(":8080")
}

func CheckBTCAddress(c *gin.Context) {
	// Decode JSON payload
	var data CoinremitterRequest
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
	response, err := http.Post(coinremitterEndpoint, "application/json", bytes.NewBuffer(requestBody))
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
	var responseData CoinremitterResponse
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
