package logic

import (
	"context"

	"github.com/flyskea/lightalk-user/rpc/internal/svc"
	"github.com/flyskea/lightalk-user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.UserResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserResp{}, nil
}
