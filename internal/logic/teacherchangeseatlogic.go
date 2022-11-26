package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeacherchangeseatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeacherchangeseatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeacherchangeseatLogic {
	return &TeacherchangeseatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeacherchangeseatLogic) Teacherchangeseat(req *types.TeacherchangeSeatreq) (resp *types.TeacherchangeSeatres, err error) {
	// todo: add your logic here and delete this line

	return
}
