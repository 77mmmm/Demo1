package user

import (
	"awesomeProject/lib/ctx"
)

type RegisterUser struct {
	Ctx *ctx.GlobalCtx
}

func (u *RegisterUser) Register(username string, password string, email string) (bool, error) {
	Bool, err := ctx.GetGlobalCtx().User.SelectUser(username, password, email)
	if err != nil {
		return Bool, err
	}

	return Bool, nil

}
