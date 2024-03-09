package rpc

import (
	"auth/app/models"
	pb "auth/proto"

	"go.uber.org/zap"
)

type Server struct {
	pb.UnimplementedAuthServer
	PG  RPCIface
	Log *zap.Logger
}

type RPCIface interface {
	GetUserByEmail(email string) (models.User, error)
}
