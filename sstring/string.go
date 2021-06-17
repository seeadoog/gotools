package sstring

import (
	"bytes"
	"strconv"
	"unsafe"
)

type String []byte

func Of(s string) String {
	p := (*[2]uintptr)(unsafe.Pointer(&s))
	t := [3]uintptr{p[0], p[1], p[1]}
	return *(*String)(unsafe.Pointer(&t))
}

func (s String) Len() int {
	return len(s)
}

func (s String) Cap() int {
	return cap(s)
}

func (s String) Slice(i, j int) String {
	return s[i:j]
}

func (s String) Equal(o String) bool {
	return bytes.Equal(s, o)
}

func (s String) HasPrefix(o String) bool {
	return bytes.HasPrefix(s, o)
}

func (s String) HasSuffix(o String) bool {
	return bytes.HasSuffix(s, o)
}

func (s String) Contains(o String) bool {
	return bytes.Contains(s, o)
}

func (s String) String() string {
	return *(*string)(unsafe.Pointer(&s))
}

func (s String) Int() (int, error) {
	return strconv.Atoi(s.String())
}

func (s String) MustInt() int {
	v, err := s.Int()
	if err != nil {
		throw(IntConvertError, err.Error())
	}
	return v
}

func (s String) Float64() (float64, error) {
	return strconv.ParseFloat(s.String(), 64)
}

func (s String) MustFloat64() float64 {
	v, err := s.Float64()
	if err != nil {
		throw(FloatConvertError, err.Error())
	}
	return v
}

func (s *String) Set(o String) {
	*s = append((*s)[:0], o...)
}

func (s String) Copy() String {
	t := make(String, s.Len())
	copy(t, s)
	return t
}

func (s String) Index(sep String) int {
	return bytes.Index(s, sep)
}
