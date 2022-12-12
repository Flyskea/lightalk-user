package logic

import (
	"context"

	"github.com/flyskea/lightalk-user-rpc/internal/svc"
	"github.com/flyskea/lightalk-user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IDReq) (*user.UserDetail, error) {
	// todo: add your logic here and delete this line

	return &user.UserDetail{}, nil
}
