package scan

import (
	"database/sql"
	"reflect"
)

func scannerOf(dest reflect.Value, column *sql.ColumnType) sql.Scanner {
	return &basicScanner{dest: dest, dbType: column.DatabaseTypeName()}
}
