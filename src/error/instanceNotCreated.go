package err

type InstanceNotCreated struct{}

func (err *InstanceNotCreated) Error() string {
	return "Instance has not been created"
}
