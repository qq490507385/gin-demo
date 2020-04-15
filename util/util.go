package util

import (
	"reflect"
)

// FlectVal struct 值
func FlectVal(s interface{}, key string) reflect.Value {
	strc := reflect.ValueOf(s)
	return strc.FieldByName(key)
}

