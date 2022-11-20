package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsgcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgcLogic {
	return &MsgcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MsgcLogic) Msgc(req *types.MsgCreq) (resp *types.MsgCres, err error) {
	// todo: add your logic here and delete this line

	return
}
