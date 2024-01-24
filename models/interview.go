package models

import "database/sql"

type Interview struct {
	InterviewStatusId sql.NullInt32
	CandidateId       sql.NullInt32
	InterviewStatus   sql.NullString
	L1ScheduledDate   sql.NullTime
	L1Panel           sql.NullString
	L2ScheduledDate   sql.NullTime
	L2Panel           sql.NullString
	DMScheduledDate   sql.NullTime
	DMPanel           sql.NullString
	OnboardingDate    sql.NullTime
	Comments          sql.NullString
}

func (interview *Interview) NewInterview() *Interview {
	return &Interview{}
}
