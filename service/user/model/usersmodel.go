package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheGinUsersIdPrefix = "cache:gin:users:id:"
)

type (
	UsersModel interface {
		Insert(data *Users) (sql.Result, error)
		FindOne(id int64) (*Users, error)
		Update(data *Users) error
		Delete(id int64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id        int64         `db:"id"`
		CreatedAt int64         `db:"created_at"`
		UpdatedAt int64         `db:"updated_at"`
		Name      string        `db:"name"` // 姓名
		Password  string        `db:"password"`
		Age       int64         `db:"age"`
		DeletedAt sql.NullInt64 `db:"deleted_at"`
	}
)

func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf) UsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Insert(data *Users) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.CreatedAt, data.UpdatedAt, data.Name, data.Password, data.Age, data.DeletedAt)

	return ret, err
}

func (m *defaultUsersModel) FindOne(id int64) (*Users, error) {
	ginUsersIdKey := fmt.Sprintf("%s%v", cacheGinUsersIdPrefix, id)
	var resp Users
	err := m.QueryRow(&resp, ginUsersIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
		return conn.QueryRow(v, query, id)
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

func (m *defaultUsersModel) Update(data *Users) error {
	ginUsersIdKey := fmt.Sprintf("%s%v", cacheGinUsersIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.Exec(query, data.CreatedAt, data.UpdatedAt, data.Name, data.Password, data.Age, data.DeletedAt, data.Id)
	}, ginUsersIdKey)
	return err
}

func (m *defaultUsersModel) Delete(id int64) error {

	ginUsersIdKey := fmt.Sprintf("%s%v", cacheGinUsersIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, ginUsersIdKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGinUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	return conn.QueryRow(v, query, primary)
}
