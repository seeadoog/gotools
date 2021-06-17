package simplehttp

import (
	"fmt"
	"testing"
)



func TestRequest_Do(t *testing.T) {
	var res string
	var err error
	Try(func() {
		res = New().GET().Url("http://10.1.87.69:8806/idcs").Do().Text()
	}, &err)
	if err != nil{
		fmt.Println("do request error:",err)
	}else{
		fmt.Println(res)
	}
}

func TestDo2(t *testing.T) {
	var res string
	TryCatchWithStack(func() {
		res = New().GET().Url("http://10.1.87.69:8808/idcs").Do().Text()
	}, func(err error,stack []byte) {
		fmt.Println("do request error:",err,string(stack))
	})

	fmt.Println(res)
}
