package dto

import (
	z "github.com/Oudwins/zog"
	"github.com/mhirii/chronobreak/users/pkg/conv"
)

type GetUserRequest struct {
	AccessToken string `json:"access_token"`
}

var GetUserRequestSchema = z.Struct(z.Schema{
	"access_token": z.String().Min(1).Max(128),
})

func (req *GetUserRequest) Parse(request any) error {
	errsMap := GetUserRequestSchema.Parse(request, &req)
	if errsMap != nil {
		return conv.MapZogErrs(errsMap)
	}
	return nil
}

func (req *GetUserRequest) Validate() error {
	err := GetUserRequestSchema.Validate(req)
	if err != nil {
		return conv.MapZogErrs(err)
	}
	return nil
}

type GetUserResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
