package natsconnector

import (
	"errors"
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var ErrNatsCantConn = errors.New("nats: can't connect")

type NatsConnSet interface {
	Close()
	Conn() (*nats.Conn, error)
}

type NatsConn struct {
	conf *Nats
	log  *zap.Logger
	conn *nats.Conn
}

func NewNatsConn(conf *Nats, log *zap.Logger) NatsConnSet {
	return &NatsConn{
		conf: conf,
		log:  log,
	}
}

func (r *NatsConn) Conn() (*nats.Conn, error) {
	if r.conn == nil || !r.conn.IsConnected() {
		var err error

		for i := 0; ; i++ {
			if i > r.conf.MaxReconnect() {
				r.log.Error("nats: can't connect")
				return nil, ErrNatsCantConn
			}

			r.log.Info("nats: trying to connect")

			r.conn, err = r.connect()
			if err == nil {
				break
			}

			r.log.Error("nats: connection error", zap.Error(err))

			time.Sleep(r.conf.RetryTimeout() * time.Second)
		}
	}

	return r.conn, nil
}

func (r *NatsConn) Close() {
	if r.conn != nil {
		r.conn.Close()
		r.log.Info("graceful shutdown nats connection")
	}
}

func (r *NatsConn) connect() (*nats.Conn, error) {
	return nats.Connect(r.conf.Addr(),
		nats.ConnectHandler(func(_ *nats.Conn) {
			r.log.Info("nats: connected")
		}),
		nats.DisconnectErrHandler(func(_ *nats.Conn, _ error) {
			r.log.Error("nats: disconnected")
		}),
		nats.UserInfo(r.conf.Login(), r.conf.Password()),
	)
}
