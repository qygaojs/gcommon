package structex

import (
	"fmt"
	"reflect"
)

/**
 * 设置对象值
 */
func SetAttr(obj interface{}, name string, value interface{}) error {
	objValue := reflect.ValueOf(obj)
	valueKind := objValue.Kind()
	structValue := objValue
	if valueKind == reflect.Ptr || valueKind == reflect.Uintptr || valueKind == reflect.UnsafePointer || valueKind == reflect.Interface {
		structValue = objValue.Elem()
	}
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}
	if !structFieldValue.CanSet() {
		return fmt.Errorf("cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	valType := val.Type()
	if structFieldType != valType {
		return fmt.Errorf("Provided value type didnot match obj field type. obj field type:%s value type:%s", structFieldType, valType)
	}

	structFieldValue.Set(val)
	return nil
}

/**
 * 获取对象值
 */
func GetAttr(obj interface{}, name string) interface{} {
	objValue := reflect.ValueOf(obj)
	valueKind := objValue.Kind()
	structValue := objValue
	if valueKind == reflect.Ptr || valueKind == reflect.Uintptr || valueKind == reflect.UnsafePointer || valueKind == reflect.Interface {
		structValue = objValue.Elem()
	}
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		// value不存在
		return nil
	}
	return structFieldValue.Interface()
}

/**
 * 获取对象值
 */
func Keys(obj interface{}) []string {
	objTypes := reflect.TypeOf(obj)
	if objTypes.Kind() == reflect.Ptr {
		objTypes = objTypes.Elem()
	}
	ret := []string{}
	for i := 0; i < objTypes.NumField(); i++ {
		ret = append(ret, objTypes.Field(i).Name)
	}
	return ret
}
