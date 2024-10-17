package repository

import (
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"context"
)

type UserRepo interface {
	CheckLogin(context context.Context, LoginReq req.ReqSignIn) (model.User, error)
	SaveUser(context context.Context, user model.User) (model.User, error)
}
