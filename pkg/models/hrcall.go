package models

type HrCall struct {
	StatusDescr   string `json:"status_descr"`
	EName         string `json:"e_name"`
	ESurname      string `json:"e_surname"`
	Email         string `json:"email"`
	Number        string `json:"number"`
	StatusUpdDate string `json:"status_upd_date"`
	Err           string `json:"error"`
}
