package imageproc

import (
    "log"
    "github.com/streadway/amqp"
)

func ProcessImages() {
    msgs, err := mq.Channel.Consume(
        mq.Queue.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Failed to consume RabbitMQ messages: %v", err)
    }

    for msg := range msgs {
        log.Printf("Processing image task: %s", msg.Body)
        // Download and compress images here
    }
}
