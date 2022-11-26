package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassgetgroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClassgetgroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassgetgroupLogic {
	return &ClassgetgroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClassgetgroupLogic) Classgetgroup(req *types.ClassGroupreq) (resp *types.ClassGroupres, err error) {
	// todo: add your logic here and delete this line

	return
}
