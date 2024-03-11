package repo

import (
	pb "user-service/genproto/user-service"

	"github.com/golang/protobuf/ptypes/empty"
)

// UserStorageI
type UserStorageI interface {
	CreateUser(*pb.User) (*pb.User, error)
	UpdateUser(*pb.User) (*pb.User, error)
	GetUserById(*pb.GetUserReqById) (*pb.User, error)
	DeleteUser(*pb.DeleteUserReq) (*empty.Empty, error)
	CheckField(*pb.CheckFieldReq) (*pb.CheckFieldResp, error)
	GetAllUsers(*pb.ListUsersReq) (*pb.ListUsersResp, error)
	IfExists(*pb.IfExistsReq) (*pb.IfExistsResp, error)
	ChangePassword(*pb.ChangeUserPasswordReq) (*pb.ChangeUserPasswordResp, error)
	UpdateRefreshToken(*pb.UpdateRefreshTokenReq) (*pb.UpdateRefreshTokenResp, error)
}
