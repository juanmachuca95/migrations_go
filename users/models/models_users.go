package users

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}
