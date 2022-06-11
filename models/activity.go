package models

type Activity struct {
	KeyWord    string `json:"key_word"`
	PagNo      int    `json:"pag_no"`
	ActivityID string `json:"activity_id"`
}

func NewActivity() *Activity {
	return &Activity{}
}
