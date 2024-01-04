package main

import (
	"invoice/app/salesInvoice"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Server started")
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/status_check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server reached",
		})
	})

	r.GET("/AllSalesInvoice", salesInvoice.GetAllInvoices)
	r.GET("/AllProducts", salesInvoice.GetAllProducts)
	r.GET("/Invoice/:id", salesInvoice.GetSpecificResp)
	r.POST("AddInvoice", salesInvoice.Add)

	r.Run(":3066")
}
