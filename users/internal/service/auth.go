package service

import (
	"github.com/mhirii/chronobreak/users/internal/dto"
	"gofr.dev/pkg/gofr"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(ctx *gofr.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	span := ctx.Trace("service.Login")
	defer span.End()

	return &dto.LoginResponse{
		AccessToken:  "token",
		RefreshToken: "token",
	}, nil
}

func (s *AuthService) Signup(ctx *gofr.Context, req dto.SignupRequest) (*dto.SignupResponse, error) {
	span := ctx.Trace("service.Signup")
	defer span.End()

	return &dto.SignupResponse{
		AccessToken:  "token",
		RefreshToken: "token",
	}, nil
}

func (s *AuthService) RefreshToken(ctx *gofr.Context, req dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	span := ctx.Trace("service.RefreshToken")
	defer span.End()

	return &dto.RefreshTokenResponse{
		AccessToken:  "token",
		RefreshToken: "token",
	}, nil
}

func (s *AuthService) GetUser(ctx *gofr.Context, req dto.GetUserRequest) (*dto.GetUserResponse, error) {
	span := ctx.Trace("service.GetUser")
	defer span.End()

	return &dto.GetUserResponse{
		ID:          "id",
		Username:    "username",
		Email:       "email",
		AccessToken: "token",
	}, nil
}

func (s *AuthService) GetUserByID(ctx *gofr.Context, req dto.GetUserByIDRequest) (*dto.GetUserByIDResponse, error) {
	span := ctx.Trace("service.GetUserByID")
	defer span.End()

	return &dto.GetUserByIDResponse{
		ID:          "id",
		Username:    "username",
		Email:       "email",
		AccessToken: "token",
	}, nil
}

func (s *AuthService) Logout(ctx *gofr.Context, req dto.LogoutRequest) (*dto.LogoutResponse, error) {
	span := ctx.Trace("service.Logout")
	defer span.End()

	return &dto.LogoutResponse{
		AccessToken: "token",
	}, nil
}
