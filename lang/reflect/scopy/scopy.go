package scopy

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func copy(dest *interface{}, src interface{}) {
	srcType := reflect.TypeOf(src)
	destType := reflect.TypeOf(dest)

	srcTypeName := srcType.Name()
	destTypeName := destType.Name()

	srcKind := srcType.Kind()
	destKind := srcType.Kind()

	switch srcKind {
	case reflect.Struct:
		switch destKind {
		case reflect.Struct:
			if srcTypeName == destTypeName {

			} else {

			}
		case reflect.Map:

		case reflect.String:
			dest,_ = json.Marshal(src)
		default:
		}
	case reflect.Map:
		switch destKind {
		case reflect.Map:

		case reflect.Struct:

		case reflect.String:

		default:

		}
	case reflect.Array:
	case reflect.Slice:
	case reflect.Func:
	case reflect.Chan:
	case reflect.String:
		switch destKind {
		case reflect.Int:
			temp := src
			dest = &temp
		case reflect.String:
			dest = strconv.Itoa(src.(int))
		default:

		}
	case reflect.Interface:
	case reflect.Int:
		switch destKind {
		case reflect.Int:
			temp := src
			dest = &temp
		case reflect.String:
			dest = strconv.Itoa(src.(int))
		default:

		}
	case reflect.Int8:
		switch destKind {
		case reflect.Int8:
			temp := src
			dest = &temp
		case reflect.String:
			dest = strconv.Itoa(src.(int))
		default:

		}
	case reflect.Int16:
		switch destKind {
		case reflect.Int16:
			temp := src
			dest = &temp
		case reflect.String:
			dest = strconv.Itoa(src.(int))
		default:

		}
	case reflect.Int32:
		switch destKind {
		case reflect.Int32:
			temp := src
			dest = &temp
		case reflect.String:
			dest = strconv.Itoa(src.(int))
		default:

		}
	case reflect.Int64:
		switch destKind {
		case reflect.Int64:
			temp := src
			dest = &temp
		case reflect.String:
			dest = strconv.Itoa(src.(int))
		default:

		}
	case reflect.Float32:
		switch destKind {
		case reflect.Float32:
			temp := src
			dest = &temp
		case reflect.String:
			dest = fmt.Sprint(src)
		default:

		}
	case reflect.Float64:
		switch destKind {
		case reflect.Int8:
			temp := src
			dest = &temp
		case reflect.String:
			dest = fmt.Sprint(src)
		default:

		}
	case reflect.Bool:
		switch destKind {
		case reflect.Int8:
			temp := src
			dest = &temp
		case reflect.String:
			temp := fmt.Sprint(src)
			dest = &(interface{}(temp))
		default:

		}
	default:

	}
}
