package serviceErrors

type ServiceError struct {
	Message string
}

func NewError(message string) *ServiceError {
	return &ServiceError{
		Message: message,
	}
}

func (err *ServiceError) Error() string {
	return err.Message
}
