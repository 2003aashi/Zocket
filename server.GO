package mq

import (
    "log"

    "github.com/streadway/amqp"
)

var Conn *amqp.Connection
var Channel *amqp.Channel
var Queue amqp.Queue

func InitQueue() {
    var err error
    Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }

    Channel, err = Conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open RabbitMQ channel: %v", err)
    }

    Queue, err = Channel.QueueDeclare(
        "image_tasks",
        false,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Failed to declare RabbitMQ queue: %v", err)
    }

    log.Println("RabbitMQ initialized.")
}
