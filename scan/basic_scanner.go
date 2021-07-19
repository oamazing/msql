package scan

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

const (
	minInt32  = -1 << 31
	maxInt32  = 1<<31 - 1
	maxUint32 = 1<<32 - 1

	minInt16  = -1 << 15
	maxInt16  = 1<<15 - 1
	maxUint16 = 1<<16 - 1

	minInt8  = -1 << 7
	maxInt8  = 1<<7 - 1
	maxUint8 = 1<<8 - 1
)

type basicScanner struct {
	dest   reflect.Value
	dbType string
}

func (bs *basicScanner) Scan(srcIfc interface{}) error {
	if srcIfc == nil {
		bs.dest.Set(reflect.Zero(bs.dest.Type()))
		return nil
	}
	src, ok := srcIfc.([]byte)
	if !ok {
		return errors.New("msql: can not assert []byte")
	}
	switch bs.dbType {
	case `BIGINT`, `TINYINT`, `INTEGER`, `SMALLINT`:
		return scanInteger(bs.dest, src)
	case `VARCHAR`, `LONGTEXT`:
		return scanString(bs.dest, src)
	default:
		return fmt.Errorf("msql: unsupport type %s", bs.dbType)
	}
}

func scanInteger(dest reflect.Value, src []byte) error {

	switch dest.Kind() {
	case reflect.Int64, reflect.Int, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Bool:
		i, err := strconv.ParseInt(string(src), 10, 64)
		if err != nil {
			return errors.New("msql: parse bigint error")
		}
		return scanInt(dest, i)
	case reflect.Uint64, reflect.Uint, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		i, err := strconv.ParseUint(string(src), 10, 64)
		if err != nil {
			return errors.New("msql: parse bigint error")
		}
		return scanUInt(dest, i)
	default:
		return errorCannotAssign(src, dest)
	}
}

func scanUInt(dest reflect.Value, i uint64) error {
	switch dest.Kind() {
	case reflect.Int64, reflect.Int:
		dest.SetUint(i)
	case reflect.Uint32:
		if i > maxUint32 {
			return errorValueOutOfRange(i, dest)
		}
		dest.SetUint(i)
	case reflect.Uint16:
		if i > maxUint16 {
			return errorValueOutOfRange(i, dest)
		}
		dest.SetUint(i)
	case reflect.Uint8:
		if i > maxUint8 {
			return errorValueOutOfRange(i, dest)
		}
		dest.SetUint(i)
	default:
		errorCannotAssign(i, dest)
	}
	return nil
}

func scanInt(dest reflect.Value, i int64) error {
	switch dest.Kind() {
	case reflect.Int64, reflect.Int:
		dest.SetInt(i)
	case reflect.Int32:
		if i < minInt32 || i > maxInt32 {
			return errorValueOutOfRange(i, dest)
		}
		dest.SetInt(i)
	case reflect.Int16:
		if i < minInt16 || i > maxInt16 {
			return errorValueOutOfRange(i, dest)
		}
		dest.SetInt(i)
	case reflect.Int8:
		if i < minInt8 || i > maxInt8 {
			return errorValueOutOfRange(i, dest)
		}
		dest.SetInt(i)
	case reflect.Bool:
		if i > 1 || i < 0 {
			return errorValueOutOfRange(i, dest)
		}
		dest.SetBool(i == 1)
	default:
		errorCannotAssign(i, dest)
	}
	return nil
}

func scanString(dest reflect.Value, src []byte) error {
	dest.SetString(string(src))
	return nil
}

func errorValueOutOfRange(src interface{}, dest reflect.Value) error {
	return fmt.Errorf("msql: cannot assign %T(%v) to %v: value out of range", src, src, dest.Type())
}

func errorCannotAssign(src interface{}, dest reflect.Value) error {
	return fmt.Errorf("msql: cannot assign %T(%v) to %v", src, src, dest.Type())
}
