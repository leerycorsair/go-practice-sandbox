package nats

import (
	"encoding/json"
	"module/internal/models"

	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func (h *NatsHandler) CreateOrder(m *stan.Msg) {
	var order models.OrderT
	err := json.Unmarshal(m.Data, &order)
	if err != nil {
		logrus.Errorf("ERROR unmarshaling msg NATS:%s", err.Error())
	}
	_, err = h.services.CreateOrder(order)
	if err != nil {
		logrus.Errorf("ERROR:%s", err)
	}
}
