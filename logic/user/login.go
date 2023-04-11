package user

import (
	"awesomeProject/lib/ctx"
	"awesomeProject/lib/jwt"
	"errors"
)

type User struct {
	Ctx *ctx.GlobalCtx
}

func NewUser() *User {
	return &User{Ctx: ctx.GetGlobalCtx()}
}

func (u *User) Login(userName, password string) (string, error) {
	pass, err := ctx.GetGlobalCtx().User.GetUser(userName)
	if err != nil {
		return "", err
	}
	if pass != password {
		return "", errors.New("身份校验失败")
	}
	token, err := jwt.GenerateToken(userName)
	if err != nil {
		return "", errors.New("创建失败")
	}
	return token, nil
}
