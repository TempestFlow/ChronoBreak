package dto

import (
	z "github.com/Oudwins/zog"
	"github.com/mhirii/chronobreak/users/pkg/conv"
)

type LogoutRequest struct {
	AccessToken string `json:"access_token"`
}

var LogoutRequestSchema = z.Struct(z.Schema{
	"access_token": z.String().Min(1).Max(128),
})

func (req *LogoutRequest) Parse(request any) error {
	errsMap := LogoutRequestSchema.Parse(request, &req)
	if errsMap != nil {
		return conv.MapZogErrs(errsMap)
	}
	return nil
}

func (req *LogoutRequest) Validate() error {
	err := LogoutRequestSchema.Validate(req)
	if err != nil {
		return conv.MapZogErrs(err)
	}
	return nil
}

type LogoutResponse struct {
	AccessToken string `json:"access_token"`
}
