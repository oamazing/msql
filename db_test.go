package msql

import (
	"fmt"
	"log"
	"testing"
	"time"
)

var db *Db

func init() {
	var err error
	db, err = Open(source)
	if err != nil {
		log.Panic(err)
	}
	db.SetTimeOut(time.Minute)
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

func Test_Query(t *testing.T) {
	// t.Run(`test bool`, func(t *testing.T) {
	// 	var disable bool
	// 	if err := db.Query(&disable, `SELECT disable FROM tests where id = 2`); err != nil {
	// 		t.Fatalf("%s", err)
	// 	}
	// 	fmt.Println(disable)
	// })
	// t.Run(`test int 1`, func(t *testing.T) {
	// 	var count int64
	// 	if err := db.Query(&count, `SELECT count(1) FROM tests`); err != nil {
	// 		t.Fatalf("%s", err)
	// 	}
	// 	fmt.Println(count)
	// })
	// t.Run(`test int 2`, func(t *testing.T) {
	// 	var id int64
	// 	if err := db.Query(&id, `SELECT id FROM tests WHERE name = "小王"`); err != nil {
	// 		t.Fatalf("%s", err)
	// 	}
	// 	fmt.Println(id)
	// })
	t.Run(`test datatime`, func(t *testing.T) {
		var createdAt time.Time
		if err := db.Query(&createdAt, `SELECT created_at FROM tests WHERE id = 2`); err != nil {
			t.Fatalf("%s", err)
		}
		fmt.Println(createdAt)
	})
}
