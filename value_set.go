package goutils

import (
	"fmt"
	"strings"
)

func setValInto(o map[string]interface{}, paths []string, val interface{}) {
	switch len(paths) {
	case 0:

	case 1:
		o[paths[0]] = val
	default:
		p, ok := o[paths[0]].(map[string]interface{})
		if !ok {
			p = map[string]interface{}{}
			o[paths[0]] = p
		}
		setValInto(p, paths[1:], val)
	}
}

func SetValInto(o map[string]interface{}, path string, val interface{}) {
	setValInto(o, strings.Split(path, "."), val)
}

func SetValueByExp(o map[string]interface{}, exp string) error {
	kvs := strings.SplitN(exp, "=", 2)
	if len(kvs) != 2 {
		return fmt.Errorf("invalid exp:%s", exp)
	}
	SetValInto(o, kvs[0], kvs[1])
	return nil
}

func parsePaths(path string) []string {
	return strings.Split(path, ".")
}
