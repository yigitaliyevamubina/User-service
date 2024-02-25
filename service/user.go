package service

import (
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes/empty"
	pb "user-service/genproto/user-service"
	"user-service/pkg/logger"
	grpcclient "user-service/service/grpc_client"
	"user-service/storage"
)

//User-service

type UserService struct {
	storage storage.IStorage
	logger  logger.Logger
	client  grpcclient.IServiceManager
	pb.UnimplementedUserServiceServer
}

func NewUserService(db *sql.DB, log logger.Logger, client grpcclient.IServiceManager) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

// rpc CreateUser(User) returns (User) {}
// rpc UpdateUser(User) returns (User) {}
// rpc GetUserById(GetUserReqById) returns (User) {}
// rpc GetUserByEmail(GetUserByEmailReq) returns (GetUserByEmailResp) {}
// rpc DeleteUser(DeleteUserReq) returns (google.protobuf.Empty) {}
// rpc CheckEmail(CheckEmailReq) returns (CheckEmailResp) {}
// rpc CheckField(CheckEmailReq) returns (CheckEmailResp) {}
// rpc GetAllUsers(ListUsersReq) returns (ListUsersResp) {}
// rpc Login(LoginReq) returns (LoginResp) {}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return s.storage.User().CreateUser(req)
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return s.storage.User().UpdateUser(req)
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserReqById) (*pb.User, error) {
	return s.storage.User().GetUserById(req)
}

func (s *UserService) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailReq) (*pb.GetUserByEmailResp, error) {
	return s.storage.User().GetUserByEmail(req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*empty.Empty, error) {
	return s.storage.User().DeleteUser(req)
}

func (s *UserService) CheckEmail(ctx context.Context, req *pb.CheckEmailReq) (*pb.CheckEmailResp, error) {
	return s.storage.User().CheckEmail(req)
}

func (s *UserService) CheckField(ctx context.Context, req *pb.CheckFieldReq) (*pb.CheckFieldResp, error) {
	return s.storage.User().CheckField(req)
}

func (s *UserService) GetAllUsers(ctx context.Context, req *pb.ListUsersReq) (*pb.ListUsersResp, error) {
	return s.storage.User().GetAllUsers(req)
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	return s.storage.User().Login(req)
}

func (s *UserService) mustEmbedUnimplementedUserServiceServer() {}
