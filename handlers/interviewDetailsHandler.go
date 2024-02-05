package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/ArpitChinmay/interview/db/datareader"
	dtomodels "github.com/ArpitChinmay/interview/main/dtoModels"
	"github.com/ArpitChinmay/interview/models"
	"github.com/gin-gonic/gin"
)

type InterviewHandler struct {
	context          context.Context
	interviewDetails []models.Interview
	reader           *datareader.DataReader
}

func NewInterviewHandler(c *gin.Context, db *sql.DB) *InterviewHandler {
	return &InterviewHandler{context: c, interviewDetails: make([]models.Interview, 0), reader: new(datareader.DataReader)}
}

func (handler *InterviewHandler) GetSelectedAndRejectedCandidatesAtLevelOne(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadInterviewDataForLevelOneSelecteOrRejected(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetSelectedAndRejectedCandidatesAtLevelTwo(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadInterviewDataForLevelTwoSelecteOrRejected(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetSelectedAndRejectedCandidatesAtLevelThree(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadInterviewDataForLevelThreeSelecteOrRejected(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetSelectedCandidatesAtLevelOne(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadInterviewDataForLevelOneSelecteCandidates(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetRejectedCandidatesAtLevelOne(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadInterviewDataForLevelOneRejectedCandidates(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetSelectedCandidatesAtLevelTwo(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadInterviewDataForLevelTwoSelectedCandidates(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetRejectedCandidatesAtLevelTwo(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadInterviewDataForLevelTwoRejectedCandidates(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetSelectedCandidatesAtDMLevel(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadInterviewDataForDMLevelSelectedCandidates(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetRejectedCandidatesAtDMLevel(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadInterviewDataForDMLevelRejectedCandidates(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetOnboardedCandidates(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadInterviewDataForOnboardedCandidates(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetCandidatesOfferedAndAccepted(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadCandidatesOfferedAndAcceptedPosition(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) GetCandidatesOfferedAndAwaited(c *gin.Context, db *sql.DB) ([]models.Interview, int, error) {
	var err error
	handler = NewInterviewHandler(c, db)
	handler.interviewDetails, err = handler.reader.ReadCandidatesOfferedAndAwaitedPosition(db)
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return handler.interviewDetails, len(handler.interviewDetails), nil
}

func (handler *InterviewHandler) CreateNewInterviewCandidate(c *gin.Context, db *sql.DB, candidate dtomodels.Candidate) (sql.Result, error) {
	// Convert the DTO model object to DB model object
	resumeModel := candidate.MapCandidateDetails(&candidate)
	response, err := handler.reader.CreateNewInterviewCandidate(db, resumeModel)

	if err != nil {
		log.Fatal(err)
	}
	return response, err
}

func (handler *InterviewHandler) UpdateInterviewCandidate(c *gin.Context, db *sql.DB, updateCandidate dtomodels.UpdateCandidate, candidateId string) (sql.Result, error) {
	response, err := handler.reader.UpdateInterviewCandidate(db, updateCandidate, candidateId)
	if err != nil {
		log.Fatal(err)
	}
	return response, err
}
