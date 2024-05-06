package hasher_service

import "github.com/matthewhartstonge/argon2"

type HasherService struct {
}

func NewHasherService() *HasherService {
	return &HasherService{}
}

// Hash creates SHA1 hash of given password.
func (h *HasherService) HashPassword(password string) (hashedPassword string, err error) {
	argon := argon2.DefaultConfig()

	hash, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}

	return string(hash), nil

}

func (h *HasherService) CompareHashAndPassword(hashedPassword string, password string) (ok bool, err error) {
	return argon2.VerifyEncoded([]byte(password), []byte(hashedPassword))
}
