Table of Contents
Project Overview
Architecture
Setup Instructions
Prerequisites
Installation
Running the Application
Features
Assumptions
Future Enhancements
Project Overview
This project is a backend system for a Product Management Application. It focuses on:

Asynchronous Image Processing: Compressing and storing images.
Caching: Using Redis to speed up frequently accessed data.
Scalability: A modular, robust design to handle high loads.
The system supports:

Creating, retrieving, and listing products.
Asynchronous tasks for image processing via message queues.
Caching product data to optimize response times.
Architecture
This system uses the following architectural components:

1. RESTful API
Built using Go with the gin framework for handling HTTP requests.
Supports endpoints for creating and retrieving product data.
2. PostgreSQL
Stores product and user data.
Ensures relational consistency between users and products.
3. Redis
Caches product details for faster retrieval, reducing database load.
Implements cache invalidation to ensure real-time accuracy.
4. Asynchronous Image Processing
RabbitMQ (or Kafka) for message queueing.
A dedicated image processing microservice:
Downloads product images.
Compresses images.
Updates the database with processed images.
5. Logging and Error Handling
Uses structured logging (logrus or zap) for debugging and monitoring.
Implements robust error handling for API, database, and queue interactions.

