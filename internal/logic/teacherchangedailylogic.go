package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeacherchangedailyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeacherchangedailyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeacherchangedailyLogic {
	return &TeacherchangedailyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeacherchangedailyLogic) Teacherchangedaily(req *types.TeacherchangeDailyreq) (resp *types.TeacherchangeDailyres, err error) {
	// todo: add your logic here and delete this line

	return
}
