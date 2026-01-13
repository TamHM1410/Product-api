package rabbitmq

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// SendMessage gửi message vào queue bình thường.
// Queue được declare durable và auto-create nếu chưa tồn tại.
func SendMessage(conn *amqp.Connection, queueName string, messageBody string) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declare queue với DLX option nếu cần
	args := amqp.Table{
		"x-dead-letter-exchange": "dlx.exchange", // message fail sẽ vào DLX
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		args,  // arguments
	)
	if err != nil {
		return err
	}

	// Tạo context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Publish message
	err = ch.PublishWithContext(ctx,
		"",     // exchange trống → gửi trực tiếp vào queue
		q.Name, // routing key = queue name
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json", // protobuf -> json
			Body:        []byte(messageBody),
			Timestamp:   time.Now(),
		},
	)
	if err != nil {
		return err
	}

	log.Printf("[x] Sent message to queue %s: %s", queueName, messageBody)
	return nil
}
