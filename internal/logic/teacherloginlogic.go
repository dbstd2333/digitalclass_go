package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeacherloginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeacherloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeacherloginLogic {
	return &TeacherloginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeacherloginLogic) Teacherlogin(req *types.TeacherLoginreq) (resp *types.TeacherLoginres, err error) {
	// todo: add your logic here and delete this line

	return
}
