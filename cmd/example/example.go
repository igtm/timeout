package main

import (
	"errors"
	"fmt"
	"github.com/igtm/timeout"
	"time"
)

func doSomethingCurry(str *A) func() error {
	return func() error {
		time.Sleep(2 * time.Second)
		str.X = "asdfa"
		return errors.New("exec error")
	}
}

type A struct {
	X string
}

func main() {
	var a A
	if err := timeout.Timeout(doSomethingCurry(&a), 3*time.Second); err != nil {
		if err == timeout.ErrTimeout {
			fmt.Printf("TIMEOUT: %s\n", err)
		} else {
			fmt.Println(err)
		}
	}
	fmt.Printf("end %+v\n", a)
}
