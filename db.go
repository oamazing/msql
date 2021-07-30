package msql

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/oamazing/msql/scan"
)

type Db struct {
	db      *sql.DB
	TimeOut time.Duration
}

// 查询
func (db *Db) Query(data interface{}, sql string, args ...interface{}) error {
	return db.QueryT(db.TimeOut, data, sql, args...)
}
func (db *Db) QueryT(d time.Duration, data interface{}, sql string, args ...interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()
	return db.query(ctx, data, sql, args...)
}
func (db *Db) Exec(sql string, args ...interface{}) (sql.Result, error) {
	return db.ExecT(db.TimeOut, sql, args...)
}
func (db *Db) ExecT(t time.Duration, sql string, args ...interface{}) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	return db.db.ExecContext(ctx, sql, args...)
}

func (db *Db) query(ctx context.Context, data interface{}, sql string, args ...interface{}) error {
	rows, err := db.db.QueryContext(ctx, sql, args...)
	if debug {
		debugSql(sql)
	}
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return err
	}
	return scan.Scan(data, rows)
}

func (db *Db) SetMaxIdleConns(n int) {
	db.db.SetMaxIdleConns(n)
}

func (db *Db) SetMaxOpenConns(n int) {
	db.db.SetMaxOpenConns(n)
}

func (db *Db) SetConnMaxIdleTime(t time.Duration) {
	db.db.SetConnMaxIdleTime(t)
}

func (db *Db) SetConnMaxLifetime(t time.Duration) {
	db.db.SetConnMaxLifetime(t)
}

var debug = os.Getenv(`DebugSql`) != ``

func debugSql(sql string) {
	fmt.Printf("\x1b[%dm%s\x1b[0m\n", 32, sql)
}

func (db *Db) RunInTransaction(fn func(*Tx) error) error {
	otx, err := db.db.Begin()
	if err != nil {
		return err
	}
	tx := &Tx{tx: otx, TimeOut: db.TimeOut}
	if err = fn(tx); err != nil {
		return tx.tx.Rollback()
	}
	return tx.tx.Commit()
}
