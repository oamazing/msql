package msql

import (
	"log"
	"reflect"
	"strings"

	"github.com/oamazing/msql/utils"
)

func Struct2Fileds(data interface{}, exclude []string) []string {
	typ := reflect.ValueOf(data)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		log.Panic("msql: data is not a struct")
	}
	var fields []string
	structFields(typ, func(name string) {
		for _, v := range exclude {
			if v == name {
				goto walk
			}
		}
		fields = append(fields, utils.Camel2Case(name))
	walk:
	})
	return fields
}

func Struct2FiledCols(data interface{}, exclude []string) string {
	return strings.Join(Struct2Fileds(data, exclude), ",")
}

func structFields(val reflect.Value, fn func(name string)) {
	typ := val.Type()
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.FieldByIndex(field.Index)
		if field.Anonymous {
			fieldTyp := field.Type
			if fieldTyp.Kind() == reflect.Ptr && fieldTyp.Elem().Kind() == reflect.Struct {
				fieldVal = fieldVal.Elem()
			}
			if fieldTyp.Kind() == reflect.Struct {
				structFields(fieldVal, fn)
				continue
			}
		}

		fn(field.Name)
	}
}
