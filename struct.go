package flag

import (
	"fmt"
	"net"
	"reflect"
	"time"
)

// StringP is like String, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Scan(target any) {
	sVal, sType := reflect.ValueOf(target), reflect.TypeOf(target)
	if sType.Kind() != reflect.Ptr {
		panic("target must be pointer")
	}

	// only struct type has tag, can parse
	sVal, sType = sVal.Elem(), sType.Elem()
	if sType.Kind() != reflect.Struct {
		panic("target must be struct")
	}

	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)
		tag := field.Tag

		name, short := tag.Get("name"), tag.Get("short")
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
		case "[]string":
			ve = newStringArrayValue(nil, (*[]string)(v))
		case "int":
			ve = newIntValue(0, (*int)(v))
		case "[]int":
			ve = newIntSliceValue([]int{}, (*[]int)(v))
		case "int8":
			ve = newInt8Value(0, (*int8)(v))
		case "int16":
			ve = newInt16Value(0, (*int16)(v))
		case "int32":
			ve = newInt32Value(0, (*int32)(v))
		case "[]int32":
			ve = newInt32SliceValue([]int32{}, (*[]int32)(v))
		case "int64":
			ve = newInt64Value(0, (*int64)(v))
		case "[]int64":
			ve = newInt64SliceValue([]int64{}, (*[]int64)(v))
		case "uint":
			ve = newUintValue(0, (*uint)(v))
		case "[]uint":
			ve = newUintSliceValue([]uint{}, (*[]uint)(v))
		case "uint8":
			ve = newUint8Value(0, (*uint8)(v))
		case "[]uint8":
			ve = newBytesHexValue([]byte{}, (*[]byte)(v))
		case "uint16":
			ve = newUint16Value(0, (*uint16)(v))
		case "uint32":
			ve = newUint32Value(0, (*uint32)(v))
		case "uint64":
			ve = newUint64Value(0, (*uint64)(v))
		case "float32":
			ve = newFloat32Value(0, (*float32)(v))
		case "[]float32":
			ve = newFloat32SliceValue([]float32{}, (*[]float32)(v))
		case "float64":
			ve = newFloat64Value(0, (*float64)(v))
		case "[]float64":
			ve = newFloat64SliceValue([]float64{}, (*[]float64)(v))
		case "bool":
			ve = newBoolValue(false, (*bool)(v))
			if noOptDefVal == "" {
				noOptDefVal = "true"
			}
		case "[]bool":
			ve = newBoolSliceValue([]bool{}, (*[]bool)(v))
		case "net.IP":
			ve = newIPValue(net.IP{}, (*net.IP)(v))
		case "[]net.IP":
			ve = newIPSliceValue([]net.IP{}, (*[]net.IP)(v))
		case "net.IPMask":
			ve = newIPMaskValue(net.IPMask{}, (*net.IPMask)(v))
		case "count":
			ve = newCountValue(0, (*int)(v))
		case "time.Duration":
			ve = newDurationValue(0, (*time.Duration)(v))
		case "[]time.Duration":
			ve = newDurationSliceValue([]time.Duration{}, (*[]time.Duration)(v))
		default:
			panic(fmt.Sprintf("unsupported parsing type: %v", kind))
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
