package socketio

import (
	socketLog "github.com/zishang520/engine.io/v2/log"
	"github.com/zishang520/engine.io/v2/types"
	"time"
)

type Options func(server *Server)

func WithAddress(addr string) Options {
	return func(s *Server) {
		s.address = addr
	}
}

func WithPath(path string) Options {
	return func(s *Server) {
		s.ServerOptions.SetPath(path)
	}
}

func WithPingInterval(d time.Duration) Options {
	return func(s *Server) {
		s.SetPingInterval(d)
	}
}

func WithPingTimeout(d time.Duration) Options {
	return func(s *Server) {
		s.SetPingTimeout(d)
	}
}

func WithAllowUpgrades(a bool) Options {
	return func(s *Server) {
		s.SetAllowUpgrades(a)
	}
}

func WithUpgradeTimeout(d time.Duration) Options {
	return func(s *Server) {
		s.SetUpgradeTimeout(d)
	}
}

func WithCors(c *types.Cors) Options {
	return func(s *Server) {
		s.SetCors(c)
	}
}

func WithAllowEIO3(b bool) Options {
	return func(s *Server) {
		s.SetAllowEIO3(b)
	}
}

func WithServerClient(b bool) Options {
	return func(s *Server) {
		s.Server.SetServeClient(b)
	}
}

func WithDebug(b bool) Options {
	return func(s *Server) {
		socketLog.DEBUG = b
	}
}
