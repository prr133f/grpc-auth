package rpc

import (
	"context"
	"time"

	"auth/libs/jwt"
	pb "auth/proto"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponce, error) {
	user, err := s.PG.GetUserByEmail(in.GetEmail())
	if err != nil {
		s.Log.Error("error while getting user",
			zap.Error(err))
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pwdhash), []byte(in.GetPassword())); err != nil {
		s.Log.Error("error while comparing pwd and hash",
			zap.Error(err))
		return nil, err
	}

	accessToken, err := jwt.Generate(jwt.Payload{
		ID:        user.ID,
		Email:     user.Email,
		ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		Role:      string(user.Role),
	})
	if err != nil {
		s.Log.Error("error while generating access token",
			zap.Error(err))
		return nil, err
	}

	refreshToken, err := jwt.Generate(jwt.Payload{
		ID:        user.ID,
		Email:     user.Email,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
		Role:      string(user.Role),
	})
	if err != nil {
		s.Log.Error("error while generating refresh token",
			zap.Error(err))
		return nil, err
	}

	return &pb.LoginResponce{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Server) UpdateSession(ctx context.Context, in *pb.RefreshToken) (*pb.AccessToken, error) {
	accessToken, err := jwt.ReissueAccessToken(in.GetToken())
	if err != nil {
		s.Log.Error("error while reissue access token", zap.Error(err))
		return nil, err
	}

	return &pb.AccessToken{
		Token: accessToken,
	}, nil
}
