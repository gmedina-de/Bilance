package authenticator

type Authenticator interface {
	Authenticate(username, password string) bool
}
