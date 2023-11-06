package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-demo/internal/proto/protouser"
)

var (
	pageSize = 5
)

type User struct {
	ID       int64
	Username string
	Password string
}

func (User) TableName() string {
	return "user"
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) *UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UserRepo) Insert(ctx context.Context, user *protouser.User) (bool, error) {
	u := User{
		Username: user.Username,
		Password: user.Password,
	}
	result := r.data.db.WithContext(ctx).Create(&u)
	if result.Error != nil || result.RowsAffected <= 0 {
		r.log.WithContext(ctx).Errorf("插入数据失败,原因:%v", result.Error)
		return false, result.Error
	}
	return true, nil
}

func (r *UserRepo) Update(ctx context.Context, user *protouser.User) (bool, error) {
	result := r.data.db.WithContext(ctx).Model(&User{}).Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"username": user.Username,
			"password": user.Password,
		})
	if result.Error != nil || result.RowsAffected <= 0 {
		r.log.WithContext(ctx).Errorf("更新数据失败,原因:%v", result.Error)
		return false, result.Error
	}
	return true, nil
}

func (r *UserRepo) SelectById(ctx context.Context, id int64) (*protouser.User, error) {
	var u User
	err := r.data.db.WithContext(ctx).First(&u, id).Error
	if err != nil {
		r.log.WithContext(ctx).Errorf("通过id查询数据失败,原因:%v", err)
		return nil, err
	}
	return &protouser.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}, nil
}

func (r *UserRepo) DeleteById(ctx context.Context, id int64) (bool, error) {
	result := r.data.db.WithContext(ctx).Delete(&User{}, id)
	if result.Error != nil || result.RowsAffected <= 0 {
		r.log.WithContext(ctx).Errorf("删除数据失败,原因:%v", result.Error)
		return false, result.Error
	}
	return true, nil
}

func (r *UserRepo) SelectByPage(ctx context.Context, username string, curPage int64) ([]*protouser.User, int64, int64, error) {
	var total int64
	// 获取总数据行
	tx := r.data.db.WithContext(ctx).Model(&User{}).Where("username like ?", "%"+username+"%")
	err := tx.Count(&total).Error
	if err != nil {
		r.log.WithContext(ctx).Errorf("获取数据总行数失败,原因:%v", err)
		return nil, total, curPage, err
	}
	// 获取总页数
	totalPage := total / int64(pageSize)
	if total%int64(pageSize) != 0 {
		totalPage++
	}
	// 越界判断
	if curPage < 1 {
		curPage = 1
	}
	if curPage > totalPage {
		curPage = totalPage
	}

	// 分页查询数据
	var users []*User
	err = r.data.db.WithContext(ctx).
		Limit(pageSize).Offset((int(curPage) - 1) * pageSize).
		Find(&users).Error
	if err != nil {
		r.log.WithContext(ctx).Errorf("查询分页数据失败,原因:%v", err)
		return nil, total, curPage, err
	}

	resp := make([]*protouser.User, 0, len(users))
	for _, user := range users {
		resp = append(resp, &protouser.User{
			ID:       user.ID,
			Username: user.Username,
			Password: user.Password,
		})
	}

	return resp, total, curPage, nil
}
