package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudentchangegroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudentchangegroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StudentchangegroupLogic {
	return &StudentchangegroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudentchangegroupLogic) Studentchangegroup(req *types.StudentchangeGroupreq) (resp *types.StudentchangeGroupres, err error) {
	// todo: add your logic here and delete this line

	return
}
