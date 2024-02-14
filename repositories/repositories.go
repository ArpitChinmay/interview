package repositories

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ArpitChinmay/interview/models"
	dtomodels "github.com/ArpitChinmay/interview/models/dtoModels"
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
	gET_SELECTED_REJECTED_CANDIDATE_AT_LEVEL_THREE = `
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
		Interview_Status = 'DM_Select'
		OR Interview_Status = 'DM_Reject';
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

	gET_ONBOARDED_CANDIDATE_DETAILS = `
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
    	Interview_Status = 'Onboarded';
	`

	gET_OFFERED_AND_ACCEPTED_CANDIDATE_DETAILS = `
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
    	Interview_Status = 'Offer_RolledOut_Accepted';
	`

	gET_OFFERED_AND_AWAITED_CANDIDATE_DETAILS = `
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
    	Interview_Status = 'Offer_RolledOut_Awaited'
	`

	iNSERT_NEW_INTERVIEW_CANDIDATE = `
	INSERT INTO Resume (
		ResumeId,
		CandidateId,
		Skill_Category,
		Name, 
		Mobile,
		Email_ID,
		Total_Experience,
		Relevant_Experience,
		Current_Company,
		Notice_Period,
		Comments,
		Screening_Status,
		Date) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	// gET_CANDIDATE_RESUME_DETAILS = `
	// SELECT * FROM int_db_data.resume WHERE candidateId = '?';
	// `

	gET_CANDIDATE_INTERVIEW_DETAILS = `
	SELECT * FROM int_db_data.interview_status WHERE candidateId = ?;`

	uPDATE_RESUME_SCREENING_STATUS = `UPDATE int_db_data.resume SET Screening_Status = ?, Date = ? WHERE CandidateId = ?;`

	uPDATE_INTERVIEW_DATA = `
	UPDATE 
		int_db_data.interview_status SET 
			Interview_status = ?, 
    		L1_Scheduled_Date = ?, 
    		L1_Panel = ?, 
    		L2_Scheduled_Date = ?, 
    		L2_Panel = ?, 
    		DM_Scheduled_Date = ?, 
    		DM_Panel = ? 
	WHERE  CandidateId = ?;`
)

type Repository interface {
	ReadInterviewDataForLevelOneSelecteOrRejected() (*[]models.Interview, error)
	ReadInterviewDataForLevelTwoSelecteOrRejected() (*[]models.Interview, error)
	ReadInterviewDataForLevelThreeSelecteOrRejected() (*[]models.Interview, error)
	ReadInterviewDataForLevelOneSelecteCandidates() (*[]models.Interview, error)
	ReadInterviewDataForLevelOneRejectedCandidates() (*[]models.Interview, error)
	ReadInterviewDataForLevelTwoSelectedCandidates() (*[]models.Interview, error)
	ReadInterviewDataForLevelTwoRejectedCandidates() (*[]models.Interview, error)
	ReadInterviewDataForDMLevelSelectedCandidates() (*[]models.Interview, error)
	ReadInterviewDataForDMLevelRejectedCandidates() (*[]models.Interview, error)
	ReadInterviewDataForOnboardedCandidates() (*[]models.Interview, error)
	ReadCandidatesOfferedAndAcceptedPosition() (*[]models.Interview, error)
	ReadCandidatesOfferedAndAwaitedPosition() (*[]models.Interview, error)
	CreateNewInterviewCandidate(*models.Resume) (sql.Result, error)
	UpdateInterviewCandidate(*dtomodels.UpdateCandidate, *string) (sql.Result, error)
}

type repository struct {
	database *sql.DB
}

func InitializeRepository(db *sql.DB) Repository {
	return &repository{database: db}
}

