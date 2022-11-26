package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudentloginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudentloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StudentloginLogic {
	return &StudentloginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudentloginLogic) Studentlogin(req *types.StudentLoginreq) (resp *types.StudentLoginres, err error) {
	// todo: add your logic here and delete this line

	return
}
