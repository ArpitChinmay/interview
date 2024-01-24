package models

type Interview struct {
	InterviewStatusId string
	CandidateId       string
	InterviewStatus   string
	L1ScheduledDate   string
	L1Panel           string
	L2ScheduledDate   string
	L2Panel           string
	DMScheduledDate   string
	DMPanel           string
	OnboardingDate    string
	Comments          string
}

func (interview *Interview) NewInterview() *Interview {
	return &Interview{}
}
