package goutils

import (
	"fmt"
	"reflect"
	"strconv"
)

func BindFromInterface(v interface{}, m interface{}) error {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Ptr {
		panic("tpye must be ptr not :" + vv.Type().String())
	}
	return bindFromMap(vv, m, "")
}

func bindFromMap(v reflect.Value, m interface{}, path string) (err error) {
	if m == nil {
		return
	}
	switch v.Kind() {
	case reflect.Struct:
		t := v.Type()
		if m == nil {
			return
		}
		mm, ok := m.(map[string]interface{})
		if !ok {
			return fmt.Errorf("bind %s value is not map", path)
		}

		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			ft := t.Field(i)
			if ft.Anonymous {
				err = bindFromMap(fv, m, path)
				if err != nil {
					return err
				}
				continue
			}

			tag := ft.Tag.Get("json")
			if tag == "" {
				tag = ft.Name
			}

			v, ok := mm[tag]
			if ok {
				err = bindFromMap(fv, v, path+"."+tag)
				if err != nil {
					return err
				}
			}

		}
	case reflect.Ptr:
		if v.IsNil() {
			vv := reflect.New(v.Elem().Type())
			if err = bindFromMap(vv.Elem(), m, path); err != nil {
				return
			}
			v.Set(v)
			return
		}
		return bindFromMap(v.Elem(), m, path)
	case reflect.String:
		str, ok := m.(string)
		if !ok {
			return fmt.Errorf("%s is not string", path)
		}
		v.SetString(str)
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int8, reflect.Int16:
		i, ok := m.(float64)
		if !ok {
			return fmt.Errorf("%s is not number", path)
		}
		if i != float64(int(i)) {
			return fmt.Errorf("%s is not int", path)
		}
		v.SetInt(int64(i))
	case reflect.Uint64, reflect.Uint, reflect.Uint32, reflect.Uint8, reflect.Uint16:
		i, ok := m.(float64)
		if !ok {
			return fmt.Errorf("%s is not int", path)
		}
		if i != float64(int(i)) {
			return fmt.Errorf("%s is not int", path)
		}
		v.SetUint(uint64(i))
	case reflect.Bool:
		i, ok := m.(bool)
		if !ok {
			return fmt.Errorf("%s is not bool", path)
		}
		v.SetBool(i)
		return nil
	case reflect.Slice:

		mslice, ok := m.([]interface{})
		if !ok {
			return fmt.Errorf("%s is not slice", path)
		}
		t := v.Type()
		et := t.Elem()

		c := v

		for i := 0; i < len(mslice); i++ {
			var itemValPtr reflect.Value
			if et.Kind() == reflect.Ptr {
				itemValPtr = reflect.New(et.Elem())
			} else if et.Kind() == reflect.Interface {
				c = reflect.Append(c, reflect.ValueOf(mslice[i]))
				continue
			} else {
				itemValPtr = reflect.New(et)
			}
			if err = bindFromMap(itemValPtr, mslice[i], fmt.Sprintf("%s[%d]", path, i)); err != nil {
				return err
			}
			if et.Kind() == reflect.Ptr {
				c = reflect.Append(c, itemValPtr)
			} else {
				c = reflect.Append(c, itemValPtr.Elem())
			}
		}
		v.Set(c)

	case reflect.Map:

		mm, ok := m.(map[string]interface{})
		if !ok {
			return fmt.Errorf("%s is not map", path)
		}
		targetV := reflect.ValueOf(m)
		if targetV.Kind() != reflect.Map {
			return fmt.Errorf("bind %s value is not map", path)
		}

		cv := v
		if cv.IsNil() {
			cv = reflect.MakeMap(v.Type())
		}
		t := v.Type()
		et := t.Elem()

		//it := targetV.MapRange()
		//for it.Next(){
		//	key := it.Key()
		//	val := it.Value()
		//}

		for key, val := range mm {
			var itemValPtr reflect.Value
			if et.Kind() == reflect.Interface {
				cv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
				continue
			}
			if et.Kind() == reflect.Ptr {
				itemValPtr = reflect.New(et.Elem())
			} else {
				itemValPtr = reflect.New(et)
			}
			if err = bindFromMap(itemValPtr, val, fmt.Sprintf("%s.%s", path, key)); err != nil {
				return err
			}
			if et.Kind() == reflect.Ptr {
				cv.SetMapIndex(reflect.ValueOf(key), itemValPtr)
			} else {
				cv.SetMapIndex(reflect.ValueOf(key), itemValPtr.Elem())
			}

		}
		v.Set(cv)
	}
	return nil
}

func I2Int(i interface{}) (res int, err error) {
	switch v := i.(type) {
	case int:
		res = v
	case float64:
		if v != float64(int(v)){
			return 0,fmt.Errorf("%v is not int",v)
		}
		res = int(v)
	case float32:
		if v != float32(int(v)){
			return 0,fmt.Errorf("%v is not int",v)
		}
		res = int(v)
	case int64:
		res = int(v)
	case string:
		res, err = strconv.Atoi(v)
	case bool:
		if v{
			res = 1
		}else{
			res = 0
		}
	case uint64:
		res = int(v)
	case uint:
		res = int(v)
	default:
		return 0,fmt.Errorf("val is not int:%v",i)
	}
	return
}

func I2Bool(i interface{})(res bool,err error){
	switch v := i.(type) {
	case bool:
		res = v
	case string:
		res,err = strconv.ParseBool(v)
	case int:
		res = v > 0
	case float64:
		res = v > 0
	case int64:
		res = v > 0
	case int32:
		res = v > 0
	case float32:
		res = v > 0
		return false,fmt.Errorf("val is not bool:%v",i)
	}
	return
}
