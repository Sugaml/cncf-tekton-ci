package queue

import "os"

type Queue struct{}

func NewQueue() *Queue {
	return &Queue{}
}

type QueueInterface interface {
	GetQueue() QueueDefInterface
}
type QueueDefInterface interface {
	Subscribe(topicName string, function func(interface{}) error)
	Publish(topic string, req interface{}) error
}

func isPubsub() bool {
	return os.Getenv("MESSAGE_TYPE") == "" || os.Getenv("MESSAGE_TYPE") == "pubsub"
}
func isRabbitmq() bool {
	return os.Getenv("MESSAGE_TYPE") == "rabbitmq"
}

func (c *Queue) GetQueue() QueueDefInterface {
	if isPubsub() {
		return NewPubsubRoutine()
	}
	if isRabbitmq() {
		return NewRabbitmqRoutine()
	}
	return nil
}
