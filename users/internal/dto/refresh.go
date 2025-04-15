package dto

import (
	z "github.com/Oudwins/zog"
	"github.com/mhirii/chronobreak/users/pkg/conv"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

var RefreshTokenRequestSchema = z.Struct(z.Schema{
	"refresh_token": z.String().Min(1).Max(128),
})

func (req *RefreshTokenRequest) Parse(request any) error {
	errsMap := RefreshTokenRequestSchema.Parse(request, &req)
	if errsMap != nil {
		return conv.MapZogErrs(errsMap)
	}
	return nil
}

func (req *RefreshTokenRequest) Validate() error {
	err := RefreshTokenRequestSchema.Validate(req)
	if err != nil {
		return conv.MapZogErrs(err)
	}
	return nil
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
