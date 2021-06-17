package excp

import "os"

func Open(file string) *os.File {
	f, err := os.Open(file)
	if err != nil {
		Throw(err)
	}
	return f
}

