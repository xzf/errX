package errX

func New(errMsg string) *Err {
	return &Err{
		data: []*levelErrInfo{
			newLevelErrInfo(errMsg, 2),
		},
	}
}

func Error(err error) *Err {
	thisErr, ok := err.(*Err)
	if ok {
		thisErr.data = append(thisErr.data, newLevelErrInfo("", 2))
		return thisErr
	}
	return New(err.Error())
}
