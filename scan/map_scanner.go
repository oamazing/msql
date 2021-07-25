package scan

import "reflect"

type mapScanner struct {
	dest reflect.Value
	dbType string
}

func (ms *mapScanner)Scan(src interface{}) error {
	switch v:=src.(type) {
	case []byte:
		src = string(v)
	}
	ms.dest.SetMapIndex(reflect.ValueOf(ms.dbType),reflect.ValueOf(src))
	return nil
}
