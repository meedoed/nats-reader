package streaming

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"nats-publisher/models"
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

	nc, err := nats.Connect(n.config.Address)
	if err != nil {
		n.logger.Error(err)
	}
	js, err := nc.JetStream()
	if err != nil {
		n.logger.Error(err)
		return
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     stream,
		Subjects: []string{fmt.Sprintf("%s.*", stream)},
	})
	if err != nil {
		n.logger.Error(err)
		return
	}
	n.nc = nc
	n.js = js

}

func (n *natsPublisher) Publish(order *models.Order) error {
	orderData, err := json.Marshal(order)
	if err != nil {
		return err
	}
	n.js.Publish(stream, orderData)
	return nil
}
