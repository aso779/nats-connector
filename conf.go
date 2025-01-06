package nats_connector

import (
	"time"

	cfgloader "github.com/aso779/config-loader"
)

type Nats struct {
	AddrProp         string `yaml:"addr"`
	LoginProp        string `yaml:"login"`
	PasswordProp     string `yaml:"password"`
	RetryTimeoutProp string `yaml:"retry_timeout"`
	MaxReconnectProp string `yaml:"max_reconnect"`
}

func (r Nats) Addr() string {
	return cfgloader.LoadStringProp(r.AddrProp)
}

func (r Nats) Login() string {
	return cfgloader.LoadStringProp(r.LoginProp)
}

func (r Nats) Password() string {
	return cfgloader.LoadStringProp(r.PasswordProp)
}

func (r Nats) RetryTimeout() time.Duration {
	return cfgloader.LoadDurationProp(r.RetryTimeoutProp)
}

func (r Nats) MaxReconnect() int {
	return cfgloader.LoadIntProp(r.MaxReconnectProp)
}
