package errors

type ValueNotFound struct{}

func (err *ValueNotFound) Error() string {
	return "value was not found"
}
