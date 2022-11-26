package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeacherchangegroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeacherchangegroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeacherchangegroupLogic {
	return &TeacherchangegroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeacherchangegroupLogic) Teacherchangegroup(req *types.TeacherchangeGroupreq) (resp *types.TeacherchangeGroupres, err error) {
	// todo: add your logic here and delete this line

	return
}
