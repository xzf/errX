package main

import (
	"fmt"
	"github.com/xzf/errX"
)

func main() {
	err := demoFunc(0)
	fmt.Println(err.Trace())
	fmt.Println(errX.Error(err).Trace())
}

func demoFunc(i int) *errX.Err {
	lastErr := demoFunc1()
	if i == 0 {
		return lastErr.Warp()
	}
	if i == 1 {
		return lastErr.Warp()
	}
	if i == 2 {
		return lastErr.Warp()
	}
	return lastErr.Warp()
}

func demoFunc1() *errX.Err {
	return errX.New("klutbgss7v")
}
