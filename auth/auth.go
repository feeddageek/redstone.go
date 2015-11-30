package auth

type Identity struct {
	Id    string          `json:"Id"`
	Name  string          `json:"Name"`
	Perms map[string]bool `json:"Perms"`
}

type AuthError struct {
	err string
}

func (e AuthError) Error() string {
	return e.err
}

type Auth interface {
	Authenticate(username string, password string) (Identity, error)
}
