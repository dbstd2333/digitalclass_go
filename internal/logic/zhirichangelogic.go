package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ZhirichangeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZhirichangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZhirichangeLogic {
	return &ZhirichangeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ZhirichangeLogic) Zhirichange(req *types.Zhirichangereq) (resp *types.Zhirichangeres, err error) {
	// todo: add your logic here and delete this line

	return
}
