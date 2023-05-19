package errX

import (
	"bytes"
	"fmt"
	"strconv"
)

type Err struct {
	data []*levelErrInfo
}

const (
	sep = `|`
)

func (e *Err) Error() string {
	errMsgBuff := bytes.NewBuffer(nil)
	for _, item := range e.data {
		if item.errMsg == "" {
			continue
		}
		if errMsgBuff.Len() != 0 {
			errMsgBuff.WriteString(sep)
			errMsgBuff.WriteString(sep)
			errMsgBuff.WriteString(sep)
		}
		errMsgBuff.WriteString(item.errMsg)
	}
	return errMsgBuff.String()
}

func (e *Err) Trace() string {
	traceMsgBuff := bytes.NewBufferString("ErrMsg : ")
	traceMsgBuff.WriteString(e.Error())
	traceMsgBuff.WriteString("\nTraceInfo : \n")
	lenNum := len(strconv.Itoa(len(e.data)))
	fmtStr := "%" + strconv.Itoa(lenNum) + "d "
	for index, item := range e.data {
		traceMsgBuff.WriteString(fmt.Sprintf(fmtStr, index))
		traceMsgBuff.WriteString(sep)
		traceMsgBuff.WriteString(" code line: ")
		traceMsgBuff.WriteString(item.codeLine)
		if index != len(e.data)-1 {
			traceMsgBuff.WriteString("\n")
		}
	}
	return traceMsgBuff.String()
}

func (e *Err) Warp() *Err {
	e.data = append(e.data, newLevelErrInfo("", 2))
	return e
}
