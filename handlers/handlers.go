package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	middleware "github.com/ArpitChinmay/interview/middleware"
	dtomodels "github.com/ArpitChinmay/interview/models/dtoModels"
	"github.com/gin-gonic/gin"
)

type interviewHandler struct {
	midware middleware.InterviewMiddleware
}

type InterviewHandler interface {
	GetSelectedAndRejectedCandidates(c *gin.Context)
	GetSepecificCandidateDetails(c *gin.Context)
	GetOnboardedCandidateDetails(c *gin.Context)
	GetCandidatesWithAcceptedOffers(c *gin.Context)
	GetCandidatesWithAwaitedOffers(c *gin.Context)
	GetAcceptedCandidatesCount(c *gin.Context)
	GetAwaitedCandidatesCount(c *gin.Context)
	AddCandidate(c *gin.Context)
	UpdateCandidate(c *gin.Context)
}

func InitializeHandler(midware middleware.InterviewMiddleware) InterviewHandler {
	return &interviewHandler{midware}
}

// Arpit Chinmay & Shaik Saisameer
func (handler interviewHandler) GetSelectedAndRejectedCandidates(c *gin.Context) {
	level, err := strconv.ParseInt(c.Query("level"), 0, 32)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading url params..."})
	}

	if level == 1 {
		// Excel file defines this method as L1
		detailsOfCandidatesDTO, _, err := handler.getCandidateInterviewDetailsAtLevelOne(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		c.JSON(http.StatusOK, detailsOfCandidatesDTO)
	} else if level == 2 {
		// Excel file defines this method as L2
		detailsOfCandidatesDTO, _, err := handler.getCandidateInterviewDetailsAtLevelTwo(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		c.JSON(http.StatusOK, detailsOfCandidatesDTO)
	} else if level == 3 {
		detailsOfCandidatesDTO, _, err := handler.getCandidateInterviewDetailsAtLevelThree(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		c.JSON(http.StatusOK, detailsOfCandidatesDTO)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong value for level param..."})
	}
}

// Arpit Chinmay & Shaik Saisameer
func (handler interviewHandler) GetSepecificCandidateDetails(c *gin.Context) {
	level, err := strconv.ParseInt(c.Param("level"), 0, 32)
	selected, err2 := strconv.ParseBool(c.Query("selected"))
	count, err3 := strconv.ParseBool(c.Query("count"))
	if err != nil || err2 != nil || err3 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "problem reading url params..."})
	}

	if level == 1 {
		if selected {
			if count {
				_, datacount, err := handler.getSelectedCandidateInterviewDetailsAtLevelOne(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, datacount)
			} else {
				detailsOfCandidatesDTO, _, err := handler.getSelectedCandidateInterviewDetailsAtLevelOne(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			}

		} else {
			if count {
				_, datacount, err := handler.getRejectedCandidateInterviewDetailsAtLevelOne(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, datacount)
			} else {
				detailsOfCandidatesDTO, _, err := handler.getRejectedCandidateInterviewDetailsAtLevelOne(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			}
		}
	} else if level == 2 {
		if selected {
			if count {
				_, datacount, err := handler.getSelectedCandidateInterviewDetailsAtLevelTwo(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, datacount)
			} else {
				detailsOfCandidatesDTO, _, err := handler.getSelectedCandidateInterviewDetailsAtLevelTwo(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			}
		} else {
			if count {
				_, datacount, err := handler.getRejectedCandidateInterviewDetailsAtLevelTwo(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, datacount)
			} else {
				detailsOfCandidatesDTO, _, err := handler.getRejectedCandidateInterviewDetailsAtLevelTwo(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			}
		}
	} else if level == 3 {
		if selected {
			if count {
				_, datacount, err := handler.getSelectedCandidateInterviewDetailsAtLevelThree(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, datacount)
			} else {
				detailsOfCandidatesDTO, _, err := handler.getSelectedCandidateInterviewDetailsAtLevelThree(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			}
		} else {
			if count {
				_, datacount, err := handler.getRejectedCandidateInterviewDetailsAtLevelThree(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, datacount)
			} else {
				detailsOfCandidatesDTO, _, err := handler.getRejectedCandidateInterviewDetailsAtLevelThree(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong request params..."})
	}
}

// Yellaling
func (handler *interviewHandler) GetOnboardedCandidateDetails(c *gin.Context) {
	//onboarded, err1 := strconv.ParseBool(c.Query("onboarded"))
	count, err2 := strconv.ParseBool(c.Query("count"))
	if err2 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "problem reading url params..."})
	}
	if count {
		_, datacount, err := handler.getOnboardedCandidateInterviewDetails(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		c.JSON(http.StatusOK, datacount)
	} else {
		detailsOfCandidatesDTO, _, err := handler.getOnboardedCandidateInterviewDetails(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		c.JSON(http.StatusOK, detailsOfCandidatesDTO)
	}
}

// Kundan Kumar
func (handler *interviewHandler) GetCandidatesWithAcceptedOffers(c *gin.Context) {
	detailsOfCandidatesDTO, _, err := handler.getCandidatesWithAcceptedOffers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
	}
	c.JSON(http.StatusOK, detailsOfCandidatesDTO)
}

func (handler *interviewHandler) GetCandidatesWithAwaitedOffers(c *gin.Context) {
	detailsOfCandidatesDTO, _, err := handler.getCandidatesWithAwaitedOffers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
	}
	c.JSON(http.StatusOK, detailsOfCandidatesDTO)
}

func (handler *interviewHandler) GetAcceptedCandidatesCount(c *gin.Context) {
	_, count, err := handler.getCandidatesWithAcceptedOffers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
	}
	c.JSON(http.StatusOK, count)
}

func (handler *interviewHandler) GetAwaitedCandidatesCount(c *gin.Context) {
	_, count, err := handler.getCandidatesWithAwaitedOffers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
	}
	c.JSON(http.StatusOK, count)
}

// SindhuShree KN
func (handler *interviewHandler) AddCandidate(c *gin.Context) {
	log.Println("Are we even getting here?")
	var candidate dtomodels.Candidate
	if err := c.ShouldBindJSON(&candidate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	sqlResult, err := handler.createNewInterviewCandidate(c, candidate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, sqlResult)
	}
}

func (handler *interviewHandler) UpdateCandidate(c *gin.Context) {
	candidateId := c.Param("id")
	var updateCandidate dtomodels.UpdateCandidate
	if err := c.ShouldBindJSON(&updateCandidate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	sqlResult, err := handler.updateInterviewCandidateData(c, updateCandidate, candidateId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, sqlResult)
	}
}

func (handler *interviewHandler) getCandidateInterviewDetailsAtLevelOne(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := handler.midware.GetSelectedAndRejectedCandidatesAtLevelOne(c)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range *DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func (handler *interviewHandler) getCandidateInterviewDetailsAtLevelTwo(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := handler.midware.GetSelectedAndRejectedCandidatesAtLevelTwo(c)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range *DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func (handler *interviewHandler) getCandidateInterviewDetailsAtLevelThree(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := handler.midware.GetSelectedAndRejectedCandidatesAtLevelThree(c)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range *DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func (handler *interviewHandler) getSelectedCandidateInterviewDetailsAtLevelOne(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := handler.midware.GetSelectedCandidatesAtLevelOne(c)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range *DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func (handler *interviewHandler) getRejectedCandidateInterviewDetailsAtLevelOne(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := handler.midware.GetRejectedCandidatesAtLevelOne(c)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range *DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func (handler *interviewHandler) getSelectedCandidateInterviewDetailsAtLevelTwo(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := handler.midware.GetSelectedCandidatesAtLevelTwo(c)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range *DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func (handler *interviewHandler) getRejectedCandidateInterviewDetailsAtLevelTwo(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := handler.midware.GetRejectedCandidatesAtLevelTwo(c)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range *DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func (handler *interviewHandler) getSelectedCandidateInterviewDetailsAtLevelThree(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := handler.midware.GetSelectedCandidatesAtDMLevel(c)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range *DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func (handler *interviewHandler) getRejectedCandidateInterviewDetailsAtLevelThree(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := handler.midware.GetRejectedCandidatesAtDMLevel(c)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range *DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func (handler *interviewHandler) getOnboardedCandidateInterviewDetails(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := handler.midware.GetOnboardedCandidates(c)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range *DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func (handler *interviewHandler) getCandidatesWithAcceptedOffers(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := handler.midware.GetCandidatesOfferedAndAccepted(c)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}
	for _, candidate := range *DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func (handler *interviewHandler) getCandidatesWithAwaitedOffers(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := handler.midware.GetCandidatesOfferedAndAwaited(c)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}
	for _, candidate := range *DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func (handler *interviewHandler) createNewInterviewCandidate(c *gin.Context, candidate dtomodels.Candidate) (sql.Result, error) {
	response, err := handler.midware.CreateNewInterviewCandidate(c, candidate)
	return *response, err
}

func (handler *interviewHandler) updateInterviewCandidateData(c *gin.Context, updatecandidate dtomodels.UpdateCandidate, candidateId string) (sql.Result, error) {
	response, err := handler.midware.UpdateInterviewCandidate(c, updatecandidate, candidateId)
	return *response, err
}
