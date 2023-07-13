package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	log "github.com/sirupsen/logrus"
)

type PubsubRoutine struct {
	Client *Client
}

func NewPubsubRoutine() *PubsubRoutine {
	client, err := NewPubsubClient()
	if err != nil {
		log.Fatalf("Client not initialized: %v", err)
	}
	return &PubsubRoutine{
		Client: client,
	}
}
func (d *PubsubRoutine) Subscribe(topic string, method func(interface{}) error) {
	topic = fmt.Sprintf("%s-%s", os.Getenv("GCLOUD_NAMESPACE"), topic)
	err := d.Client.ReceiveMessage(topic, method)
	if err != nil {
		log.Error(err)
	}
}

func (d *PubsubRoutine) Publish(topic string, req interface{}) error {
	topic = fmt.Sprintf("%s-%s", os.Getenv("GCLOUD_NAMESPACE"), topic)
	_, err := d.Client.Publish(topic, req)
	if err != nil {
		return err
	}
	log.Infof("Published message in %s", topic)
	return nil
}

type Client struct {
	Client *pubsub.Client
}

func NewPubsubClient() (*Client, error) {
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, os.Getenv("GCLOUD_PROJECT"))
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}
	return &Client{
		Client: pubsubClient,
	}, err
}

func (c *Client) CreateTopicIfNotExists(topic string) *pubsub.Topic {
	ctx := context.Background()
	t := c.Client.Topic(topic)
	ok, err := t.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		return t
	}
	t, err = c.Client.CreateTopic(ctx, topic)
	if err != nil {
		log.Fatalf("Failed to create the topic: %v", err)
	}
	return t
}

func (c *Client) CreateSubscription(name string, topic *pubsub.Topic) error {
	ctx := context.Background()
	t := c.Client.Subscription(name)
	ok, _ := t.Exists(ctx)
	if ok {
		return nil
	}
	sub, err := c.Client.CreateSubscription(ctx, name, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 2 * time.Minute,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Created subscription: %v\n", sub)
	return nil
}

func (c *Client) ClosePubConnection() {
	err := c.Client.Close()
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("PubSub Connection Closed")
}

func (c *Client) Publish(name string, input interface{}) (string, error) {
	t := c.CreateTopicIfNotExists(name)
	msg, err := json.Marshal(input)
	if err != nil {
		log.Error("Unable to convert model to json;")
		return "", err
	}
	ctx := context.Background()
	result := t.Publish(ctx, &pubsub.Message{
		Data: msg,
	})
	return result.Get(ctx)
}

func (c *Client) ReceiveMessage(name string, release func(input interface{}) error) error {
	t := c.CreateTopicIfNotExists(name)
	if err := c.CreateSubscription(name, t); err != nil {
		log.Error(err)
		return err
	}
	sub := c.Client.Subscription(name)
	cctx, ct := context.WithCancel(context.Background())
	defer ct()
	err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		log.Printf("Got message: %s :: %s", name, msg.Data)
		//defer msg.Ack()
		err := release(msg.Data)
		if err != nil {
			log.Error(err)
		}
		msg.Ack()
	})
	if err != nil {
		return err
	}
	return nil
}
