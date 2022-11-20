package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CleanchangeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCleanchangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanchangeLogic {
	return &CleanchangeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CleanchangeLogic) Cleanchange(req *types.Cleanchangereq) (resp *types.Cleanchangeres, err error) {
	// todo: add your logic here and delete this line

	return
}
