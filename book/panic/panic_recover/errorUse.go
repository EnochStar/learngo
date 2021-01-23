package panic_recover

import (
	"errors"
	"fmt"
)

var ErrDivByZero = errors.New("division by zero")

func Div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}
func ErrorCatch() {
	defer func() {
		fmt.Println(recover())
	}()
	switch z, err := Div(10, 0); err {
	case nil:
		println(z)
	case ErrDivByZero:
		panic(err)
	}
}
