package model

type Patient struct {
	StudentID  int    `json:"StudentID"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	RoomNumber int    `json:"roomNumber"`
	course  string `json:"course"`
}