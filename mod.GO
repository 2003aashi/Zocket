package api

import (
    "github.com/gin-gonic/gin"
)

func RunServer() {
    r := gin.Default()

    r.POST("/products", createProduct)
    r.GET("/products/:id", getProduct)
    r.GET("/products", listProducts)

    r.Run(":8080") // Start server on port 8080
}
