package logic

import (
	"context"

	"github.com/flyskea/lightalk-user-rpc/rpc/internal/svc"
	"github.com/flyskea/lightalk-user-rpc/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListUserLogic) ListUser(in *user.PageReq) (*user.UserListResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserListResp{}, nil
}
