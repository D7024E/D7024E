package errors

type ValueTimeout struct{}

func (err *ValueTimeout) Error() string {
	return "value has timed out"
}
