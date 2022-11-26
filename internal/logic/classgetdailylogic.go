package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassgetdailyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClassgetdailyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassgetdailyLogic {
	return &ClassgetdailyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClassgetdailyLogic) Classgetdaily(req *types.ClassDailyreq) (resp *types.ClassDailyres, err error) {
	// todo: add your logic here and delete this line

	return
}
