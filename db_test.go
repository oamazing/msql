package msql

import (
	"fmt"
	"log"
	"testing"
)

var db *Db

func init() {
	var err error
	db, err = Open(source)
	if err != nil {
		log.Panic(err)
	}

}

// func TestQuery(t *testing.T) {
// 	var count int64
// 	if err := db.Query(&count, `SELECT count(1) FROM tests`); err != nil {
// 		t.Fatalf("%s", err)
// 	}
// 	fmt.Println(count)

// 	var id int64
// 	if err := db.Query(&id, `SELECT id FROM tests WHERE name = "小王"`); err != nil {
// 		t.Fatalf("%s", err)
// 	}
// 	fmt.Println(id)
// }

func TestQuery(t *testing.T) {
	var disable bool
	if err := db.Query(&disable, `SELECT disable FROM tests where id = 2`); err != nil {
		t.Fatalf("%s", err)
	}
	fmt.Println(disable)
}
