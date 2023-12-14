package main

import (
	"invoice/app/salesInvoice"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Loading")
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/AllSalesInvoice", salesInvoice.GetAllInvoices)

	r.POST("", salesInvoice.Add)

	r.Run()
}
