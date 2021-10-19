package login

var GetUser = func() string {
	return "SELECT * FROM users WHERE username=?"
}
