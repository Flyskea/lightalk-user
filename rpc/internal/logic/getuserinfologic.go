package logic

import (
	"context"

	"github.com/flyskea/lightalk-user/rpc/internal/svc"
	"github.com/flyskea/lightalk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.IDReq) (*user.UserInfo, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfo{}, nil
}
