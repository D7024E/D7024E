package errors

type ValueAlreadyExist struct{}

func (err *ValueAlreadyExist) Error() string {
	return "Value already exist"
}
