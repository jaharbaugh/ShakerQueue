package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType, // an enum to represent "durable" or "transient"
) (*amqp.Channel, amqp.Queue, error) {

	ch, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	var durable, autoDelete, exclusive bool

	if queueType == SimpleQueueDurable {
		durable = true
		autoDelete = false
		exclusive = false
	} else {
		durable = false
		autoDelete = true
		exclusive = true
	}

    // --- 1) Declare DLX and DLQ ---
    err = ch.ExchangeDeclare(
        "ShakerQueue_dlx",
        "direct", 
        true,     
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        return nil, amqp.Queue{}, err
    }

    dlq, err := ch.QueueDeclare(
        "orders.dead", 
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        return nil, amqp.Queue{}, err
    }

    // Bind DLQ to DLX with a routing key
    err = ch.QueueBind(dlq.Name, "orders.failed", "ShakerQueue_dlx", false, nil)
    if err != nil {
        return nil, amqp.Queue{}, err
    }

    // --- 2) Declare main queue with DLX attached ---
    q, err := ch.QueueDeclare(
        queueName,
        durable,
        autoDelete,
        exclusive,
        false,
        amqp.Table{
            "x-dead-letter-exchange":    "ShakerQueue_dlx",
            "x-dead-letter-routing-key": "orders.failed",
        },
    )
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	err = ch.QueueBind(q.Name, key, exchange, false, amqp.Table(nil))
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	return ch, q, err
}
