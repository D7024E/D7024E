package err

type ObjectNotFound struct{}

func (err *ObjectNotFound) Error() string {
	return "Object was not found"
}
