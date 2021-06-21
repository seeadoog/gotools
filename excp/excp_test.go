package excp

import (
	"errors"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTry(t *testing.T) {
	var err error
	Try(func() {
		Throw(NewError("error occur"))
	}, &err)

	if err != nil {
		fmt.Println("err=>", err)
		switch e := err.(type) {
		case *DefaultError:
			fmt.Println("e is default error:", e)
		case *Error:
			fmt.Println("is error:", e)
		}
	} else {
		fmt.Println("success")
	}

}

type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func NewError(msg string) *Error {
	return &Error{Message: msg}
}

func TestCatch(t *testing.T) {
	TryCatch(func() {
		Throw(NewError("throw error"))
	}, func(err error) {
		switch e := err.(type) {
		case *Error:
			fmt.Println(e.Message)
		}
	})
}

func BenchmarkTryCatch(b *testing.B) {
	err := NewError("throw error")
	for i := 0; i < b.N; i++ {
		TryCatch(func() {
			Throw(err)
		},
			func(err error) {

			})
	}
}

func TestTryCatchWithStack(t *testing.T) {

	TryCatchWithStack(func() {
		testfunc()
	}, func(err error, stack []byte) {
		fmt.Println(string(stack))
	})
}

func testfunc() {
	panic("gg")
}

func TestTr(t *testing.T) {
	err := TryR(func() {
		Throw(errors.New("hello "))
	})
	if err != nil {
		fmt.Println("err is: ", err)
	}

	TryCatch(func() {

	},
		func(err error) {

		})
}

func TestStack(t *testing.T) {
	err, stack := TryRWithStack(func() {
		testfunc()
	})
	if err != nil {
		fmt.Println(err, string(stack))
	}

}

func TestG(t *testing.T) {
	for i := 0; i < 1e6; i++ {
		go func() {
			for{
				time.Sleep(100*time.Minute)
			}
		}()
	}
	select {

	}
}

func TestName(t *testing.T) {

	a := testdf()
	fmt.Println(a)
}

func testdf()(res int){
	defer func() {
		res = 5
	}()
	return 4
}

func TestL(t *testing.T) {
	TryCatchWithStack(ttt, func(err error,stack []byte) {
		log.Println("error is :",err,"stack:",string(stack))
	})

	TryCatch(ttt, func(err error,) {
		log.Println("error is :",err,"stack:")
	})
}


func ttt(){
	_,err := fmt.Println()
	Throw(fmt.Errorf("%w",err))
}
