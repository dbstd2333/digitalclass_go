package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudentupphoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudentupphoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StudentupphoneLogic {
	return &StudentupphoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudentupphoneLogic) Studentupphone(req *types.StudentupPhonereq) (resp *types.StudentupPhoneres, err error) {
	// todo: add your logic here and delete this line

	return
}
