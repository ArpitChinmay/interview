package dtomodels

type UpdateCandidate struct {
	ScreeningStatus string `json:"screeningStatus"`
	Date            string `json:"date"`
	InterviewStatus string `json:"interviewStatus"`
	L1Date          string `json:"levelOneDate"`
	L1Panel         string `json:"levelOnePanel"`
	L2Date          string `json:"levelTwoDate"`
	L2Panel         string `json:"levelTwoPanel"`
	DMDate          string `json:"dMDate"`
	DMPanel         string `json:"dMPanel"`
}
