package sqlx

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type (
	Session interface {
		Exec(query string, args ...any) (sql.Result, error)
		ExecCtx(ctx context.Context, query string, args ...any) (sql.Result, error)

		QueryRow(v any, query string, args ...any) error
		QueryRowCtx(ctx context.Context, v any, query string, args ...any) error

		QueryRows(v any, query string, args ...any) error
		QueryRowsCtx(ctx context.Context, v any, query string, args ...any) error
	}

	SqlConn interface {
		Session
		Begin() TxSqlConn
		Transact(fn func(Session) error) error
		TransactCtx(ctx context.Context, fn func(context.Context, Session) error) error
	}
)

type commonSqlConn struct {
	conn *sqlx.DB
}

var _ SqlConn = &commonSqlConn{}

// NewSqlConn returns a SqlConn with given driver name and datasource.
func NewSqlConn(driverName, dataSourceName string) SqlConn {
	conn, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("db connect error,err=%s", err))
	}
	return &commonSqlConn{
		conn: conn,
	}
}

// NewSqlConnFromDB returns a SqlConn with the given sql.DB.
// Use it with caution, it's provided for other ORM to interact with.
func NewSqlConnFromDB(db *sql.DB, driverName string) SqlConn {
	return &commonSqlConn{
		conn: sqlx.NewDb(db, driverName),
	}
}

func (db *commonSqlConn) Exec(query string, args ...any) (sql.Result, error) {
	return db.ExecCtx(context.Background(), query, args...)
}

func (db *commonSqlConn) ExecCtx(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return db.conn.ExecContext(ctx, query, args...)
}

func (db *commonSqlConn) QueryRow(v any, query string, args ...any) error {
	return db.QueryRowCtx(context.Background(), v, query, args...)
}

func (db *commonSqlConn) QueryRowCtx(ctx context.Context, v any, query string, args ...any) error {
	return db.conn.GetContext(ctx, v, query, args...)
}

func (db *commonSqlConn) QueryRows(v any, query string, args ...any) error {
	return db.QueryRowsCtx(context.Background(), v, query, args...)
}

func (db *commonSqlConn) QueryRowsCtx(ctx context.Context, v any, query string, args ...any) error {
	return db.conn.SelectContext(ctx, v, query, args...)
}

func (db *commonSqlConn) Begin() TxSqlConn {
	tx, err := db.conn.Beginx()
	if err != nil {
		panic(fmt.Sprintf("begin transact failed,err=%s", err))
	}
	return &txConn{
		conn: tx,
	}
}

func (db *commonSqlConn) Transact(fn func(Session) error) error {
	f := func(ctx context.Context, sess Session) error {
		return fn(sess)
	}
	return db.TransactCtx(context.Background(), f)
}

func (db *commonSqlConn) TransactCtx(ctx context.Context, fn func(ctx context.Context, sess Session) error) error {

	tx := db.Begin()
	err := fn(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
