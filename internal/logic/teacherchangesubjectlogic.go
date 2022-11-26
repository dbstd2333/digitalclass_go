package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeacherchangesubjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeacherchangesubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeacherchangesubjectLogic {
	return &TeacherchangesubjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeacherchangesubjectLogic) Teacherchangesubject(req *types.TeacherchangeSubjectreq) (resp *types.TeacherchangeSubjectres, err error) {
	// todo: add your logic here and delete this line

	return
}
