// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	kebiaoFieldNames          = builder.RawFieldNames(&Kebiao{})
	kebiaoRows                = strings.Join(kebiaoFieldNames, ",")
	kebiaoRowsExpectAutoSet   = strings.Join(stringx.Remove(kebiaoFieldNames, "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), ",")
	kebiaoRowsWithPlaceHolder = strings.Join(stringx.Remove(kebiaoFieldNames, "`uuid`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), "=?,") + "=?"

	cacheCloudclassKebiaoUuidPrefix = "cache:cloudclass:kebiao:uuid:"
)

type (
	kebiaoModel interface {
		Insert(ctx context.Context, data *Kebiao) (sql.Result, error)
		FindOne(ctx context.Context, uuid int64) (*Kebiao, error)
		Update(ctx context.Context, data *Kebiao) error
		Delete(ctx context.Context, uuid int64) error
	}

	defaultKebiaoModel struct {
		sqlc.CachedConn
		table string
	}

	Kebiao struct {
		Uuid    int64  `db:"uuid"`
		Class   int64  `db:"class"`
		Weekly  string `db:"weekly"`
		Lession int64  `db:"lession"`
		Subject string `db:"subject"`
		Src     string `db:"src"`
	}
)

func newKebiaoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultKebiaoModel {
	return &defaultKebiaoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`kebiao`",
	}
}

func (m *defaultKebiaoModel) Delete(ctx context.Context, uuid int64) error {
	cloudclassKebiaoUuidKey := fmt.Sprintf("%s%v", cacheCloudclassKebiaoUuidPrefix, uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `uuid` = ?", m.table)
		return conn.ExecCtx(ctx, query, uuid)
	}, cloudclassKebiaoUuidKey)
	return err
}

func (m *defaultKebiaoModel) FindOne(ctx context.Context, uuid int64) (*Kebiao, error) {
	cloudclassKebiaoUuidKey := fmt.Sprintf("%s%v", cacheCloudclassKebiaoUuidPrefix, uuid)
	var resp Kebiao
	err := m.QueryRowCtx(ctx, &resp, cloudclassKebiaoUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? limit 1", kebiaoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, uuid)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultKebiaoModel) Insert(ctx context.Context, data *Kebiao) (sql.Result, error) {
	cloudclassKebiaoUuidKey := fmt.Sprintf("%s%v", cacheCloudclassKebiaoUuidPrefix, data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, kebiaoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Class, data.Weekly, data.Lession, data.Subject, data.Src)
	}, cloudclassKebiaoUuidKey)
	return ret, err
}

func (m *defaultKebiaoModel) Update(ctx context.Context, data *Kebiao) error {
	cloudclassKebiaoUuidKey := fmt.Sprintf("%s%v", cacheCloudclassKebiaoUuidPrefix, data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `uuid` = ?", m.table, kebiaoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Class, data.Weekly, data.Lession, data.Subject, data.Src, data.Uuid)
	}, cloudclassKebiaoUuidKey)
	return err
}

func (m *defaultKebiaoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCloudclassKebiaoUuidPrefix, primary)
}

func (m *defaultKebiaoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `uuid` = ? limit 1", kebiaoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultKebiaoModel) tableName() string {
	return m.table
}
