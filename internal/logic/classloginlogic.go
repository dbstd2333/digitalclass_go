package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassloginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClassloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassloginLogic {
	return &ClassloginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClassloginLogic) Classlogin(req *types.Classloginreq) (resp *types.Classloginres, err error) {
	// todo: add your logic here and delete this line

	return
}
