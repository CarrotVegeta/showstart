package models

type Activity struct {
	ActivityID string `json:"activity_id"`
}

func NewActivity() *Activity {
	return &Activity{}
}
