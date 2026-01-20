package queue

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T, priority uint8, delay uint8) error {
	jsonData, err := json.Marshal(val)
	if err != nil {
		return err
	}

	return ch.PublishWithContext(
		context.Background(),
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Priority: priority,
			/*Headers: amqp.Table{
                "x-delay": int32(delay * 1000),
            },*/
			Body: jsonData,
		},
	)

}
