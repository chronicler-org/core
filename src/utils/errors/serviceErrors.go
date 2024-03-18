package servoceErrors

type ServiceError struct {
	Status  int
	Message string
}

func NewError(status int, message string) *ServiceError {
	return &ServiceError{
		Status:  status,
		Message: message,
	}
}

func (err *ServiceError) Error() string {
	return err.Message
}

const (
	BadRequestError     = 400
	NotFoundError       = 404
	InternalServerError = 500
)
