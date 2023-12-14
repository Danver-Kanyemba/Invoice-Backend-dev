package salesInvoice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"invoice/app/client"
	"log"
	"net/http"
)

func GetAllInvoices(c *gin.Context) {
	log.Println("Getting sales invoice")

	// Query the MongoDB collection to retrieve all documents
	cur, err := client.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("Error executing find query: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer cur.Close(context.Background())

	// Define a slice to store the query results
	var results []bson.M

	// Iterate through the cursor and append the documents to the results slice
	for cur.Next(context.Background()) {
		var result bson.M
		if err := cur.Decode(&result); err != nil {
			log.Printf("Error decoding document: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		results = append(results, result)
	}

	// Check if any error occurred during cursor iteration
	if err := cur.Err(); err != nil {
		log.Printf("Error iterating through cursor: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Convert the query results to a JSON string and send it as the response
	jsonData, err := json.Marshal(results)
	if err != nil {
		log.Printf("Error converting to JSON: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", jsonData)
}
