package service

import (
	"bcc-hackathon-go/pkg/models"
	"bcc-hackathon-go/pkg/modules"
	"bytes"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type HrCallService struct{}

func NewHrCallService() *HrCallService {
	return &HrCallService{}
}

func (u *HrCallService) HrCallLogic(jsonInput modules.Request) (code int, any models.HrCall) {
	requestBodyBytes := new(bytes.Buffer)
	json.NewEncoder(requestBodyBytes).Encode(jsonInput)

	var request modules.Request
	json.Unmarshal(requestBodyBytes.Bytes(), &request)

	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/hackathon")
	if err != nil {
		log.Print(err.Error())
	}

	defer db.Close()

	//var points models.Points
	//

	var call models.HrCall
	req := db.QueryRow("SELECT s.status_descr, h.e_name, h.e_surname, h.email, h.number, a.status_upd_date FROM Status s INNER JOIN Apply_status a USING(status_id) INNER JOIN Apply ap USING(apply_id) INNER JOIN Hr_employee h USING(stuff_id) INNER JOIN Candidate c USING(candidate_id) WHERE c.number = ? AND a.status_upd_date = (SELECT MAX(a.status_upd_date) FROM Apply_status a)", request.Number).Scan(&call.StatusDescr, &call.EName, &call.ESurname, &call.Email, &call.Number, &call.StatusUpdDate)
	switch {
	case req == sql.ErrNoRows:
		var response models.HrCall
		response.StatusDescr = ""
		response.EName = ""
		response.ESurname = ""
		response.Email = ""
		response.Number = ""
		response.StatusUpdDate = ""
		response.Err = req.Error()
		return 500, response
	case req != nil:
		var response models.HrCall
		response.StatusDescr = ""
		response.EName = ""
		response.ESurname = ""
		response.Email = ""
		response.Number = ""
		response.StatusUpdDate = ""
		return 500, response
	}
	//
	var response models.HrCall
	response.StatusDescr = call.StatusDescr
	response.EName = call.EName
	response.ESurname = call.ESurname
	response.Email = call.Email
	response.Number = call.Number
	response.StatusUpdDate = call.StatusUpdDate
	response.Err = ""
	return 200, response
}
