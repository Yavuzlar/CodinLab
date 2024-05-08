package service_errors

/*
- Hata kodları ve mesajları burada tanımlanır.
- Kodlar 1000'den başlar ve 1000'er artar.
- Örnek Hatalar için predefined.go dosyasına bakınız.
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
