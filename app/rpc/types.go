package rpc

import (
	"auth/app/models"
	pb "auth/proto"

	"github.com/rs/zerolog"
)

type Server struct {
	pb.UnimplementedAuthServer
	PG  RPCIface
	Log *zerolog.Logger
}

type RPCIface interface {
	GetUserByEmail(email string) (models.User, error)
}
