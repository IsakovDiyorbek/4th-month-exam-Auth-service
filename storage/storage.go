package storage

import (
	"context"

	pb "github.com/Exam4/4th-month-exam-Auth-service/genproto/auth"
	pu "github.com/Exam4/4th-month-exam-Auth-service/genproto/user"
)

type StorageI interface {
	Auth() Auth
	User() User
}

type Auth interface {
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error)
	ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error)
}

type User interface {
	GetProfile(ctx context.Context, req *pu.GetProfileRequest) (*pu.GetProfileResponse, error)
	UpdateProfile(ctx context.Context, req *pu.UpdateProfileRequest) (*pu.UpdateProfileResponse, error)
	ChangePassword(ctx context.Context, req *pu.ChangePasswordRequest) (*pu.ChangePasswordResponse, error)

}
