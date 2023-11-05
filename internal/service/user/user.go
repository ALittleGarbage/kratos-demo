package user

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-demo/internal/biz/user"

	pb "kratos-demo/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *user.UserUsecase
	l  log.Logger
}

func NewUserService(uc *user.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, l: logger}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	u, err := s.uc.GetById(ctx, req.GetId())
	if err != nil {
		log.NewHelper(s.l).Errorf("通关id查询user失败,原因:%v", err)
		return nil, err
	}
	return &pb.GetUserReply{
		User: &pb.UserDto{
			Id:       u.ID,
			Username: u.Username,
		},
	}, nil
}
func (s *UserService) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersReply, error) {
	users, total, curPage, err := s.uc.GetByPage(ctx, req.GetUsername(), req.GetCurPage())
	if err != nil {
		log.NewHelper(s.l).Errorf("分页查询users失败,原因:%v", err)
		return nil, err
	}
	userDtos := make([]*pb.UserDto, 0, len(users))
	for _, u := range users {
		userDtos = append(userDtos, &pb.UserDto{
			Id:       u.ID,
			Username: u.Username,
		})
	}

	return &pb.GetUsersReply{
		Users:   userDtos,
		CurPage: curPage,
		Total:   total,
	}, nil
}
func (s *UserService) RemoveUser(ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserReply, error) {
	res, err := s.uc.DelById(ctx, req.GetId())
	if err != nil {
		log.NewHelper(s.l).Errorf("通过id删除user失败,原因:%v", err)
		return nil, err
	}
	return &pb.RemoveUserReply{Res: res}, nil
}
func (s *UserService) SaveUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserReply, error) {
	userVo := req.GetUser()
	if userVo == nil {
		log.NewHelper(s.l).Errorf("保存user失败,原因:userInfo为空")
		return nil, errors.New("user info is empty")
	}
	res, err := s.uc.Add(ctx, &user.User{
		ID:       userVo.GetId(),
		Username: userVo.GetUsername(),
		Password: userVo.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.AddUserReply{Res: res}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	userVo := req.GetUser()
	res, err := s.uc.Upd(ctx, &user.User{
		ID:       userVo.GetId(),
		Username: userVo.GetUsername(),
		Password: userVo.GetPassword(),
	})
	if err != nil {
		log.NewHelper(s.l).Errorf("更新user失败,原因:%v", err)
		return nil, err
	}
	return &pb.UpdateUserReply{Res: res}, nil
}
