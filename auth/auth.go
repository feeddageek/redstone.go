package auth

type Identity struct {
	id    string
	name  string
	perms map[string]bool
}

type Auth interface {
	Authenticate(username string, password string) (Identity, error)
}

