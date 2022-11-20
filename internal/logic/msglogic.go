package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgLogic {
	return &MsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MsgLogic) Msg(req *types.Msgreq) (resp *types.Msgres, err error) {
	// todo: add your logic here and delete this line

	return
}
