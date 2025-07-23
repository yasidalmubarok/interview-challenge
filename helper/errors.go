package helper

type Error interface {
	Error() string
	StatusCode() int
	Message() string
}

type err struct {
	error    string
	message  string
	statusCode int
}

func (e *err) Error() string {
	return e.message
}

func (e *err) StatusCode() int {
	return e.statusCode
}

func (e *err) Message() string {
	return e.message
}

func NewStatusNotFoundError(message string) Error {
	return &err{
		error:     "Not Found",
		message:   message,
		statusCode: 404,
	}
}

func NewStatusBadRequestError(message string) Error {
	return &err{
		error:     "Bad Request",
		message:   message,
		statusCode: 400,
	}
}

func NewStatusInternalServerError(message string) Error {
	return &err{
		error:     "Internal Server Error",
		message:   message,
		statusCode: 500,
	}
}

func NewStatusUnProcessableEntityError(message string) Error {
	return &err{
		error:     "Unprocessable Entity",
		message:   message,
		statusCode: 422,
	}
}