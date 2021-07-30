package utils

import (
	"strings"
	"time"
	"unicode"
)

var loc *time.Location

func GetTimeZone() *time.Location {
	return loc
}

func SetTimeZone(name string) error {
	var err error
	loc, err = time.LoadLocation(name)
	if err != nil {
		return err
	}
	return nil
}

// 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	buffer := strings.Builder{}
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteByte('_')
			}
			buffer.WriteRune(r + 32)
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}
