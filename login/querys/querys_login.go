package login

var GetUser = func() string {
	return "SELECT username, password FROM users WHERE username=?"
}
