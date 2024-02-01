package dtomodels

import (
	"time"

	"github.com/ArpitChinmay/interview/models"
)

type InterviewDTO struct {
	CandidateId     int       `json:"candidateId"`
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
	result.CandidateId = int(databaseModel.CandidateId.Int32)
	result.InterviewStatus = databaseModel.InterviewStatus.String

	if databaseModel.L1ScheduledDate.Valid {
		result.L1Date = databaseModel.L1ScheduledDate.Time
	} else {
		result.L1Date = time.Time{}
	}
	result.L1Panel = databaseModel.L1Panel.String

	if databaseModel.L2ScheduledDate.Valid {
		result.L2Date = databaseModel.L2ScheduledDate.Time
	} else {
		result.L2Date = time.Time{}
	}
	result.L2Panel = databaseModel.L2Panel.String

	if databaseModel.DMScheduledDate.Valid {
		result.DMDate = databaseModel.DMScheduledDate.Time
	} else {
		result.DMDate = time.Time{}
	}
	result.DMPanel = databaseModel.DMPanel.String

	if databaseModel.OnboardingDate.Valid {
		result.OnboardingDate = databaseModel.OnboardingDate.Time
	} else {
		result.OnboardingDate = time.Time{}
	}

	result.Comments = databaseModel.Comments.String

	return result
}
