package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func createProduct(c *gin.Context) {
    var product struct {
        UserID          string   `json:"user_id"`
        ProductName     string   `json:"product_name"`
        ProductDesc     string   `json:"product_description"`
        ProductImages   []string `json:"product_images"`
        ProductPrice    float64  `json:"product_price"`
    }

    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save to database and queue images for processing
    go queueImages(product.ProductImages) // Async image processing task

    c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}
