package msql

import (
	"database/sql"
	"time"

	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/oamazing/msql/utils"
)

func init() {
	// 初始化时区
	if err := utils.SetTimeZone(`Asia/Shanghai`); err != nil {
		panic(err)
	}
}

const driver = "mysql"

type DbOrTx interface {
	Query(data interface{}, sql string, args ...interface{})
	QueryT(ctx context.Context, data interface{}, sql string, args ...interface{})
	Exec(sql string, args ...interface{})
	ExecT(ctx context.Context, sql string, args ...interface{})
}

func Open(dataSourceName string) (*Db, error) {
	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &Db{db: db, TimeOut: time.Duration(time.Second)}, nil
}

func (db *Db) SetTimeOut(d time.Duration) {
	db.TimeOut = d
}
