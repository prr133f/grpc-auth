package rpc

import (
	pb "auth/proto"

	"github.com/rs/zerolog"
)

type Server struct {
	pb.UnimplementedAuthServer
	PG  RPCIface
	Log *zerolog.Logger
}

type RPCIface interface {
}
