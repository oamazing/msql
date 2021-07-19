package msql

import "database/sql"

type Tx struct {
	Tx *sql.Tx
}
