# Product Management System with Asynchronous Image Processing


## Table of Contents


a)Project Overview

b)Architecture

c)Setup Instructions

  i)Prerequisites
  
  ii)Installation
  
  iii)Running the Application
  
d)Features

e)Assumptions

f)Future Enhancements

## Project Overview

This project is a backend system for a Product Management Application. It focuses on:

**Asynchronous Image Processing**: Compressing and storing images.

**Caching**: Using Redis to speed up frequently accessed data.

**Scalability**: A modular, robust design to handle high loads.

The system supports:

a)Creating, retrieving, and listing products.

b)Asynchronous tasks for image processing via message queues.

c)Caching product data to optimize response times.

## Architecture

This system uses the following architectural components:

### 1. RESTful API
Built using Go with the gin framework for handling HTTP requests.
Supports endpoints for creating and retrieving product data.
### 2. PostgreSQL
Stores product and user data.
Ensures relational consistency between users and products.
### 3. Redis
Caches product details for faster retrieval, reducing database load.
Implements cache invalidation to ensure real-time accuracy.
### 4. Asynchronous Image Processing
RabbitMQ (or Kafka) for message queueing.
A dedicated image processing microservice:
Downloads product images.
Compresses images.
Updates the database with processed images.
### 5. Logging and Error Handling
Uses structured logging (logrus or zap) for debugging and monitoring.
Implements robust error handling for API, database, and queue interactions.

## Setup Instructions
Prerequisites
Go 1.20+ installed on your system.
PostgreSQL, Redis, and RabbitMQ installed (or Docker if available).
Installation
### 1. Clone the Repository

git clone https://github.com/yourusername/product-management.git
cd product-management
### 2. Install Dependencies

go mod tidy
### 3. Set Up Environment Variables
Create a .env file in the root directory:

dotenv

#Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=productuser
DB_PASSWORD=secret
DB_NAME=productdb

#Redis
REDIS_ADDR=localhost:6379

#RabbitMQ
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
### 4. Initialize Database
Run the SQL schema located in configs/db.sql:


psql -U productuser -d productdb -f configs/db.sql
### 5. Start Services
Run PostgreSQL, Redis, and RabbitMQ. Use service commands or respective GUIs for your OS.

## Running the Application

### 1. Start the API Server

go run cmd/main.go
### 2. Verify API Endpoints
Use a tool like Postman or curl to interact with the API:

**Create a Product:**


curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{
      "user_id": "1234-5678-9012",
      "product_name": "Sample Product",
      "product_description": "A test product.",
      "product_images": ["https://example.com/image1.jpg", "https://example.com/image2.jpg"],
      "product_price": 100.0
    }'
    
**Retrieve Product Details:**


curl http://localhost:8080/products/1234

**List Products:**


curl http://localhost:8080/products?user_id=1234&min_price=50&max_price=200
### 3. Start the Image Processing Microservice

go run internal/imageproc/worker.go

## Features

### Implemented

1)Create, retrieve, and list products.

2)Asynchronous image processing with RabbitMQ.

3)Caching product data with Redis.

4)Robust logging and error handling.

### In Progress

1)Pagination for product listings.

2)Improved benchmark testing.

### Assumptions

**1)Image URLs**: The URLs provided are valid and accessible.

**2)Single Database**: PostgreSQL is used for both user and product data.

**3)Redis as Primary Cache**: Redis caching is sufficient for handling load spikes.

**4)Asynchronous Processing**: Image processing is handled entirely by the queue and worker service.

## Future Enhancements

### 1. Scalability Improvements:

a) Add horizontal scaling for the API and image processing services.

b)Introduce distributed caching mechanisms.

### 2. Enhanced Security:

a)Add OAuth 2.0 for user authentication.

b)Use HTTPS for secure communication.

### 3. Monitoring:

Integrate with Prometheus and Grafana for system monitoring.

### 4. Additional Features:

a)Add search functionality for product descriptions.
b)Implement support for image resizing options.
