package scan

import (
	"database/sql"
	"errors"
	"reflect"
	"time"
)

func Scan(data interface{}, rows *sql.Rows) error {
	typ := reflect.ValueOf(data)
	if typ.Kind() != reflect.Ptr {
		return errors.New("msql:data must be a point")
	}
	target := typ.Elem()
	switch target.Kind() {
	case reflect.Slice:

	default:
		if rows.Next() {
			return ScanRow(rows, target)
		}
	}
	return nil
}

func ScanRow(rows *sql.Rows, target reflect.Value) error {
	addr := target.Addr().Interface()
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}

	switch addr.(type) {
	case *time.Time:
		return rows.Scan(scannerOf(target, columnTypes[0]))
	}

	// addr := target.Addr().Interface()
	switch target.Kind() {
	case reflect.Struct:
	case reflect.Map:
	default:
		return rows.Scan(scannerOf(target, columnTypes[0]))
	}
	return nil
}
