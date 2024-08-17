package flag

import (
	"reflect"
)

func Scan(target any) {
	sVal, sType := reflect.ValueOf(target), reflect.TypeOf(target)

	if sType.Kind() == reflect.Ptr {
		//用Elem()获得实际的value
		sVal = sVal.Elem()
		sType = sType.Elem()
	}
	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)

		tag := field.Tag
		name := tag.Get("name")
		if name == "" || !sVal.Field(i).CanSet() {
			continue
		}
		v := sVal.FieldByName(field.Name).Addr().UnsafePointer()
		var ve Value
		switch field.Type.Kind() {
		case reflect.String:
			ve = newStringValue("", (*string)(v))
			// todo: 可以解析更多特殊类型
		case reflect.Int:
			ve = newIntValue(0, (*int)(v))
		case reflect.Int8:
			ve = newInt8Value(0, (*int8)(v))
		case reflect.Int16:
			ve = newInt16Value(0, (*int16)(v))
		case reflect.Int32:
			ve = newInt32Value(0, (*int32)(v))
		case reflect.Int64:
			ve = newInt64Value(0, (*int64)(v))
		case reflect.Uint:
			ve = newUintValue(0, (*uint)(v))
		case reflect.Uint8:
			ve = newUint8Value(0, (*uint8)(v))
		case reflect.Uint16:
			ve = newUint16Value(0, (*uint16)(v))
		case reflect.Uint32:
			ve = newUint32Value(0, (*uint32)(v))
		case reflect.Uint64:
			ve = newUint64Value(0, (*uint64)(v))
		case reflect.Float32:
			ve = newFloat32Value(0, (*float32)(v))
		case reflect.Float64:
			ve = newFloat64Value(0, (*float64)(v))
		case reflect.Array:
			ve = newStringArrayValue(nil, (*[]string)(v))
		}

		if value := tag.Get("value"); value != "" {
			_ = ve.Set(value)
		}
		VarP(ve, name, tag.Get("short"), tag.Get("usage"))
	}
}
