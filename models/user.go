package models

type User struct {
	UserID    int    `json:"user_id"`
	CUUSERREF string `json:"cuuserref"`
	StFlpv    string `json:"st_flpv"`
	CUSUT     string `json:"cusut"`
}

func NewUser() *User {
	return &User{}
}
