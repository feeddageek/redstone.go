package auth

import (
	"crypto/sha1"
	"encoding/base32"
	"encoding/json"
)
type Json struct {
	j authj
}

type authj struct {
	Salt  string              `json:"Salt"`
	Users map[string]Identity `json:"Users"`
}

func (a *Json) New (data []byte)(err error){
	err = json.Unmarshal(data, &a.j)
	return
}

func (a *Json) Save()(data []byte,err error){
	data,err = json.Marshal(&a.j)
	return
}

func (a Json) Authenticate(username string, password string) (id Identity, err error) {
	id,ok := a.j.Users[username]
	toHash := sha1.Sum([]byte(password + a.j.Salt))
	key := base32.StdEncoding.EncodeToString(toHash[:])
	if ok && a.j.Users[username].Id == key {
		return
	}
	id = Identity{}
	err = AuthError{"Invalid username or password. "+key}
	return
}
