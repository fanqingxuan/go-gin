package sqlx

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type txConn struct {
	conn *sqlx.Tx
}

type TxSqlConn interface {
	Session
	Commit() error
	Rollback() error
}

var _ TxSqlConn = &txConn{}

func (db *txConn) Exec(query string, args ...any) (sql.Result, error) {
	return db.ExecCtx(context.Background(), query, args...)
}

func (db *txConn) ExecCtx(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return db.conn.ExecContext(ctx, query, args...)
}

func (db *txConn) QueryRow(v any, query string, args ...any) error {
	return db.QueryRowCtx(context.Background(), v, query, args...)
}

func (db *txConn) QueryRowCtx(ctx context.Context, v any, query string, args ...any) error {
	return db.conn.GetContext(ctx, v, query, args...)
}

func (db *txConn) QueryRows(v any, query string, args ...any) error {
	return db.QueryRowsCtx(context.Background(), v, query, args...)
}

func (db *txConn) QueryRowsCtx(ctx context.Context, v any, query string, args ...any) error {
	return db.conn.SelectContext(ctx, v, query, args...)
}

func (db *txConn) Commit() error {
	return db.conn.Commit()
}

func (db *txConn) Rollback() error {
	return db.conn.Rollback()
}
