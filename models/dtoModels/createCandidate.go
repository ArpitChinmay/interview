package dtomodels

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/ArpitChinmay/interview/models"
	"github.com/rs/xid"
)

type Candidate struct {
	Candidate_ID        string `json:"candidate_id"`
	Name                string `json:"name"`
	Date                string `json:"date"`
	Skill_Category      string `json:"skill_category"`
	Mobile              string `json:"mobile"`
	Email_ID            string `json:"email_id"`
	Total_Experience    string `json:"total_experience"`
	Relevant_Experience string `json:"relevant_experience"`
	Current_Company     string `json:"current_company"`
	Notice_Period       string `json:"notice_period"`
	Comments            string `json:"comments"`
	Screening_Status    string `json:"screening_status"`
}

func (r Candidate) MapCandidateDetails(dtoModel *Candidate) models.Resume {
	var resumeModel models.Resume
	resumeModel.ResumeID = xid.New().String()
	resumeModel.CandidateID = dtoModel.Candidate_ID
	resumeModel.Name = dtoModel.Name

	date, err := time.Parse("2006-01-02", dtoModel.Date)
	if err == nil {
		resumeModel.Date = sql.NullTime{Time: date, Valid: true}
	} else {
		resumeModel.Date = sql.NullTime{Time: time.Now(), Valid: true}
	}

	resumeModel.Skill_Category = sql.NullString{String: dtoModel.Skill_Category, Valid: true}
	resumeModel.Mobile = sql.NullString{String: dtoModel.Mobile, Valid: true}
	resumeModel.Email_ID = sql.NullString{String: dtoModel.Email_ID, Valid: true}

	totalExperience, err := strconv.ParseFloat(dtoModel.Total_Experience, 64)
	if err == nil {
		resumeModel.Total_Experience = sql.NullFloat64{Float64: totalExperience, Valid: true}
	} else {
		resumeModel.Total_Experience = sql.NullFloat64{}
	}

	relevantExperience, err := strconv.ParseFloat(dtoModel.Relevant_Experience, 64)
	if err == nil {
		resumeModel.Relevant_Experience = sql.NullFloat64{Float64: relevantExperience, Valid: true}
	} else {
		resumeModel.Relevant_Experience = sql.NullFloat64{}
	}

	resumeModel.Current_Company = dtoModel.Current_Company

	noticePeriod, err := strconv.ParseInt(dtoModel.Notice_Period, 10, 32)
	if err == nil {
		resumeModel.Notice_Period = sql.NullInt32{Int32: int32(noticePeriod), Valid: true}
	} else {
		resumeModel.Notice_Period = sql.NullInt32{Int32: 90, Valid: true}
	}

	resumeModel.Comments = sql.NullString{String: dtoModel.Comments, Valid: true}
	resumeModel.Screening_Status = sql.NullString{String: dtoModel.Screening_Status, Valid: true}

	return resumeModel
}
