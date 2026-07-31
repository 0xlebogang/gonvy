package common

import "golang.org/x/crypto/bcrypt"

type Hasher interface {
	Hash(raw string) (string, error)
	Check(hash, raw string) error
}

type hasher struct{}

func NewHasher() Hasher {
	return &hasher{}
}

func (h *hasher) Hash(raw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (h *hasher) Check(hash, raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
}
