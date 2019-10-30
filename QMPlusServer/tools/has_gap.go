package tools

import (
	"fmt"
	"reflect"
)

// HasGap ...
func HasGap(input interface{}) error {
	getType := reflect.TypeOf(input)
	getValue := reflect.ValueOf(input)
	// 获取方法字段
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		switch value.(type) {
		case string:
			if value == "" {
				fmt.Printf("%s为空", field.Name)
				return fmt.Errorf("%s为空", field.Name)
			}
		default:
			if value == nil {
				fmt.Printf("%s为空", field.Name)
				return fmt.Errorf("%s为空", field.Name)
			}
		}
	}
	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	//for i := 0; i < getType.NumMethod(); i++ {
	//	m := getType.Method(i)
	//	fmt.Printf("%s: %v\n", m.Name, m.Type)
	//}
	return nil
}
