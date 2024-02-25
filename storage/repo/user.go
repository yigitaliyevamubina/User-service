package repo

import (
	"github.com/golang/protobuf/ptypes/empty"
	pb "user-service/genproto/user-service"
)

// rpc CreateUser(User) returns (User) {}
// rpc UpdateUser(User) returns (User) {}
// rpc GetUserById(GetUserReqById) returns (User) {}
// rpc GetUserByEmail(GetUserByEmailReq) returns (GetUserByEmailResp) {}
// rpc DeleteUser(DeleteUserReq) returns (google.protobuf.Empty) {}
// rpc CheckEmail(CheckEmailReq) returns (CheckEmailResp) {}
// rpc CheckField(CheckEmailReq) returns (CheckEmailResp) {}
// rpc GetAllUsers(ListUsersReq) returns (ListUsersResp) {}
// rpc Login(LoginReq) returns (LoginResp) {}
// UserStorageI
type UserStorageI interface {
	CreateUser(*pb.User) (*pb.User, error)
	UpdateUser(*pb.User) (*pb.User, error)
	GetUserById(*pb.GetUserReqById) (*pb.User, error)
	GetUserByEmail(*pb.GetUserByEmailReq) (*pb.GetUserByEmailResp, error)
	DeleteUser(*pb.DeleteUserReq) (*empty.Empty, error)
	CheckEmail(*pb.CheckEmailReq) (*pb.CheckEmailResp, error)
	CheckField(*pb.CheckFieldReq) (*pb.CheckFieldResp, error)
	GetAllUsers(*pb.ListUsersReq) (*pb.ListUsersResp, error)
	Login(*pb.LoginReq) (*pb.LoginResp, error)
}
