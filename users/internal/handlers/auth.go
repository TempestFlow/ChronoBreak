package handlers

import (
	"github.com/mhirii/chronobreak/users/internal/dto"
	"github.com/mhirii/chronobreak/users/internal/service"
	"gofr.dev/pkg/gofr"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler() AuthHandler {
	return AuthHandler{}
}

func (h AuthHandler) Login(ctx *gofr.Context) (any, error) {
	req := dto.LoginRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		ctx.Error(err)
		return nil, err
	}

	err = req.Validate()
	if err != nil {
		ctx.Error(err)
		return nil, err
	}
	res, err := h.service.Login(ctx, req)
	if err != nil {
		ctx.Error(err)
		return nil, err
	}

	return res, nil
}

func (h AuthHandler) Signup(ctx *gofr.Context) (any, error) {
	return &dto.SignupResponse{}, nil
}

func (h AuthHandler) RefreshToken(ctx *gofr.Context) (any, error) {
	return &dto.RefreshTokenResponse{}, nil
}

func (h AuthHandler) GetUser(ctx *gofr.Context) (any, error) {
	return &dto.GetUserResponse{}, nil
}

func (h AuthHandler) GetUserByID(ctx *gofr.Context) (any, error) {
	return &dto.GetUserByIDResponse{}, nil
}

func (h AuthHandler) Logout(ctx *gofr.Context) (any, error) {
	return &dto.LogoutResponse{}, nil
}
