package util

import (
	"reflect"
	"unicode"
)

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func MapFirstToLower(data map[string] interface{}) map[string] interface{} {
	var newMap = make(map[string] interface{})
	for k,v := range data {
		newMap[Lcfirst(k)] = v
	}
	return newMap
}

func IsStartUpper(s string) bool {
	return unicode.IsUpper([]rune(s)[0])
}

func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func Typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func MapHasKey(data map[string]interface{}, key string) bool  {
	_, ok := data[key]
	return ok
}

func MapTypeConvert()  {
	
}