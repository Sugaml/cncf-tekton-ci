package queue

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitmqRoutine struct {
	Client *amqp.Channel
}

func NewRabbitmqRoutine() *RabbitmqRoutine {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatal("Failed to connect", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to create channel", err)
	}
	return &RabbitmqRoutine{
		Client: ch,
	}
}

func (c *RabbitmqRoutine) Subscribe(name string, method func(interface{}) error) {
	_, err := c.Client.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("queue not declared: %v", err)
	}
	msgs, err := c.Client.Consume(
		name, "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("Failed to consume", err)
	}
	for d := range msgs {
		log.Infof("Got message in queue :: %s ", name)
		log.Debug(string(d.Body))

		go func(msg amqp.Delivery) {
			err := method(msg.Body)
			if err != nil {
				log.Error(err)
			}
		}(d)
	}
}

func (c *RabbitmqRoutine) Publish(name string, req interface{}) error {
	_, err := c.Client.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("queue not declared: %v", err)
	}
	msg, err := json.Marshal(req)
	if err != nil {
		log.Error("Unable to convert model to json;")
		return err
	}
	err = c.Client.Publish(
		"", name, false, false,
		amqp.Publishing{Body: msg},
	)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("Published a message in queue " + name)
	return nil
}
