package salesInvoice

import (
	"context"
	"encoding/json"
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

func GetSpecificResp(c *gin.Context) {
	log.Println("Getting specific Invoice")
	// Execute the query
	// params := c.Request.PostForm

	// // Print or use the parameters as needed
	// for key, values := range params {
	// 	for _, value := range values {
	// 		log.Println("%s = %s\n", key, value)
	// 	}
	// }
	objectIDStr := c.Param("id")
	log.Println("id", objectIDStr)
	objectID, err := primitive.ObjectIDFromHex(objectIDStr)
	if err != nil {
		log.Println(err)
	}
	filter := bson.M{"_id": objectID}

	// options := options.Find().SetLimit(1)

	cur, err := client.Collection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("Error executing find query for single response on edit: %v\n", err)
		return
	}
	defer cur.Close(context.Background())

	log.Println("getting Sales Invoice From DB")

	// Define a slice to store the query results
	var results []bson.M
	count := 0
	// Iterate through the cursor and decode each document into a bson.M map
	for cur.Next(context.Background()) {
		count++
		var result bson.M

		err := cur.Decode(&result)
		if err != nil {
			log.Printf("Error decoding document for edit: %v\n", err)
			return
		}
		results = append(results, result)
		jsonData, err := json.Marshal(results)
		if err != nil {
			log.Printf("Error converting to JSON: %v\n", err)
			return
		}

		c.Data(http.StatusOK, "application/json; charset=utf-8", jsonData)
	}
	if count == 0 {
		log.Println("No values in DB")
		data := gin.H{
			"message": "API token not found in DB",
			"status":  400,
		}
		// Set the content type to JSON
		c.Header("Content-Type", "application/json")

		c.IndentedJSON(http.StatusBadRequest, data)
	}
}

func GetAllProducts(c *gin.Context) {
	log.Println("Getting sales invoice")

	// Query the MongoDB collection to retrieve all documents
	cur, err := client.ProductCollection.Find(context.Background(), bson.D{})
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
