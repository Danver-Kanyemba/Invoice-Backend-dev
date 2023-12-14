package salesInvoice

import (
	"context"
	"github.com/gin-gonic/gin"
	"invoice/app/client"
	"net/http"
)

func Add(c *gin.Context) {
	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	_, err := client.Collection.InsertOne(context.Background(), jsonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert JSON data: " + err})

		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Invoice added successfully"})
}
