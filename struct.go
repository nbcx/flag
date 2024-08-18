package flag

import (
	"reflect"
)

// StringP is like String, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Scan(target any) {
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
		short := tag.Get("short")
		if (short == "" && name == "") || !sVal.Field(i).CanSet() {
			continue
		}
		v := sVal.FieldByName(field.Name).Addr().UnsafePointer()
		var ve Value

		kind := tag.Get("type")
		if kind == "" {
			kind = field.Type.String()
		}
		noOptDefVal := tag.Get("def")
		switch kind {
		case "string":
			ve = newStringValue("", (*string)(v))
		case "int":
			ve = newIntValue(0, (*int)(v))
		case "int8":
			ve = newInt8Value(0, (*int8)(v))
		case "int16":
			ve = newInt16Value(0, (*int16)(v))
		case "int32":
			ve = newInt32Value(0, (*int32)(v))
		case "int64":
			ve = newInt64Value(0, (*int64)(v))
		case "uint":
			ve = newUintValue(0, (*uint)(v))
		case "uint8":
			ve = newUint8Value(0, (*uint8)(v))
		case "uint16":
			ve = newUint16Value(0, (*uint16)(v))
		case "uint32":
			ve = newUint32Value(0, (*uint32)(v))
		case "uint64":
			ve = newUint64Value(0, (*uint64)(v))
		case "float32":
			ve = newFloat32Value(0, (*float32)(v))
		case "float64":
			ve = newFloat64Value(0, (*float64)(v))
		case "bool":
			ve = newBoolValue(false, (*bool)(v))
			if noOptDefVal == "" {
				noOptDefVal = "true"
			}
		case "[]string":
			ve = newStringArrayValue(nil, (*[]string)(v))
		case "[]slice":
			ve = newStringArrayValue(nil, (*[]string)(v))
		}
		if value := tag.Get("value"); value != "" {
			_ = ve.Set(value)
		}
		pf := f.VarPF(ve, name, short, tag.Get("usage"))
		if noOptDefVal != "" {
			pf.NoOptDefVal = noOptDefVal
		}
	}
}

func Scan(target any) {
	CommandLine.Scan(target)
}
