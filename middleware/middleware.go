package handlers

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ArpitChinmay/interview/models"
	dtomodels "github.com/ArpitChinmay/interview/models/dtoModels"
	repository "github.com/ArpitChinmay/interview/repositories"
	"github.com/gin-gonic/gin"
)

type interviewMiddleware struct {
	repository repositories.Repository
}

type InterviewMiddleware interface {
	GetSelectedAndRejectedCandidatesAtLevelOne(c *gin.Context) (*[]models.Interview, int, error)
	GetSelectedAndRejectedCandidatesAtLevelTwo(c *gin.Context) (*[]models.Interview, int, error)
	GetSelectedAndRejectedCandidatesAtLevelThree(c *gin.Context) (*[]models.Interview, int, error)
	GetSelectedCandidatesAtLevelOne(c *gin.Context) (*[]models.Interview, int, error)
	GetRejectedCandidatesAtLevelOne(c *gin.Context) (*[]models.Interview, int, error)
	GetSelectedCandidatesAtLevelTwo(c *gin.Context) (*[]models.Interview, int, error)
	GetRejectedCandidatesAtLevelTwo(c *gin.Context) (*[]models.Interview, int, error)
	GetSelectedCandidatesAtDMLevel(c *gin.Context) (*[]models.Interview, int, error)
	GetRejectedCandidatesAtDMLevel(c *gin.Context) (*[]models.Interview, int, error)
	GetOnboardedCandidates(c *gin.Context) (*[]models.Interview, int, error)
	GetCandidatesOfferedAndAccepted(c *gin.Context) (*[]models.Interview, int, error)
	GetCandidatesOfferedAndAwaited(c *gin.Context) (*[]models.Interview, int, error)
	CreateNewInterviewCandidate(c *gin.Context, candidate dtomodels.Candidate) (*sql.Result, error)
	UpdateInterviewCandidate(c *gin.Context, updateCandidate dtomodels.UpdateCandidate, candidateId string) (*sql.Result, error)
}

func InitializeInterviewMiddleware(repository repository.Repository) InterviewMiddleware {
	return &interviewMiddleware{repository}
}

func (midware *interviewMiddleware) GetSelectedAndRejectedCandidatesAtLevelOne(c *gin.Context) (*[]models.Interview, int, error) {
	interviewDetails, err := midware.repository.ReadInterviewDataForLevelOneSelecteOrRejected()
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return interviewDetails, len(*interviewDetails), nil
}

func (midware *interviewMiddleware) GetSelectedAndRejectedCandidatesAtLevelTwo(c *gin.Context) (*[]models.Interview, int, error) {
	interviewDetails, err := midware.repository.ReadInterviewDataForLevelTwoSelecteOrRejected()
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return interviewDetails, len(*interviewDetails), nil
}

func (midware *interviewMiddleware) GetSelectedAndRejectedCandidatesAtLevelThree(c *gin.Context) (*[]models.Interview, int, error) {
	interviewDetails, err := midware.repository.ReadInterviewDataForLevelThreeSelecteOrRejected()
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return interviewDetails, len(*interviewDetails), nil
}

func (midware *interviewMiddleware) GetSelectedCandidatesAtLevelOne(c *gin.Context) (*[]models.Interview, int, error) {
	interviewDetails, err := midware.repository.ReadInterviewDataForLevelOneSelecteCandidates()
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return interviewDetails, len(*interviewDetails), nil
}

func (midware *interviewMiddleware) GetRejectedCandidatesAtLevelOne(c *gin.Context) (*[]models.Interview, int, error) {
	interviewDetails, err := midware.repository.ReadInterviewDataForLevelOneRejectedCandidates()
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return interviewDetails, len(*interviewDetails), nil
}

func (midware *interviewMiddleware) GetSelectedCandidatesAtLevelTwo(c *gin.Context) (*[]models.Interview, int, error) {
	interviewDetails, err := midware.repository.ReadInterviewDataForLevelTwoSelectedCandidates()
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return interviewDetails, len(*interviewDetails), nil
}

func (midware *interviewMiddleware) GetRejectedCandidatesAtLevelTwo(c *gin.Context) (*[]models.Interview, int, error) {
	interviewDetails, err := midware.repository.ReadInterviewDataForLevelTwoRejectedCandidates()
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return interviewDetails, len(*interviewDetails), nil
}

func (midware *interviewMiddleware) GetSelectedCandidatesAtDMLevel(c *gin.Context) (*[]models.Interview, int, error) {
	interviewDetails, err := midware.repository.ReadInterviewDataForDMLevelSelectedCandidates()
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return interviewDetails, len(*interviewDetails), nil
}

func (midware *interviewMiddleware) GetRejectedCandidatesAtDMLevel(c *gin.Context) (*[]models.Interview, int, error) {
	interviewDetails, err := midware.repository.ReadInterviewDataForDMLevelRejectedCandidates()
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return interviewDetails, len(*interviewDetails), nil
}

func (midware *interviewMiddleware) GetOnboardedCandidates(c *gin.Context) (*[]models.Interview, int, error) {
	interviewDetails, err := midware.repository.ReadInterviewDataForOnboardedCandidates()
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return interviewDetails, len(*interviewDetails), nil
}

func (midware *interviewMiddleware) GetCandidatesOfferedAndAccepted(c *gin.Context) (*[]models.Interview, int, error) {
	interviewDetails, err := midware.repository.ReadCandidatesOfferedAndAcceptedPosition()
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return interviewDetails, len(*interviewDetails), nil
}

func (midware *interviewMiddleware) GetCandidatesOfferedAndAwaited(c *gin.Context) (*[]models.Interview, int, error) {
	interviewDetails, err := midware.repository.ReadCandidatesOfferedAndAwaitedPosition()
	if err != nil {
		log.Fatal(err)
		return nil, 0, errors.New("An error occurred while fetching the data")
	}
	return interviewDetails, len(*interviewDetails), nil
}

func (midware *interviewMiddleware) CreateNewInterviewCandidate(c *gin.Context, candidate dtomodels.Candidate) (*sql.Result, error) {
	// Convert the DTO model object to DB model object
	resumeModel := candidate.MapCandidateDetails(&candidate)
	response, err := midware.repository.CreateNewInterviewCandidate(&resumeModel)

	if err != nil {
		log.Fatal(err)
	}
	return &response, err
}

func (midware *interviewMiddleware) UpdateInterviewCandidate(c *gin.Context, updateCandidate dtomodels.UpdateCandidate, candidateId string) (*sql.Result, error) {
	response, err := midware.repository.UpdateInterviewCandidate(&updateCandidate, &candidateId)
	if err != nil {
		log.Fatal(err)
	}
	return &response, err
}
