package queue

import (
	"context"
	"encoding/json"

	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type queueRabbitMQ struct {
	ch   *amqp.Channel
	conf config.QueueInfo
}

func initRabbitMQ(conf *config.Config) (q *queueRabbitMQ, err error) {
	q = new(queueRabbitMQ)

	q.conf = conf.Queue

	conn, err := amqp.Dial(conf.Queue.Server)
	if err != nil {
		return
	}

	q.ch, err = conn.Channel()
	if err != nil {
		return
	}

	conf.AddCloser(func() error {
		q.ch.Close()
		conn.Close()
		return nil
	})

	err = q.ch.ExchangeDeclare(
		conf.Queue.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange": "peril_dlx",
		},
	)

	return
}

func (q *queueRabbitMQ) PublishJSON(key string, val any) (err error) {
	body, err := json.Marshal(val)
	if err != nil {
		return
	}

	if key == "" {
		key = q.conf.Key
	}

	err = q.ch.PublishWithContext(context.Background(),
		q.conf.Exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	return
}
