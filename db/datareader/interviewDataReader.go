package datareader

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ArpitChinmay/interview/models"
)

const (
	gET_SELECTE_REJECTED_CANDIDATE_AT_LEVEL_ONE = `
	SELECT 
		InterviewStatusId, 
		CandidateId, 
		Interview_Status, 
		L1_Scheduled_Date, 
		L1_Panel, 
		L2_Scheduled_Date, 
		L2_Panel, 
		DM_Scheduled_Date, 
		DM_Panel, 
		Onboarding_Date, 
		Comments
	FROM 
		int_db_data.interview_status
	WHERE 
		Interview_Status = 'L1_Select'
		OR Interview_Status = 'L1_Reject';`

	gET_SELECTED_REJECTED_CANDIDATE_AT_LEVEL_TWO = `
	SELECT 
    	InterviewStatusId,
    	CandidateId,
    	Interview_Status,
    	L1_Scheduled_Date,
    	L1_Panel,
    	L2_Scheduled_Date,
    	L2_Panel,
    	DM_Scheduled_Date,
    	DM_Panel,
    	Onboarding_Date,
    	Comments
	FROM
    	int_db_data.interview_status
	WHERE
    	Interview_Status = 'L2_Reject'
    	OR Interview_Status = 'L2_Select';
	`
	gET_SELECTED_CANDIDATE_AT_LEVEL_ONE = `
	SELECT 
		InterviewStatusId, 
		CandidateId, 
		Interview_Status, 
		L1_Scheduled_Date, 
		L1_Panel, 
		L2_Scheduled_Date, 
		L2_Panel, 
		DM_Scheduled_Date, 
		DM_Panel, 
		Onboarding_Date, 
		Comments
	FROM 
		int_db_data.interview_status
	WHERE 
		Interview_Status = 'L1_Select';`

	gET_REJECTED_CANDIDATE_AT_LEVEL_ONE = `
	SELECT 
		InterviewStatusId, 
		CandidateId, 
		Interview_Status, 
		L1_Scheduled_Date, 
		L1_Panel, 
		L2_Scheduled_Date, 
		L2_Panel, 
		DM_Scheduled_Date, 
		DM_Panel, 
		Onboarding_Date, 
		Comments
	FROM 
		int_db_data.interview_status
	WHERE 
		Interview_Status = 'L1_Reject';
	`

	gET_SELECTED_CANDIDATE_AT_LEVEL_TWO = `
	SELECT 
		InterviewStatusId, 
		CandidateId, 
		Interview_Status, 
		L1_Scheduled_Date, 
		L1_Panel, 
		L2_Scheduled_Date, 
		L2_Panel, 
		DM_Scheduled_Date, 
		DM_Panel, 
		Onboarding_Date, 
		Comments
	FROM 
		int_db_data.interview_status
	WHERE 
		Interview_Status = 'L2_Select';`

	gET_REJECTED_CANDIDATE_AT_LEVEL_TWO = `
	SELECT 
    	InterviewStatusId,
    	CandidateId,
    	Interview_Status,
    	L1_Scheduled_Date,
    	L1_Panel,
    	L2_Scheduled_Date,
    	L2_Panel,
    	DM_Scheduled_Date,
    	DM_Panel,
    	Onboarding_Date,
    	Comments
	FROM
    	int_db_data.interview_status
	WHERE
    	Interview_Status = 'L2_Reject';`

	gET_SELECTED_CANDIDATE_AT_LEVEL_DM = `
	SELECT 
    	InterviewStatusId,
    	CandidateId,
    	Interview_Status,
    	L1_Scheduled_Date,
    	L1_Panel,
    	L2_Scheduled_Date,
    	L2_Panel,
    	DM_Scheduled_Date,
    	DM_Panel,
    	Onboarding_Date,
    	Comments
	FROM
    	int_db_data.interview_status
	WHERE
    	Interview_Status = 'DM_Select';`

	gET_REJECTED_CANDIDATE_AT_LEVEL_DM = `
	SELECT 
    	InterviewStatusId,
    	CandidateId,
    	Interview_Status,
    	L1_Scheduled_Date,
    	L1_Panel,
    	L2_Scheduled_Date,
    	L2_Panel,
    	DM_Scheduled_Date,
    	DM_Panel,
    	Onboarding_Date,
    	Comments
	FROM
    	int_db_data.interview_status
	WHERE
    	Interview_Status = 'DM_Reject';`
)

type DataReader struct {
	database         *sql.DB
	interviewDetails []models.Interview
}

func NewDataReader(db *sql.DB) *DataReader {
	return &DataReader{database: db, interviewDetails: make([]models.Interview, 0)}
}

func (datareader *DataReader) ReadInterviewDataForLevelOneSelecteOrRejected(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_SELECTE_REJECTED_CANDIDATE_AT_LEVEL_ONE)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
			&candidate.Comments)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadInterviewDataForLevelTwoSelecteOrRejected(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_SELECTED_REJECTED_CANDIDATE_AT_LEVEL_TWO)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
			&candidate.Comments)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadInterviewDataForLevelOneSelecteCandidates(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_SELECTED_CANDIDATE_AT_LEVEL_ONE)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
			&candidate.Comments)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadInterviewDataForLevelOneRejectedCandidates(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_REJECTED_CANDIDATE_AT_LEVEL_ONE)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
			&candidate.Comments)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadInterviewDataForLevelTwoSelectedCandidates(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_SELECTED_CANDIDATE_AT_LEVEL_TWO)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
			&candidate.Comments)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadInterviewDataForLevelTwoRejectedCandidates(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_REJECTED_CANDIDATE_AT_LEVEL_TWO)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
			&candidate.Comments)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadInterviewDataForDMLevelSelectedCandidates(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_SELECTED_CANDIDATE_AT_LEVEL_DM)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
			&candidate.Comments)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}

func (datareader *DataReader) ReadInterviewDataForDMLevelRejectedCandidates(db *sql.DB) ([]models.Interview, error) {
	datareader = NewDataReader(db)
	log.Println("Attempting to read the data from database...")
	rows, err := datareader.database.Query(gET_REJECTED_CANDIDATE_AT_LEVEL_DM)

	if err != nil {
		log.Println("error occurred while reading the database...")
		log.Fatal(err)
		return nil, errors.New("There was an error encountered while trying to read the database...")
	}

	log.Println("rows data:")
	for rows.Next() {
		candidate := models.Interview{}
		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
			&candidate.Comments)

		if err != nil {
			log.Println("error reading the data into rows...")
			log.Fatal(err)
			return nil, errors.New("There was an error reading the data from rows...")
		}
		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return datareader.interviewDetails, nil
}
