package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType,
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

	// --- 1) Dead Letter Exchange ---
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

	// --- 2) Dead Letter Queue ---
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

	err = ch.QueueBind(
		dlq.Name,
		"orders.failed",
		"ShakerQueue_dlx",
		false,
		nil,
	)
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	// --- 3) Delayed Exchange ---
	/*
	err = ch.ExchangeDeclare(
		exchange,
		"x-delayed-message",
		true,
		false,
		false,
		false,
		/*amqp.Table{
			"x-delayed-type": "direct",
		},
	)
	if err != nil {
		return nil, amqp.Queue{}, err
	}
	*/
	// --- 4) Main Queue (Priority + DLX) ---
	q, err := ch.QueueDeclare(
		queueName,
		durable,
		autoDelete,
		exclusive,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    "ShakerQueue_dlx",
			"x-dead-letter-routing-key": "orders.failed",
			"x-max-priority":           10,
		},
	)
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	// --- 5) Bind Queue to Delayed Exchange ---
	err = ch.QueueBind(
		q.Name,
		key,
		exchange,
		false,
		nil,
	)
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	return ch, q, nil
}