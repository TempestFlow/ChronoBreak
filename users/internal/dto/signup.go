package dto

import (
	z "github.com/Oudwins/zog"
	"github.com/mhirii/chronobreak/users/pkg/conv"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var SignupRequestSchema = z.Struct(z.Schema{
	"email":    z.String().Email(),
	"username": z.String().Min(3).Max(32),
	"password": z.String().Min(8).Max(128),
})

func (req *SignupRequest) Parse(request any) error {
	errsMap := SignupRequestSchema.Parse(request, &req)
	if errsMap != nil {
		return conv.MapZogErrs(errsMap)
	}
	return nil
}

func (req *SignupRequest) Validate() error {
	err := SignupRequestSchema.Validate(req)
	if err != nil {
		return conv.MapZogErrs(err)
	}
	return nil
}

type SignupResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
