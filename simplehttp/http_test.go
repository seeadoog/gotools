package simplehttp

import (
	"fmt"
	"testing"
	"unsafe"
)



func TestRequest_Do(t *testing.T) {
	var res string
	var err error
	Try(func() {
		res = New().GET().Urls("http://10.1.87.69:8805/idcs","http://10.1.87.69:8806/idcs").Do().Text()
	}, &err)
	if err != nil{
		switch e := err.(type) {
		case *Error:
			fmt.Println("http error;",e.Type,e.Message)
		}
		fmt.Println("do request error:",err)
	}else{
		fmt.Println(res)
	}
}



func TestT2(t *testing.T) {
	var err error
	var r int
	fmt.Println(uintptr(unsafe.Pointer(&err)))
	fmt.Println(uintptr(unsafe.Pointer(&r)))
	Try2(func() {

	},&err)

}


func Try2(f func(), exception *error) {
	if exception == nil {
		panic("exception is nil")
	}
	defer func() {
		fmt.Println(uintptr(unsafe.Pointer(exception)))
		if err := recover(); err != nil {
			switch e := err.(type) {
			case error:
				*exception = e
			default:
				*exception = NewError(0,fmt.Sprintf("%v",err))
			}
		}
	}()
	f()
}

//

