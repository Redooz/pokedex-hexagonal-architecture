package error

type NoDataFound struct {
	Message string
}

func (e *NoDataFound) Error() string {
	return e.Message
}
