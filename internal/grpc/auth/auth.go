package auth

import (
	"context"

	ssov1 "github.com/paniccaaa/protos/gen/golang/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	panic("unimplemented log")
}

// GetRole implements ssov1.AuthServer.
func (s *serverAPI) GetRole(ctx context.Context, req *ssov1.UserIdRequest) (*ssov1.RoleResponse, error) {
	panic("unimplemented")
}

// IsAdmin implements ssov1.AuthServer.
func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	panic("unimplemented")
}

// Logout implements ssov1.AuthServer.
func (s *serverAPI) Logout(ctx context.Context, req *ssov1.LogoutRequest) (*ssov1.LogoutResponse, error) {
	panic("unimplemented")
}

// Register implements ssov1.AuthServer.
func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	panic("unimplemented")
}
