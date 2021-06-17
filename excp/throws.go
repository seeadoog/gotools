package excp

import "os"

func Open(file string) *os.File {
	f, err := os.Open(file)
	Throw(err)
	return f
}
