package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassgetmsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClassgetmsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassgetmsgLogic {
	return &ClassgetmsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClassgetmsgLogic) Classgetmsg(req *types.ClassMsgreq) (resp *types.ClassMsgres, err error) {
	// todo: add your logic here and delete this line

	return
}
