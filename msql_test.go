package msql

import "testing"

var source = "root:sqdas1994210@tcp(cdb-fuxb3dbb.gz.tencentcdb.com:10174)/msql_test?charset=utf8"

func TestOpen(t *testing.T) {
	_, err := Open(source)
	if err != nil {
		t.Fail()
	}
}
