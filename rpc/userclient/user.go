// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"github.com/flyskea/lightalk-user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseUserInfo = user.BaseUserInfo
	IDReq        = user.IDReq
	LoginReq     = user.LoginReq
	PageReq      = user.PageReq
	RegisterReq  = user.RegisterReq
	UserDetail   = user.UserDetail
	UserInfo     = user.UserInfo
	UserListResp = user.UserListResp
	UserResp     = user.UserResp

	User interface {
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*UserResp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*UserDetail, error)
		GetUser(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*UserDetail, error)
		ListUser(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*UserListResp, error)
		CreateUserInfo(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserInfo, error)
		GetUserInfo(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*UserInfo, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*UserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*UserDetail, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) GetUser(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*UserDetail, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) ListUser(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.ListUser(ctx, in, opts...)
}

func (m *defaultUser) CreateUserInfo(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserInfo, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateUserInfo(ctx, in, opts...)
}

func (m *defaultUser) GetUserInfo(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*UserInfo, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}
