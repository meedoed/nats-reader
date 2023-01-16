package streaming

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

var stream = "WH_ITEMS"

type natsPublisher struct {
	config *Config
	nc     *nats.Conn
	js     nats.JetStream
	sub    *nats.Subscription
	logger *logrus.Logger
}

func New() *natsPublisher {
	return &natsPublisher{}
}

func (n *natsPublisher) Init() {
	config := NewConfig()
	logger := logrus.New()
	n.config = config
	n.logger = logger

	natsconn, err := nats.Connect(n.config.Address)
	if err != nil {
		n.logger.Error(err)
	}
	jetstream, err := natsconn.JetStream()
	if err != nil {
		n.logger.Error(err)
		return
	}

	_, err = jetstream.AddStream(&nats.StreamConfig{
		Name:     stream,
		Subjects: []string{fmt.Sprintf("%s.*", stream)},
	})
	if err != nil {
		n.logger.Error(err)
		return
	}
	n.nc = natsconn
	n.js = jetstream

	subscription, err := n.js.Subscribe(stream+".*", n.getMessage)

	if err != nil {
		n.logger.Error(err)
	}
}
