package main

import (
    "product-management/internal/api"
    "product-management/internal/db"
    "product-management/internal/mq"
    "product-management/pkg/logger"
)

func main() {
    logger.InitLogger()
    db.InitDB()
    mq.InitQueue()

    api.RunServer() // Start the API server
}
