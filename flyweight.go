package flyweight

import (
	"reflect"
)

func Reset(v interface{}) {
	r := reflect.ValueOf(v)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		ty := field.Type()

		switch field.Kind() {
		case reflect.Struct:
			tryCallReset(field, ty)
		case reflect.Slice:
			for i := 0; i < field.Len(); i++ {
				element := field.Index(i)
				elementType := element.Type()
				tryCallReset(element, elementType)
			}
			field.SetLen(0)
		case reflect.Ptr:
			if !field.IsNil() {
				callReset(field)
			}
		default:
			field.Set(reflect.Zero(ty))
		}
	}
}

func tryCallReset(field reflect.Value, ty reflect.Type) {
	if called := callReset(field); called {
		return
	}

	pt := ptr(field)
	if called := callReset(pt); called {
		field.Set(pt.Elem())
		return
	}

	field.Set(reflect.Zero(ty))
}

func callReset(field reflect.Value) bool {
	reset := field.MethodByName("Reset")
	if reset.IsValid() {
		reset.Call(nil)
		return true
	}
	return false
}

func ptr(v reflect.Value) reflect.Value {
	pt := reflect.PtrTo(v.Type())
	pv := reflect.New(pt.Elem())
	pv.Elem().Set(v)
	return pv
}
