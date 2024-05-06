package service_errors

import "strings"

/*
- Hata kodları ve mesajları burada tanımlanır.
- Kodlar 1000'den başlar ve 1000'er artar.
- Örnek Hatalar için predefined.go dosyasına bakınız.
*/
type ServiceError struct {
	Code    int
	Key     string
	Message string
	err     error
}

func (e *ServiceError) Error() string { // Error() fonksiyonu error interface'ini implemente eder. Bu sayede error olarak döndürülebilir.
	if e.err != nil { // Eğer err varsa err'i döndürür.
		return e.err.Error()
	}
	if e.Key != "" { // Key varsa key'i döndürür.
		return e.Key
	}
	// Mesajı küçük harfe çevirir.
	return strings.ToLower(e.Message)
}
func NewServiceErrorWithMessage(code int, message string) error {
	return &ServiceError{
		Code:    code,
		Message: message,
	}
}
func NewServiceErrorWithKey(code int, key string) error {
	return &ServiceError{
		Code: code,
		Key:  key,
	}
}
func NewServiceErrorWithMessageAndError(code int, message string, err error) error {
	return &ServiceError{
		Code:    code,
		Message: message,
		err:     err,
	}
}
