package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-demo/internal/biz/user"
)

var (
	pageSize = 5
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) user.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Insert(ctx context.Context, user *user.User) (bool, error) {
	result := r.data.db.Create(&user)
	if result.Error != nil || result.RowsAffected <= 0 {
		r.log.WithContext(ctx).Errorf("插入数据失败,原因:%v", result.Error)
		return false, result.Error
	}
	return true, nil
}

func (r *userRepo) Update(ctx context.Context, user *user.User) (bool, error) {
	result := r.data.db.Updates(&user)
	if result.Error != nil || result.RowsAffected <= 0 {
		r.log.WithContext(ctx).Errorf("更新数据失败,原因:%v", result.Error)
		return false, result.Error
	}
	return true, nil
}

func (r *userRepo) SelectById(ctx context.Context, id int64) (*user.User, error) {
	user := new(user.User)
	err := r.data.db.First(&user, id).Error
	if err != nil {
		r.log.WithContext(ctx).Errorf("通过id查询数据失败,原因:%v", err)
		return nil, err
	}
	return user, nil
}

func (r *userRepo) DeleteById(ctx context.Context, id int64) (bool, error) {
	result := r.data.db.Delete(&user.User{}, id)
	if result.Error != nil || result.RowsAffected <= 0 {
		r.log.WithContext(ctx).Errorf("删除数据失败,原因:%v", result.Error)
		return false, result.Error
	}
	return true, nil
}

func (r *userRepo) SelectByPage(ctx context.Context, username string, curPage int64) ([]*user.User, int64, int64, error) {
	var total int64
	// 获取总数据行
	err := r.data.db.Model(&user.User{}).Where("username like ?", "%"+username+"%").Count(&total).Error
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
	var users []*user.User
	err = r.data.db.Limit(pageSize).Offset((int(curPage)-1)*pageSize).Where("username like ?", "%"+username+"%").Find(&users).Error
	if err != nil {
		r.log.WithContext(ctx).Errorf("查询分页数据失败,原因:%v", err)
		return nil, total, curPage, err
	}
	return users, total, curPage, nil
}
