package main

import (
	"custom/myerror"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		// Menggunakan custom interface MyError
		return 0, myerror.NewErrorInvalidArgument("division by zero")
	}
	return a / b, nil
}

func main() {
	res, err := divide(1, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}
