package errX

import (
	"runtime"
	"strconv"
)

type levelErrInfo struct {
	errMsg   string
	codeLine string
}

func newLevelErrInfo(msg string, level int) *levelErrInfo {
	_, f, l, ok := runtime.Caller(level)
	if ok == false {
		return nil
	}
	return &levelErrInfo{
		codeLine: f + ":" + strconv.Itoa(l),
		errMsg:   msg,
	}
}
