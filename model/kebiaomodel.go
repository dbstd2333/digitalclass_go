package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ KebiaoModel = (*customKebiaoModel)(nil)

type (
	// KebiaoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customKebiaoModel.
	KebiaoModel interface {
		kebiaoModel
	}

	customKebiaoModel struct {
		*defaultKebiaoModel
	}
)

// NewKebiaoModel returns a model for the database table.
func NewKebiaoModel(conn sqlx.SqlConn, c cache.CacheConf) KebiaoModel {
	return &customKebiaoModel{
		defaultKebiaoModel: newKebiaoModel(conn, c),
	}
}
