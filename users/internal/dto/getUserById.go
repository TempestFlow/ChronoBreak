package dto

import (
	z "github.com/Oudwins/zog"
	"github.com/mhirii/chronobreak/users/pkg/conv"
)

type GetUserByIDRequest struct {
	ID string `json:"id"`
}

var GetUserByIDRequestSchema = z.Struct(z.Schema{
	"id": z.String().Min(1).Max(128),
})

func (req *GetUserByIDRequest) Parse(request any) error {
	errsMap := GetUserByIDRequestSchema.Parse(request, &req)
	if errsMap != nil {
		return conv.MapZogErrs(errsMap)
	}
	return nil
}

func (req *GetUserByIDRequest) Validate() error {
	err := GetUserByIDRequestSchema.Validate(req)
	if err != nil {
		return conv.MapZogErrs(err)
	}
	return nil
}

type GetUserByIDResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
