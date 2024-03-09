package rpc

import (
	"context"
	"time"

	jwt "auth/libs/jwt"
	pb "auth/proto"

	"golang.org/x/crypto/bcrypt"
)

func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponce, error) {
	user, err := s.PG.GetUserByEmail(in.GetEmail())
	if err != nil {
		s.Log.Error().Err(err).Stack()
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pwdhash), []byte(in.GetPassword())); err != nil {
		s.Log.Error().Err(err).Stack()
		return nil, err
	}

	accessToken, err := jwt.Generate(jwt.Payload{
		ID:        user.ID,
		Email:     user.Email,
		ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		Role:      string(user.Role),
	})
	if err != nil {
		s.Log.Error().Err(err).Stack()
		return nil, err
	}

	refreshToken, err := jwt.Generate(jwt.Payload{
		ID:        user.ID,
		Email:     user.Email,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
		Role:      string(user.Role),
	})
	if err != nil {
		s.Log.Error().Err(err).Stack()
		return nil, err
	}

	return &pb.LoginResponce{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
