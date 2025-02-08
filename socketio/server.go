package socketio

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/v2/transport"
	socketio "github.com/zishang520/socket.io/v2/socket"
)

var _ transport.Server = (*Server)(nil)

type Server struct {
	*socketio.Server
	conf     *Config
	address  string
	path     string
	opts     []Option
	startCtx context.Context
}

func NewServer(opts ...Option) *Server {

	server := socketio.NewServer(nil, nil)

	srv := &Server{
		Server:  server,
		path:    "/socket.io/",
		address: ":3000",
	}

	for _, o := range opts {
		o(srv.conf)
	}

	return srv
}

func (s *Server) Start(ctx context.Context) error {
	s.startCtx = ctx

	http.Handle(s.path, s.Server.ServeHandler(nil))
	go func() {
		if err := http.ListenAndServe(s.address, nil); err != nil {
			panic(err)
		}
	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.Server.Close(nil)
	return nil
}
