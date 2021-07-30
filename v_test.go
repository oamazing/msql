package msql

import (
	"fmt"
	"time"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Base
}

type Base struct {
	CreatedAt time.Time `json:"created_at"`
}

func Example_Struct2Fileds1() {
	fields := Struct2Fileds(User{}, []string{})
	for _, v := range fields {
		fmt.Println(v)
	}
	// Output:
	// id
	// name
	// created_at
}

func Example_Struct2Fileds2() {
	fields := Struct2Fileds(User{}, []string{`Name`})
	for _, v := range fields {
		fmt.Println(v)
	}
	// Output:
	// id
	// created_at
}

func Example_Struct2FiledCols() {
	fmt.Println(Struct2FiledCols(User{}, []string{}))
	// Output:
	// id,name,created_at
}
