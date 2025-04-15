package v1

// versions:
// 	gofr-cli v0.6.0
// 	gofr.dev v1.37.0
// 	source: users.proto

import (
	"github.com/mhirii/chronobreak/users/internal/dto"
	"github.com/mhirii/chronobreak/users/internal/service"
	"gofr.dev/pkg/gofr"
)

type AuthServiceGoFrServer struct {
	health  *healthServer
	service service.AuthService
}

func (s *AuthServiceGoFrServer) Login(ctx *gofr.Context) (any, error) {
	span := ctx.Trace("grpc.AuthService.Login")
	defer span.End()

	request := LoginRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		return nil, err
	}

	r := dto.LoginRequest{
		Email:    request.GetEmail(),
		Password: request.GetPassword(),
	}

	err = r.Validate()
	if err != nil {
		ctx.Error(err)
		return nil, err
	}

	res, err := s.service.Login(ctx, r)
	if err != nil {
		ctx.Error(err)
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}

func (s *AuthServiceGoFrServer) Signup(ctx *gofr.Context) (any, error) {
	span := ctx.Trace("grpc.AuthService.Signup")
	defer span.End()

	request := SignupRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		return nil, err
	}

	r := dto.SignupRequest{
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
	}

	err = r.Validate()
	if err != nil {
		ctx.Error(err)
		return nil, err
	}

	res, err := s.service.Signup(ctx, r)
	if err != nil {
		ctx.Error(err)
		return nil, err
	}

	return &SignupResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}

func (s *AuthServiceGoFrServer) RefreshToken(ctx *gofr.Context) (any, error) {
	span := ctx.Trace("grpc.AuthService.RefreshToken")
	defer span.End()

	request := RefreshTokenRequest{}
	err := ctx.Bind(&request)
	if err != nil || request.RefreshToken == "" {
		ctx.Error(err)
		return nil, err
	}

	r := dto.RefreshTokenRequest{RefreshToken: request.RefreshToken}

	res, err := s.service.RefreshToken(ctx, r)
	if err != nil {
		ctx.Error(err)
		return nil, err
	}

	return &RefreshTokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}

func (s *AuthServiceGoFrServer) GetUser(ctx *gofr.Context) (any, error) {
	span := ctx.Trace("grpc.AuthService.GetUser")
	defer span.End()
	return &GetUserResponse{}, nil
}

func (s *AuthServiceGoFrServer) GetUserByID(ctx *gofr.Context) (any, error) {
	span := ctx.Trace("grpc.AuthService.GetUserByID")
	defer span.End()

	return &GetUserByIDResponse{}, nil
}

func (s *AuthServiceGoFrServer) Logout(ctx *gofr.Context) (any, error) {
	span := ctx.Trace("grpc.AuthService.Logout")
	defer span.End()

	return &LogoutResponse{}, nil
}

func (s *AuthServiceGoFrServer) Validate(ctx *gofr.Context) (any, error) {
	span := ctx.Trace("grpc.AuthService.Validate")
	defer span.End()

	return &ValidateResponse{}, nil
}
