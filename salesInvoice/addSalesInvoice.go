package salesInvoice

import (
	"context"
	"invoice/app/client"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// add time
	jsonData["created_at"] = time.Now()

	_, err := client.Collection.InsertOne(context.Background(), jsonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Invoice added successfully",
		"status":  200,
	})
}
