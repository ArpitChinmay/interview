package dtomodels

import (
	"interview-dashboard/models"
	"time"
)

type InterviewDTO struct {
	CandidateId     string    `json:"candidateId"`
	InterviewStatus string    `json:"interviewStatus"`
	L1Date          time.Time `json:"levelOneDate"`
	L1Panel         string    `json:"levelOnePanel"`
	L2Date          time.Time `json:"levelTwoDate"`
	L2Panel         string    `json:"levelTwoPanel"`
	DMDate          time.Time `json:"dMDate"`
	DMPanel         string    `json:"dMPanel"`
	OnboardingDate  time.Time `json:"onboardingDate"`
	Comments        string    `json:"comments"`
}

func (r InterviewDTO) MapInterviewDetails(databaseModel *models.Interview) InterviewDTO {
	result := InterviewDTO{}
	result.CandidateId = databaseModel.CandidateId
	result.InterviewStatus = databaseModel.InterviewStatus

	l1date, err := time.Parse("YYYY-MM-DD HH:MM:SS", databaseModel.L1ScheduledDate)
	if err != nil {
		result.L1Date = time.Time{}
	} else {
		result.L1Date = l1date
	}
	result.L1Panel = databaseModel.L1Panel

	l2date, err := time.Parse("YYYY-MM-DD HH:MM:SS", databaseModel.L2ScheduledDate)
	if err != nil {
		result.L2Date = time.Time{}
	} else {
		result.L2Date = l2date
	}
	result.L2Panel = databaseModel.L2Panel

	dmdate, err := time.Parse("YYYY-MM-DD HH:MM:SS", databaseModel.DMScheduledDate)
	if err != nil {
		result.DMDate = time.Time{}
	} else {
		result.DMDate = dmdate
	}
	result.DMPanel = databaseModel.DMPanel

	onboardingDate, err := time.Parse("YYYY-MM-DD HH:MM:SS", databaseModel.OnboardingDate)
	if err != nil {
		result.OnboardingDate = time.Time{}
	} else {
		result.OnboardingDate = onboardingDate
	}

	result.Comments = databaseModel.Comments

	return result
}
