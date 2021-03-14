package main

import (
	"errors"
	"fmt"
)

//error is communicate use same return value, not by raising exception

func testError(param int) (int, error) {
	if param == 42 {
		return -1, errors.New("Can't do with it")
	}
	return param + 3, nil
}

type argError struct {
	param  int
	errMsg string
}

// make own error by create struct implement Error method
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.param, e.errMsg)
}

func testError2(param int) (int, error) {
	if param == 42 {
		return -1, &argError{param: param, errMsg: "Can't work with this value"}
	}
	return param + 3, nil
}

func main() {
	for _, v := range []int{23, 42, 22} {
		if r, e := testError(v); e == nil {
			fmt.Println("Value not error : ", r)
		} else {
			fmt.Println("Error: ", e)
		}
	}

	for _, v := range []int{23, 42, 22} {
		// we can use multiple statement inside if
		if r, e := testError2(v); e == nil {
			fmt.Println("Value not error : ", r)
		} else {
			fmt.Println("Error: ", e)
		}
	}

	// get all var inside error struct use assertion
	_, e := testError2(42)
	// assertion is like data type convertion, but for interface
	ae, isAssertOk := e.(*argError)
	if isAssertOk {
		fmt.Println(ae.param)
		fmt.Println(ae.errMsg)
	}
}
