package scan

import (
	"database/sql"
	"errors"
	"reflect"
	"time"

	"github.com/oamazing/msql/utils"
)

func Scan(data interface{}, rows *sql.Rows) error {
	typ := reflect.ValueOf(data)
	if typ.Kind() != reflect.Ptr {
		return errors.New("msql:data must be a point")
	}
	target := typ.Elem()
	switch target.Kind() {
	case reflect.Slice:
		typ := target.Type().Elem()
		for rows.Next() {
			elem := reflect.New(typ).Elem()
			if err := scanRow(rows, elem); err != nil {
				return err
			}
			target.Set(reflect.Append(target, elem))
		}
	default:
		if rows.Next() {
			return scanRow(rows, target)
		}
	}
	return nil
}

func scanRow(rows *sql.Rows, target reflect.Value) error {
	addr := target.Addr().Interface()
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}

	// parse time
	switch addr.(type) {
	case *time.Time:
		return rows.Scan(scannerOf(target, columnTypes[0]))
	}

	if len(columnTypes) == 1 && (target.Kind() == reflect.Struct || target.Kind() == reflect.Map) {
		return rows.Scan(scannerOf(target, columnTypes[0]))
	}

	// addr := target.Addr().Interface()
	switch target.Kind() {
	case reflect.Struct:
		return scan2Struct(rows, target, columnTypes)
	case reflect.Map:
		return scan2Map(rows, target, columnTypes)
	default:
		return rows.Scan(scannerOf(target, columnTypes[0]))
	}
}

func scan2Map(rows *sql.Rows, target reflect.Value, columns []*sql.ColumnType) error {
	if target.IsNil() {
		target.Set(reflect.MakeMap(target.Type()))
	}
	var scanners []interface{}
	for _, column := range columns {
		scanners = append(scanners, &mapScanner{target, column.Name()})
	}
	if err := rows.Scan(scanners...); err != nil {
		return err
	}
	return nil
}

func scan2Struct(rows *sql.Rows, target reflect.Value, columns []*sql.ColumnType) error {
	var scanners []interface{}
	for _, column := range columns {
		field := FieldByName(target, column.Name())
		if !field.IsValid() {
			return errors.New("bsql: no or multiple field '" + column.Name() + "' in struct")
		}
		scanners = append(scanners, scannerOf(field, column))
	}
	if err := rows.Scan(scanners...); err != nil {
		return err
	}
	return nil
}

func FieldByName(v reflect.Value, name string) reflect.Value {
	if f, ok := v.Type().FieldByName(utils.Case2Camel(name)); ok {
		return FieldByIndex(v, f.Index)
	}
	return reflect.Value{}
}

func FieldByIndex(v reflect.Value, index []int) reflect.Value {
	if len(index) == 1 {
		return v.Field(index[0])
	}
	for i, x := range index {
		if i > 0 {
			if v.Kind() == reflect.Ptr && v.Type().Elem().Kind() == reflect.Struct {
				if v.IsNil() {
					v.Set(reflect.New(v.Type().Elem()))
				}
				v = v.Elem()
			}
		}
		v = v.Field(x)
	}
	return v
}
