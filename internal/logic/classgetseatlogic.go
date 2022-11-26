package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassgetseatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClassgetseatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassgetseatLogic {
	return &ClassgetseatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClassgetseatLogic) Classgetseat(req *types.ClassSeatreq) (resp *types.ClassSeatres, err error) {
	// todo: add your logic here and delete this line

	return
}
