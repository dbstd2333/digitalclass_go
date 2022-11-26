package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudentchangeseatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudentchangeseatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StudentchangeseatLogic {
	return &StudentchangeseatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudentchangeseatLogic) Studentchangeseat(req *types.StudentchangeSeatreq) (resp *types.StudentchangeSeatres, err error) {
	// todo: add your logic here and delete this line

	return
}
