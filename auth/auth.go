package auth

type Identity struct {
	Id    string
	Name  string
	Perms []string
}

type Auth interface {
	Authenticate(username string, password string) (Identity, error)
}
