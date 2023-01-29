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

type TechCallService struct{}

func NewTechCallService() *TechCallService {
	return &TechCallService{}
}

func (u *TechCallService) TechCallLogic(jsonInput modules.Request) (code int, any models.TechCall) {
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

	var call models.TechCall
	req := db.QueryRow("SELECT job_name, He.email AS Interviewer, C.email, C.number, Il.interview_date FROM Jobs j INNER JOIN Apply A on j.job_id = A.job_id INNER JOIN Candidate C on A.candidate_id = C.candidate_id INNER JOIN Interview_log Il on C.candidate_id = Il.candidate_id INNER JOIN Hr_employee He on Il.stuff_id = He.stuff_id WHERE C.number = ?", request.Number).Scan(&call.JobName, &call.Interviewer, &call.Email, &call.Number, &call.InterviewDate)
	switch {
	case req == sql.ErrNoRows:
		var response models.TechCall
		response.JobName = ""
		response.Interviewer = ""
		response.Email = ""
		response.Number = ""
		response.InterviewDate = ""
		response.Err = req.Error()
		return 500, response
	case req != nil:
		var response models.TechCall
		response.JobName = ""
		response.Interviewer = ""
		response.Email = ""
		response.Number = ""
		response.InterviewDate = ""
		response.Err = req.Error()
		return 500, response
	}
	//
	var response models.TechCall
	response.JobName = call.JobName
	response.Interviewer = call.Interviewer
	response.Email = call.Email
	response.Number = call.Number
	response.InterviewDate = call.InterviewDate
	response.Err = ""
	return 200, response
}
