package socketio

import "time"

type Config struct {
	Network      string
	Address      string
	Path         string
	PingInterval time.Duration
	PingTimeout  time.Duration
}

type Option func(*Config)

func WithNetwork(network string) Option {
	return func(c *Config) {
		c.Network = network
	}
}

func WithAddress(addr string) Option {
	return func(c *Config) {
		c.Address = addr
	}
}

func WithPath(path string) Option {
	return func(c *Config) {
		c.Path = path
	}
}

func WithPingInterval(d time.Duration) Option {
	return func(c *Config) {
		c.PingInterval = d
	}
}

func WithPingTimeout(d time.Duration) Option {
	return func(c *Config) {
		c.PingTimeout = d
	}
}
