package gotest

import "errors"

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return b, errors.New("除数不能为0")
	}
	return a / b, nil
}
