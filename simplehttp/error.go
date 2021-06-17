package simplehttp

import (
	"fmt"
	"github.com/seeadoog/goutils/excp"
)

var (
	Try               = excp.Try
	TryCatch          = excp.TryCatch
	TryCatchWithStack = excp.TryCatchWithStack
	TryWithStack      = excp.TryWithStack
	TryR              = excp.TryR
)

type ErrorType int

const (
	InvalidBodyType ErrorType = iota
	CreateRequestError
	DoRequestError
	ResponseReadBodyError
	ResponseUnmarshalError
	ResponseWriteError
)

//go:generate stringer -type=ErrorType -output error_string.go
type Error struct {
	Type    ErrorType
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s | %s", e.Type.String(), e.Message)
}

func NewError(t ErrorType, msg string) error {
	return &Error{
		Type:    t,
		Message: msg,
	}
}

func throw(t ErrorType, msg string) {
	excp.Throw(NewError(t, msg))
}
