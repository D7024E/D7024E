package errors

type ValueNotFound struct{}

func (err *ValueNotFound) Error() string {
	return "Value was not found"
}
