package models

type User struct {
	CUUSERREF string `json:"cuuserref"`
	StFlpv    string `json:"st_flpv"`
	CUSUT     string `json:"cusut"`
}

func NewUser() *User {
	return &User{}
}
