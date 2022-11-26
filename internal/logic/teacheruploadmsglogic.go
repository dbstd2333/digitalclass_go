package logic

import (
	"context"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeacheruploadmsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeacheruploadmsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeacheruploadmsgLogic {
	return &TeacheruploadmsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeacheruploadmsgLogic) Teacheruploadmsg(req *types.TeacherupMsgreq) (resp *types.TeacherupMsgres, err error) {
	// todo: add your logic here and delete this line

	return
}
