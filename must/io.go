package must

import (
	"github.com/seeadoog/goutils/excp"
	"os"
	"strconv"
)

var throw = excp.Throw

//@throw error
func Open(f string) *os.File {
	file, err := os.Open(f)
	throw(err)
	return file
}

//@throw error
func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	throw(err)
	return v
}

// func a(s string)(i int)
/*
	var err error
	try(func(){
		a("ss")
	},&err)
	switch e := err.(type){
		case nil:
	}
 */
