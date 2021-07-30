package utils

import (
	"fmt"
	"testing"
)

func ExampleCamel2Case() {
	var names = []string{
		`Abb`,
		`ApllePencel`,
		`CreatedAt`,
	}
	for _, name := range names {
		fmt.Println(Camel2Case(name))
	}
	// output:
	// abb
	// aplle_pencel
	// created_at
}

func ExampleCase2Camel() {
	var names = []string{
		`abb`,
		`aplle_pencel`,
		`created_at`,
	}
	for _, name := range names {
		fmt.Println(Case2Camel(name))
	}
	// output:
	// Abb
	// ApllePencel
	// CreatedAt
}

func Test_GetTimeZone(t *testing.T) {
	if GetTimeZone() == nil {
		t.Failed()
	}
}

func Test_SetTimeZone(t *testing.T) {
	if err := SetTimeZone("Asia/Shanghai"); err != nil {
		t.Failed()
	}
}
