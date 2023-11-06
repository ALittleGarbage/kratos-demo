package protouser

import (
	"context"
)

type IUserRepo interface {
	Insert(context.Context, *User) (bool, error)
	Update(context.Context, *User) (bool, error)
	SelectById(context.Context, int64) (*User, error)
	DeleteById(context.Context, int64) (bool, error)
	SelectByPage(context.Context, string, int64) ([]*User, int64, int64, error)
}

type User struct {
	ID       int64
	Username string
	Password string
}
