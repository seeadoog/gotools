package goutils

import (
	"fmt"
	"testing"
)

func TestParseRangeIps(t *testing.T) {
	fmt.Println(ParseRangeIps("127.0.0.1:{8000...8005}"))
	fmt.Println(ParseRangeIps("127.0.0.{1...4}:{8000...8005}"))
	fmt.Println(ParseRangeIps("127.0.{0...1}.{1...3}:{8000...8005}"))
}
