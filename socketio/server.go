package socketio

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	socketio "github.com/zishang520/socket.io/v2/socket"
	"net/http"
)

var _ transport.Server = (*Server)(nil)

type Server struct {
	*socketio.Server
	*socketio.ServerOptions

	address string
	debug   bool

	err error
}

func NewServer(opts ...Options) *Server {
	server := socketio.NewServer(nil, nil)

	srv := &Server{
		Server:        server,
		ServerOptions: socketio.DefaultServerOptions(),
	}

	for _, o := range opts {
		o(srv)
	}

	return srv
}

func (s *Server) Start(_ context.Context) error {
	http.Handle(s.ServerOptions.Path(), s.ServeHandler(s.ServerOptions))

	log.Infof("[socketIo] server listening on: %s", s.address)

	go func() {
		if err := http.ListenAndServe(s.address, nil); err != nil {
			log.Infof("[socketIo] server listening on: %s failed", s.address)
			panic(err)
		}
	}()

	return nil
}

func (s *Server) Stop(_ context.Context) error {
	log.Info("[socketIo] server stopping")
	s.Close(nil)
	return nil
}

func (s *Server) RegisterConnectHandler(clients func(clients ...interface{})) error {
	return s.On("connection", clients)
}
