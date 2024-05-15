package login

type user struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	EmailId  string `json:"emailid`
	Password string `json:"password"`
}
