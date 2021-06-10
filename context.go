package goutils

import (
	"context"
	"time"
)

func NewTimeoutContext(timeout time.Duration) context.Context {
	c, _ := context.WithTimeout(context.Background(), timeout)
	return c
}

