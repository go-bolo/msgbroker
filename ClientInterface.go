package msgbroker

import "github.com/go-bolo/bolo"

type Client interface {
	Init(app bolo.App) error
	// Register a handler function to receive the new messages
	// Now we support only one method for each queues
	Subscribe(queueName string, handler MessageHandler) (string, error)
	// Unsubscribe handler function with subscriberID
	UnSubscribe(subscriberID string)
	// Publish one message to queue
	Publish(queueName string, data []byte) error
	// Get queue by queueName
	GetQueue(name string) Queue
	// Set one queue in queue list
	SetQueue(name string, queue Queue) error
}

type Queue interface {
	GetName() string
	SetName(name string) error
	GetHandler() MessageHandler
	SetHandler(handler MessageHandler) error
}

type MessageHandler interface {
	HandleMessage(queueName string, message Message) error
}

type Message interface {
	GetData() *[]byte
}
