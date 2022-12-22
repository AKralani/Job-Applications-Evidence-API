package evidence

type Evidence struct {
	EvidenceID           *int   `json:"evidenceId"`
	Date                 string `json:"date"`
	CompanyName          string `json:"companyName"`
	Link                 string `json:"link"`
	JobDescription       string `json:"jobDescription"`
	Location             string `json:"location"`
	JobType              string `json:"jobType"`
	Field                string `json:"field"`
	InterviewDate        string `json:"interviewDate"`
	InterviewDescription string `json:"interviewDescription"`
	Accepted             bool   `json:"accepted"`
}
