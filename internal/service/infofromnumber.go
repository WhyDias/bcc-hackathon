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

type InfoFromNumberService struct{}

func NewInfoFromNumberService() *InfoFromNumberService {
	return &InfoFromNumberService{}
}

func (u *InfoFromNumberService) InfoFromNumberLogic(jsonInput modules.Request) (code int, any models.InfoFromNumber) {
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

	var info models.InfoFromNumber
	req := db.QueryRow("SELECT c_name, c_surname, email FROM Candidate WHERE number =  ?", request.Number).Scan(&info.CName, &info.CSurname, &info.Email)
	switch {
	case req == sql.ErrNoRows:
		var response models.InfoFromNumber
		response.CName = ""
		response.CSurname = ""
		response.Email = ""
		response.Err = req.Error()
		return 500, response
	case req != nil:
		var response models.InfoFromNumber
		response.CName = ""
		response.CSurname = ""
		response.Email = ""
		response.Err = req.Error()
		return 500, response
	}
	//
	var response models.InfoFromNumber
	response.CName = info.CName
	response.CSurname = info.CSurname
	response.Email = info.Email
	response.Err = ""
	return 200, response
}
