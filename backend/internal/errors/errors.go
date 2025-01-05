package service_errors

/*
- Error codes and messages are specified here.
- Error codes start from 1000 and increase by 1000.
- For example errors, you can refer to the predefined.go file.
*/
type ServiceError struct {
	Code    int
	Message string
	err     error
}

func (e *ServiceError) Error() string { // Error() fonksiyonu error interface'ini implemente eder. Bu sayede error olarak döndürülebilir.
	if e.err != nil { // Eğer err varsa err'i döndürür.
		return e.err.Error()
	}
	return ""
}
func NewServiceErrorWithMessage(code int, message string) error {
	return &ServiceError{
		Code:    code,
		Message: message,
	}
}

func NewServiceErrorWithMessageAndError(code int, message string, err error) error {
	return &ServiceError{
		Code:    code,
		Message: message,
		err:     err,
	}
}
