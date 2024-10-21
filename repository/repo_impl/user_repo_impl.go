package repo_impl

import (
	"backend-github-trending/db"
	"backend-github-trending/lemon"
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"backend-github-trending/repository"
	"context"
	"database/sql"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}

func (u UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement :=
		` INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)
	 VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)`

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if _, err := u.sql.Db.NamedExecContext(context, statement, user); err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, lemon.UserConflict
			}
		}
		return user, lemon.SignUpFail
	}
	return user, nil
}

func (u *UserRepoImpl) CheckLogin(context context.Context, LoginReq req.ReqSignIn) (model.User, error) {
	var user = model.User{}
	statement := `SELECT * FROM users WHERE email =$1`
	err := u.sql.Db.GetContext(context, &user, statement, LoginReq.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, lemon.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}
	return user, nil
}

func (u *UserRepoImpl) SelectUserById(context context.Context, userId string) (model.User, error) {
	var user model.User
	err := u.sql.Db.GetContext(context, &user, `SELECT * FROM users WHERE user_id =$1`, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, lemon.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}
	return user, nil
}
