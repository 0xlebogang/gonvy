package password

import "golang.org/x/crypto/bcrypt"

type Password interface {
	Hash(raw string) (string, error)
	Check(hash, raw string) error
}

type password struct{}

func New() Password {
	return &password{}
}

func (p *password) Hash(raw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (p *password) Check(hash, raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
}
