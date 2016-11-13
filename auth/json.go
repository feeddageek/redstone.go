package auth

import (
	"crypto/sha1"
	"encoding/base32"
	"errors"
)

type Authj struct {
	Salt  string
	Users map[string]Identity
}

func (a Authj) Authenticate(username string, password string) (id Identity, err error) {
	id, ok := a.Users[username]
	toHash := sha1.Sum([]byte(password + a.Salt))
	key := base32.StdEncoding.EncodeToString(toHash[:])
	if ok && a.Users[username].Id == key {
		return
	}
	id = Identity{}
	err = errors.New("Invalid username or password. " + key)
	return
}
