package models

type TechCall struct {
	JobName       string `json:"job_name"`
	Interviewer   string `json:"Interviewer"`
	Email         string `json:"email"`
	Number        string `json:"number"`
	InterviewDate string `json:"interview_date"`
	Err           string `json:"error"`
}