func (repo *repository) ReadInterviewDataForLevelOneSelecteOrRejected() (*[]models.Interview, error) {
	var interviewDetails []models.Interview
	log.Println("Attempting to read the data from database...")
	rows, err := repo.database.Query(gET_SELECTE_REJECTED_CANDIDATE_AT_LEVEL_ONE)

	if err != nil {
		log.Println("error occurred while executing the read query in database...")
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
			log.Println("error reading the data into row...")
			log.Fatal(err)
			continue
		}
		interviewDetails = append(interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return &interviewDetails, nil
}

func (repo *repository) ReadInterviewDataForLevelTwoSelecteOrRejected() (*[]models.Interview, error) {
	var interviewDetails []models.Interview
	log.Println("Attempting to read the data from database...")
	rows, err := repo.database.Query(gET_SELECTED_REJECTED_CANDIDATE_AT_LEVEL_TWO)

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
			continue
		}
		interviewDetails = append(interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return &interviewDetails, nil
}

func (repo *repository) ReadInterviewDataForLevelThreeSelecteOrRejected() (*[]models.Interview, error) {
	var interviewDetails []models.Interview
	log.Println("Attempting to read the data from database...")
	rows, err := repo.database.Query(gET_SELECTED_REJECTED_CANDIDATE_AT_LEVEL_THREE)

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
			continue
		}
		interviewDetails = append(interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return &interviewDetails, nil
}

func (repo *repository) ReadInterviewDataForLevelOneSelecteCandidates() (*[]models.Interview, error) {
	var interviewDetails []models.Interview
	log.Println("Attempting to read the data from database...")
	rows, err := repo.database.Query(gET_SELECTED_CANDIDATE_AT_LEVEL_ONE)

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
			continue
		}
		interviewDetails = append(interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return &interviewDetails, nil
}

func (repo *repository) ReadInterviewDataForLevelOneRejectedCandidates() (*[]models.Interview, error) {
	var interviewDetails []models.Interview
	log.Println("Attempting to read the data from database...")
	rows, err := repo.database.Query(gET_REJECTED_CANDIDATE_AT_LEVEL_ONE)

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
		interviewDetails = append(interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return &interviewDetails, nil
}

func (repo *repository) ReadInterviewDataForLevelTwoSelectedCandidates() (*[]models.Interview, error) {
	var interviewDetails []models.Interview
	log.Println("Attempting to read the data from database...")
	rows, err := repo.database.Query(gET_SELECTED_CANDIDATE_AT_LEVEL_TWO)

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
		interviewDetails = append(interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return &interviewDetails, nil
}

func (repo *repository) ReadInterviewDataForLevelTwoRejectedCandidates() (*[]models.Interview, error) {
	var interviewDetails []models.Interview
	log.Println("Attempting to read the data from database...")
	rows, err := repo.database.Query(gET_REJECTED_CANDIDATE_AT_LEVEL_TWO)

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
		interviewDetails = append(interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return &interviewDetails, nil
}

func (repo *repository) ReadInterviewDataForDMLevelSelectedCandidates() (*[]models.Interview, error) {
	var interviewDetails []models.Interview
	log.Println("Attempting to read the data from database...")
	rows, err := repo.database.Query(gET_SELECTED_CANDIDATE_AT_LEVEL_DM)

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
		interviewDetails = append(interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return &interviewDetails, nil
}

func (repo *repository) ReadInterviewDataForDMLevelRejectedCandidates() (*[]models.Interview, error) {
	var interviewDetails []models.Interview
	log.Println("Attempting to read the data from database...")
	rows, err := repo.database.Query(gET_REJECTED_CANDIDATE_AT_LEVEL_DM)

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
		interviewDetails = append(interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return &interviewDetails, nil
}

func (repo *repository) ReadInterviewDataForOnboardedCandidates() (*[]models.Interview, error) {
	var interviewDetails []models.Interview
	log.Println("Attempting to read the data from database...")
	rows, err := repo.database.Query(gET_ONBOARDED_CANDIDATE_DETAILS)

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
		interviewDetails = append(interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return &interviewDetails, nil
}

func (repo *repository) ReadCandidatesOfferedAndAcceptedPosition() (*[]models.Interview, error) {
	var interviewDetails []models.Interview
	log.Println("Attempting to read the data from database...")
	rows, err := repo.database.Query(gET_OFFERED_AND_ACCEPTED_CANDIDATE_DETAILS)

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
		interviewDetails = append(interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return &interviewDetails, nil
}

func (repo *repository) ReadCandidatesOfferedAndAwaitedPosition() (*[]models.Interview, error) {
	var interviewDetails []models.Interview
	log.Println("Attempting to read the data from database...")
	rows, err := repo.database.Query(gET_OFFERED_AND_AWAITED_CANDIDATE_DETAILS)

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
		interviewDetails = append(interviewDetails, candidate)
	}

	if err = rows.Err(); err != nil {
		log.Println("Some error: ")
		log.Println(err)
	}

	defer rows.Close()
	return &interviewDetails, nil
}

func (repo *repository) CreateNewInterviewCandidate(resume *models.Resume) (sql.Result, error) {
	log.Println("Attempting to write the data to database...")

	response, err := repo.database.Exec(iNSERT_NEW_INTERVIEW_CANDIDATE, resume.ResumeID, resume.CandidateID, resume.Skill_Category,
		resume.Name, resume.Mobile, resume.Email_ID, resume.Total_Experience, resume.Relevant_Experience, resume.Current_Company,
		resume.Notice_Period, resume.Comments, resume.Screening_Status, resume.Date)

	if err != nil {
		log.Println("There was an error inserting data into the database...")
		log.Fatal(err)
	}
	return response, err
}

func (repo *repository) UpdateInterviewCandidate(updateCandidate *dtomodels.UpdateCandidate, candidateId *string) (sql.Result, error) {
	log.Println("Attempting to update the data in database...")

	responseResumeUpdate, err := repo.database.Exec(uPDATE_RESUME_SCREENING_STATUS, updateCandidate.ScreeningStatus, updateCandidate.Date, candidateId)
	log.Println("Resume update successfull")

	responseInterviewUpdate, err2 := repo.database.Exec(uPDATE_INTERVIEW_DATA, updateCandidate.InterviewStatus, updateCandidate.L1Date, updateCandidate.L1Panel, updateCandidate.L2Date, updateCandidate.L2Panel,
		updateCandidate.DMDate, updateCandidate.DMPanel, candidateId)
	log.Println("Interview update successful...")
	if err != nil || err2 != nil {
		log.Println("There was an error updating data into the database...")
		log.Println(err)
		log.Println(err2)
		log.Println("Resume update response: ", responseResumeUpdate)
		log.Println("Interview update resposne: ", responseInterviewUpdate)
	}

	interviewData, err := repo.database.Exec(gET_CANDIDATE_INTERVIEW_DETAILS, candidateId)
	return interviewData, err
}

// type DataReader struct {
// 	database         *sql.DB
// 	interviewDetails []models.Interview
// }

// type DataReader interface {
// 	ReadInterviewDataForLevelOneSelecteOrRejected() (*[]models.Interview, error)
// 	ReadInterviewDataForLevelTwoSelecteOrRejected() (*[]models.Interview, error)
// 	ReadInterviewDataForLevelOneSelecteCandidates() (*[]models.Interview, error)
// 	ReadInterviewDataForLevelOneRejectedCandidates() (*[]models.Interview, error)
// 	ReadInterviewDataForLevelTwoSelectedCandidates() (*[]models.Interview, error)
// 	ReadInterviewDataForLevelTwoRejectedCandidates() (*[]models.Interview, error)
// 	ReadInterviewDataForDMLevelSelectedCandidates() (*[]models.Interview, error)
// 	ReadInterviewDataForDMLevelRejectedCandidates() (*[]models.Interview, error)
// }

// func InitializeDataReader(db *sql.DB) DataReader {
// 	return &dataReader{database: db, interviewDetails: make([]models.Interview, 0)}
// }

// func (datareader *dataReader) ReadInterviewDataForLevelOneSelecteOrRejected() (*[]models.Interview, error) {
// 	// datareader = NewDataReader(db)
// 	log.Println("Attempting to read the data from database...")
// 	rows, err := datareader.database.Query(gET_SELECTE_REJECTED_CANDIDATE_AT_LEVEL_ONE)

// 	if err != nil {
// 		log.Println("error occurred while reading the database...")
// 		log.Fatal(err)
// 		return nil, errors.New("There was an error encountered while trying to read the database...")
// 	}

// 	log.Println("rows data:")
// 	for rows.Next() {
// 		candidate := models.Interview{}
// 		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
// 			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
// 			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
// 			&candidate.Comments)

// 		if err != nil {
// 			log.Println("error reading the data into rows...")
// 			log.Fatal(err)
// 			return nil, errors.New("There was an error reading the data from rows...")
// 		}
// 		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
// 	}

// 	if err = rows.Err(); err != nil {
// 		log.Println("Some error: ")
// 		log.Println(err)
// 	}

// 	defer rows.Close()
// 	return &datareader.interviewDetails, nil
// }

// func (datareader *dataReader) ReadInterviewDataForLevelTwoSelecteOrRejected() (*[]models.Interview, error) {
// 	//datareader = NewDataReader(db)
// 	log.Println("Attempting to read the data from database...")
// 	rows, err := datareader.database.Query(gET_SELECTED_REJECTED_CANDIDATE_AT_LEVEL_TWO)

// 	if err != nil {
// 		log.Println("error occurred while reading the database...")
// 		log.Fatal(err)
// 		return nil, errors.New("There was an error encountered while trying to read the database...")
// 	}

// 	log.Println("rows data:")
// 	for rows.Next() {
// 		candidate := models.Interview{}
// 		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
// 			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
// 			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
// 			&candidate.Comments)

// 		if err != nil {
// 			log.Println("error reading the data into rows...")
// 			log.Fatal(err)
// 			return nil, errors.New("There was an error reading the data from rows...")
// 		}
// 		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
// 	}

// 	if err = rows.Err(); err != nil {
// 		log.Println("Some error: ")
// 		log.Println(err)
// 	}

// 	defer rows.Close()
// 	return &datareader.interviewDetails, nil
// }

// func (datareader *dataReader) ReadInterviewDataForLevelOneSelecteCandidates() (*[]models.Interview, error) {
// 	// datareader = NewDataReader(db)
// 	log.Println("Attempting to read the data from database...")
// 	rows, err := datareader.database.Query(gET_SELECTED_CANDIDATE_AT_LEVEL_ONE)

// 	if err != nil {
// 		log.Println("error occurred while reading the database...")
// 		log.Fatal(err)
// 		return nil, errors.New("There was an error encountered while trying to read the database...")
// 	}

// 	log.Println("rows data:")
// 	for rows.Next() {
// 		candidate := models.Interview{}
// 		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
// 			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
// 			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
// 			&candidate.Comments)

// 		if err != nil {
// 			log.Println("error reading the data into rows...")
// 			log.Fatal(err)
// 			return nil, errors.New("There was an error reading the data from rows...")
// 		}
// 		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
// 	}

// 	if err = rows.Err(); err != nil {
// 		log.Println("Some error: ")
// 		log.Println(err)
// 	}

// 	defer rows.Close()
// 	return &datareader.interviewDetails, nil
// }

// func (datareader *dataReader) ReadInterviewDataForLevelOneRejectedCandidates() (*[]models.Interview, error) {
// 	// datareader = NewDataReader(db)
// 	log.Println("Attempting to read the data from database...")
// 	rows, err := datareader.database.Query(gET_REJECTED_CANDIDATE_AT_LEVEL_ONE)

// 	if err != nil {
// 		log.Println("error occurred while reading the database...")
// 		log.Fatal(err)
// 		return nil, errors.New("There was an error encountered while trying to read the database...")
// 	}

// 	log.Println("rows data:")
// 	for rows.Next() {
// 		candidate := models.Interview{}
// 		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
// 			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
// 			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
// 			&candidate.Comments)

// 		if err != nil {
// 			log.Println("error reading the data into rows...")
// 			log.Fatal(err)
// 			return nil, errors.New("There was an error reading the data from rows...")
// 		}
// 		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
// 	}

// 	if err = rows.Err(); err != nil {
// 		log.Println("Some error: ")
// 		log.Println(err)
// 	}

// 	defer rows.Close()
// 	return &datareader.interviewDetails, nil
// }

// func (datareader *dataReader) ReadInterviewDataForLevelTwoSelectedCandidates() (*[]models.Interview, error) {
// 	// datareader = NewDataReader(db)
// 	log.Println("Attempting to read the data from database...")
// 	rows, err := datareader.database.Query(gET_SELECTED_CANDIDATE_AT_LEVEL_TWO)

// 	if err != nil {
// 		log.Println("error occurred while reading the database...")
// 		log.Fatal(err)
// 		return nil, errors.New("There was an error encountered while trying to read the database...")
// 	}

// 	log.Println("rows data:")
// 	for rows.Next() {
// 		candidate := models.Interview{}
// 		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
// 			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
// 			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
// 			&candidate.Comments)

// 		if err != nil {
// 			log.Println("error reading the data into rows...")
// 			log.Fatal(err)
// 			return nil, errors.New("There was an error reading the data from rows...")
// 		}
// 		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
// 	}

// 	if err = rows.Err(); err != nil {
// 		log.Println("Some error: ")
// 		log.Println(err)
// 	}

// 	defer rows.Close()
// 	return &datareader.interviewDetails, nil
// }

// func (datareader *dataReader) ReadInterviewDataForLevelTwoRejectedCandidates() (*[]models.Interview, error) {
// 	// datareader = NewDataReader(db)
// 	log.Println("Attempting to read the data from database...")
// 	rows, err := datareader.database.Query(gET_REJECTED_CANDIDATE_AT_LEVEL_TWO)

// 	if err != nil {
// 		log.Println("error occurred while reading the database...")
// 		log.Fatal(err)
// 		return nil, errors.New("There was an error encountered while trying to read the database...")
// 	}

// 	log.Println("rows data:")
// 	for rows.Next() {
// 		candidate := models.Interview{}
// 		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
// 			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
// 			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
// 			&candidate.Comments)

// 		if err != nil {
// 			log.Println("error reading the data into rows...")
// 			log.Fatal(err)
// 			return nil, errors.New("There was an error reading the data from rows...")
// 		}
// 		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
// 	}

// 	if err = rows.Err(); err != nil {
// 		log.Println("Some error: ")
// 		log.Println(err)
// 	}

// 	defer rows.Close()
// 	return &datareader.interviewDetails, nil
// }

// func (datareader *dataReader) ReadInterviewDataForDMLevelSelectedCandidates() (*[]models.Interview, error) {
// 	// datareader = NewDataReader(db)
// 	log.Println("Attempting to read the data from database...")
// 	rows, err := datareader.database.Query(gET_SELECTED_CANDIDATE_AT_LEVEL_DM)

// 	if err != nil {
// 		log.Println("error occurred while reading the database...")
// 		log.Fatal(err)
// 		return nil, errors.New("There was an error encountered while trying to read the database...")
// 	}

// 	log.Println("rows data:")
// 	for rows.Next() {
// 		candidate := models.Interview{}
// 		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
// 			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
// 			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
// 			&candidate.Comments)

// 		if err != nil {
// 			log.Println("error reading the data into rows...")
// 			log.Fatal(err)
// 			return nil, errors.New("There was an error reading the data from rows...")
// 		}
// 		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
// 	}

// 	if err = rows.Err(); err != nil {
// 		log.Println("Some error: ")
// 		log.Println(err)
// 	}

// 	defer rows.Close()
// 	return &datareader.interviewDetails, nil
// }

// func (datareader *dataReader) ReadInterviewDataForDMLevelRejectedCandidates() (*[]models.Interview, error) {
// 	// datareader = NewDataReader(db)
// 	log.Println("Attempting to read the data from database...")
// 	rows, err := datareader.database.Query(gET_REJECTED_CANDIDATE_AT_LEVEL_DM)

// 	if err != nil {
// 		log.Println("error occurred while reading the database...")
// 		log.Fatal(err)
// 		return nil, errors.New("There was an error encountered while trying to read the database...")
// 	}

// 	log.Println("rows data:")
// 	for rows.Next() {
// 		candidate := models.Interview{}
// 		err = rows.Scan(&candidate.InterviewStatusId, &candidate.CandidateId, &candidate.InterviewStatus,
// 			&candidate.L1ScheduledDate, &candidate.L1Panel, &candidate.L2ScheduledDate,
// 			&candidate.L2Panel, &candidate.DMScheduledDate, &candidate.DMPanel, &candidate.OnboardingDate,
// 			&candidate.Comments)

// 		if err != nil {
// 			log.Println("error reading the data into rows...")
// 			log.Fatal(err)
// 			return nil, errors.New("There was an error reading the data from rows...")
// 		}
// 		datareader.interviewDetails = append(datareader.interviewDetails, candidate)
// 	}

// 	if err = rows.Err(); err != nil {
// 		log.Println("Some error: ")
// 		log.Println(err)
// 	}

// 	defer rows.Close()
// 	return &datareader.interviewDetails, nil
// }
