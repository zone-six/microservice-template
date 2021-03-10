package pubsub

import (
	"github.com/nats-io/stan.go"
	"github.com/zone-six/microservice-template/internal/config"
	"github.com/zone-six/microservice-template/internal/managers"
	"github.com/zone-six/microservice-template/internal/utilities"
)

// NOTE: This is untested. Please test before using. There may be a much better way to handle this.

// Client is the PubSub Client
type Client interface {
	// RegisterSubscriptions creates all of the subscriptions
	RegisterSubscriptions()
	// CleanUp unsubscribes from all subscriptions
	CleanUp()
}

type pubsubClient struct {
	config     *config.Config
	managers   *managers.Container
	utilities  *utilities.Container
	closeChans []chan struct{}
}

// New returns a new PubSub client
func New(cfg *config.Config, managers *managers.Container, utilities *utilities.Container) Client {
	return &pubsubClient{config: cfg, managers: managers, utilities: utilities, closeChans: []chan struct{}{}}
}

type subscriberFunc func(<-chan struct{})

// RegisterSubscriptions creates all of the subscriptions
func (c *pubsubClient) RegisterSubscriptions() {
	subs := []subscriberFunc{
		c.exampleSubscription,
	}

	for _, f := range subs {
		close := make(chan struct{})
		c.closeChans = append(c.closeChans, close)
		// Fire off a go routine for each subscriber.
		// Subscribers must block.
		go f(close)
	}
}

// CleanUp unsubscribes from all subscriptions
func (c *pubsubClient) CleanUp() {
	for _, v := range c.closeChans {
		v <- struct{}{}
	}
}

// TODO: Could probably have a wrapper around this pattern.
func (c *pubsubClient) exampleSubscription(close <-chan struct{}) {
	// Typically use a service's library to get the subject.
	// QueueGroups and DurableGroupNames should be unique to the subscribing service.
	sub, err := c.utilities.PubSub.DurableQueueSubscribe(
		"Subject",
		"QueueGroup",
		"DurableGroupName",
		func(msg *stan.Msg) {
			// Unmarshal the message and process it.
		},
	)

	if err != nil {
		// Handle error
	}

	// Block until we recieve a signal to close
	<-close
	// Unsubscribe
	err = sub.Unsubscribe()
	if err != nil {
		// Handle error
	}
}
