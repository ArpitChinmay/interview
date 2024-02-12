package models

import "database/sql"

type Resume struct {
	ResumeID            string
	CandidateID         string
	Skill_Category      sql.NullString
	Name                string
	Mobile              sql.NullString
	Email_ID            sql.NullString
	Total_Experience    sql.NullFloat64
	Relevant_Experience sql.NullFloat64
	Current_Company     string
	Notice_Period       sql.NullInt32
	Comments            sql.NullString
	Screening_Status    sql.NullString
	Date                sql.NullTime
}

func (resume *Resume) NewResume() *Resume {
	return &Resume{}
}
