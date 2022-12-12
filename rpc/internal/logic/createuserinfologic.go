package logic

import (
	"context"

	"github.com/flyskea/lightalk-user-rpc/internal/svc"
	"github.com/flyskea/lightalk-user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserInfoLogic {
	return &CreateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserInfoLogic) CreateUserInfo(in *user.UserInfo) (*user.UserInfo, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfo{}, nil
}
