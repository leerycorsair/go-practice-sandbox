package nats

import (
	"module/internal/service"
	"time"

	"github.com/nats-io/stan.go"
	_ "github.com/nats-io/stan.go/pb"
)

type NatsConfig struct {
	ClusterId string
	ClientId  string
	Host      string
	Port      string
}

const (
	connectWait        = time.Second * 30
	pubAckWait         = time.Second * 30
	interval           = 10
	maxOut             = 5
	maxPubAcksInflight = 25
)

func NewNatsConnect(conf NatsConfig) (stan.Conn, error) {
	return stan.Connect(
		conf.ClusterId,
		conf.ClientId,
		stan.ConnectWait(connectWait),
		stan.NatsURL("nats://"+conf.Host+":"+conf.Port),
		stan.PubAckWait(pubAckWait),
		stan.Pings(interval, maxOut),
		stan.MaxPubAcksInflight(maxPubAcksInflight),
	)
}

type NatsHandler struct {
	conn     stan.Conn
	conf     NatsConfig
	services *service.Service
}

func NewNatsHandler(conn stan.Conn, conf NatsConfig, services *service.Service) *NatsHandler {
	return &NatsHandler{conn: conn, conf: conf, services: services}
}

func (h *NatsHandler) InitSubs() (stan.Subscription, error) {
	return h.conn.QueueSubscribe(h.conf.ClientId, h.conf.ClusterId, h.CreateOrder)
}
