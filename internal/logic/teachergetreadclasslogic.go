package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeachergetreadclassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeachergetreadclassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeachergetreadclassLogic {
	return &TeachergetreadclassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeachergetreadclassLogic) Teachergetreadclass(req *types.TeachergetReadclassreq) (resp *types.TeachergetReadclassres, err error) {
	// todo: add your logic here and delete this line

	return
}
