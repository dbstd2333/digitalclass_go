package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KebiaochangeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKebiaochangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KebiaochangeLogic {
	return &KebiaochangeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KebiaochangeLogic) Kebiaochange(req *types.Kebiaocreq) (resp *types.Kebiaocres, err error) {
	// todo: add your logic here and delete this line

	return
}
