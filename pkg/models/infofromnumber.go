package models

type InfoFromNumber struct {
	CName    string `json:"c_name"`
	CSurname string `json:"c_surname"`
	Email    string `json:"email"`
	Err      string `json:"error"`
}
