package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-demo/internal/proto/protouser"
)

type UserUsecase struct {
	user protouser.IUserRepo
	log  *log.Helper
}

func NewUserUsecase(user protouser.IUserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{user: user, log: log.NewHelper(logger)}
}

func (u *UserUsecase) Add(ctx context.Context, user *protouser.User) (bool, error) {
	return u.user.Insert(ctx, user)
}

func (u *UserUsecase) Upd(ctx context.Context, user *protouser.User) (bool, error) {
	return u.user.Update(ctx, user)
}

func (u *UserUsecase) GetById(ctx context.Context, id int64) (*protouser.User, error) {
	return u.user.SelectById(ctx, id)
}

func (u *UserUsecase) DelById(ctx context.Context, id int64) (bool, error) {
	return u.user.DeleteById(ctx, id)
}

func (u *UserUsecase) GetByPage(ctx context.Context, username string, curPage int64) ([]*protouser.User, int64, int64, error) {
	return u.user.SelectByPage(ctx, username, curPage)
}
