package msql

import (
	"context"
	"database/sql"
	"time"
)

type Tx struct {
	tx      *sql.Tx
	TimeOut time.Duration
}

func (tx *Tx) Query(data interface{}, sql string, args ...interface{}) error {
	// TODO 事务查询
	return nil
}
func (tx *Tx) QueryT(ctx context.Context, data interface{}, sql string, args ...interface{}) error {
	// TODO 事务查询
	return nil
}
func (tx *Tx) Exec(sql string, args ...interface{}) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), tx.TimeOut)
	defer cancel()
	return tx.ExecT(ctx, sql, args...)
}
func (tx *Tx) ExecT(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) {
	return tx.tx.ExecContext(ctx, sql, args...)
}
