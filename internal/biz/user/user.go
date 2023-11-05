package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID       int64
	Username string
	Password string
}

func (User) TableName() string {
	return "user"
}

type UserRepo interface {
	Insert(context.Context, *User) (bool, error)
	Update(context.Context, *User) (bool, error)
	SelectById(context.Context, int64) (*User, error)
	DeleteById(context.Context, int64) (bool, error)
	SelectByPage(context.Context, string, int64) ([]*User, int64, int64, error)
}

type UserUsecase struct {
	user UserRepo
	log  *log.Helper
}

func NewUserUsecase(user UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{user: user, log: log.NewHelper(logger)}
}

func (u *UserUsecase) Add(ctx context.Context, user *User) (bool, error) {
	return u.user.Insert(ctx, user)
}

func (u *UserUsecase) Upd(ctx context.Context, user *User) (bool, error) {
	return u.user.Update(ctx, user)
}

func (u *UserUsecase) GetById(ctx context.Context, id int64) (*User, error) {
	return u.user.SelectById(ctx, id)
}

func (u *UserUsecase) DelById(ctx context.Context, id int64) (bool, error) {
	return u.user.DeleteById(ctx, id)
}

func (u *UserUsecase) GetByPage(ctx context.Context, username string, curPage int64) ([]*User, int64, int64, error) {
	return u.user.SelectByPage(ctx, username, curPage)
}
