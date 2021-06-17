package sstring

import (
	"fmt"
	"github.com/seeadoog/goutils/excp"
)

var (
	Try               = excp.Try
	TryCatch          = excp.TryCatch
	TryCatchWithStack = excp.TryCatchWithStack
	TryWithStack      = excp.TryWithStack
)


//go:generate stringer -type=ErrorType -output=error_string.go
type ErrorType int

const (
	IntConvertError ErrorType = iota
	FloatConvertError
)

type Error struct {
	Type ErrorType
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s | %s",e.Type.String(),e.Message)
}

func throw(t ErrorType,msg string){
	excp.Throw(&Error{Type: t,Message: msg})
}
