package service

import (
	"context"
	"database/sql"
	pb "user-service/genproto/user-service"
	"user-service/pkg/logger"
	grpcclient "user-service/service/grpc_client"
	"user-service/storage"

	"github.com/golang/protobuf/ptypes/empty"
)

//User-service

type UserService struct {
	storage storage.IStorage
	logger  logger.Logger
	client  grpcclient.IServiceManager
}

func NewUserService(db *sql.DB, log logger.Logger, client grpcclient.IServiceManager) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().CreateUser(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().UpdateUser(req)

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserReqById) (*pb.User, error) {
	user, err := s.storage.User().GetUserById(req)
	
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*empty.Empty, error) {
	resp, err := s.storage.User().DeleteUser(req)

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (s *UserService) CheckField(ctx context.Context, req *pb.CheckFieldReq) (*pb.CheckFieldResp, error) {
	resp, err := s.storage.User().CheckField(req)

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (s *UserService) GetAllUsers(ctx context.Context, req *pb.ListUsersReq) (*pb.ListUsersResp, error) {
	users, err := s.storage.User().GetAllUsers(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return users, nil
}

func (s *UserService) IfExists(ctx context.Context, req *pb.IfExistsReq) (*pb.IfExistsResp, error) {
	resp, err := s.storage.User().IfExists(req)

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (s *UserService) ChangePassword(ctx context.Context, req *pb.ChangeUserPasswordReq) (*pb.ChangeUserPasswordResp, error) {
	resp, err := s.storage.User().ChangePassword(req)

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (s *UserService) UpdateRefreshToken(ctx context.Context, req *pb.UpdateRefreshTokenReq) (*pb.UpdateRefreshTokenResp, error) {
	resp, err := s.storage.User().UpdateRefreshToken(req)

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return resp, nil
}
