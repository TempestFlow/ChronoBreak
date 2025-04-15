package dto

import (
	z "github.com/Oudwins/zog"
	"github.com/mhirii/chronobreak/users/pkg/conv"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var LoginRequestSchema = z.Struct(z.Schema{
	"email":    z.String().Email(),
	"password": z.String().Min(8).Max(128),
})

func (req *LoginRequest) Parse(request any) error {
	var res LoginRequest
	errsMap := LoginRequestSchema.Parse(request, &res)
	if errsMap != nil {
		return conv.MapZogErrs(errsMap)
	}
	return nil
}

func (req *LoginRequest) Validate() error {
	err := LoginRequestSchema.Validate(req)
	if err != nil {
		return conv.MapZogErrs(err)
	}
	return nil
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
