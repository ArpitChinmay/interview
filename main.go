package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	dtomodels "github.com/ArpitChinmay/interview/setup/dtoModels"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var serve *ServerObject

func main() {
	// The SetupServer() method opens a db connection
	// and registers the routes.
	var ginEngine *gin.Engine
	serve, ginEngine = SetupServer()
	ginEngine.Run(":5000")
}

// Arpit Chinmay & Shaik Saisameer
func GetSelectedAndRejectedCandidates(c *gin.Context) {
	level, err := strconv.ParseInt(c.Query("level"), 0, 32)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading url params..."})
	}

	if level == 1 {
		// Excel file defines this method as L1
		detailsOfCandidatesDTO, _, err := getCandidateInterviewDetailsAtLevelOne(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		c.JSON(http.StatusOK, detailsOfCandidatesDTO)
	} else if level == 2 {
		// Excel file defines this method as L2
		detailsOfCandidatesDTO, _, err := getCandidateInterviewDetailsAtLevelTwo(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		c.JSON(http.StatusOK, detailsOfCandidatesDTO)
	} else if level == 3 {
		detailsOfCandidatesDTO, _, err := getCandidateInterviewDetailsAtLevelThree(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		c.JSON(http.StatusOK, detailsOfCandidatesDTO)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong value for level param..."})
	}
}

// Arpit Chinmay & Shaik Saisameer
func GetSepecificCandidateDetails(c *gin.Context) {
	level, err := strconv.ParseInt(c.Param("level"), 0, 32)
	selected, err2 := strconv.ParseBool(c.Query("selected"))
	count, err3 := strconv.ParseBool(c.Query("count"))
	if err != nil || err2 != nil || err3 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "problem reading url params..."})
	}

	if level == 1 {
		if selected {
			if count {
				_, datacount, err := getSelectedCandidateInterviewDetailsAtLevelOne(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, datacount)
			} else {
				detailsOfCandidatesDTO, _, err := getSelectedCandidateInterviewDetailsAtLevelOne(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			}

		} else {
			if count {
				_, datacount, err := getRejectedCandidateInterviewDetailsAtLevelOne(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, datacount)
			} else {
				detailsOfCandidatesDTO, _, err := getRejectedCandidateInterviewDetailsAtLevelOne(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			}
		}
	} else if level == 2 {
		if selected {
			if count {
				_, datacount, err := getSelectedCandidateInterviewDetailsAtLevelTwo(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, datacount)
			} else {
				detailsOfCandidatesDTO, _, err := getSelectedCandidateInterviewDetailsAtLevelTwo(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			}
		} else {
			if count {
				_, datacount, err := getRejectedCandidateInterviewDetailsAtLevelTwo(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, datacount)
			} else {
				detailsOfCandidatesDTO, _, err := getRejectedCandidateInterviewDetailsAtLevelTwo(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			}
		}
	} else if level == 3 {
		if selected {
			if count {
				_, datacount, err := getSelectedCandidateInterviewDetailsAtLevelThree(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, datacount)
			} else {
				detailsOfCandidatesDTO, _, err := getSelectedCandidateInterviewDetailsAtLevelThree(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, detailsOfCandidatesDTO)
			}
		} else {
			if count {
				_, datacount, err := getRejectedCandidateInterviewDetailsAtLevelThree(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, datacount)
			} else {
				detailsOfCandidatesDTO, _, err := getRejectedCandidateInterviewDetailsAtLevelThree(c)
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
func GetOnboardedCandidateDetails(c *gin.Context) {
	//onboarded, err1 := strconv.ParseBool(c.Query("onboarded"))
	count, err2 := strconv.ParseBool(c.Query("count"))
	if err2 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "problem reading url params..."})
	}
	if count {
		_, datacount, err := getOnboardedCandidateInterviewDetails(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		c.JSON(http.StatusOK, datacount)
	} else {
		detailsOfCandidatesDTO, _, err := getOnboardedCandidateInterviewDetails(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
		}
		c.JSON(http.StatusOK, detailsOfCandidatesDTO)
	}
}

// Kundan Kumar
func GetCandidatesWithAcceptedOffers(c *gin.Context) {
	detailsOfCandidatesDTO, _, err := getCandidatesWithAcceptedOffers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
	}
	c.JSON(http.StatusOK, detailsOfCandidatesDTO)
}

func GetCandidatesWithAwaitedOffers(c *gin.Context) {
	detailsOfCandidatesDTO, _, err := getCandidatesWithAwaitedOffers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
	}
	c.JSON(http.StatusOK, detailsOfCandidatesDTO)
}

func GetAcceptedCandidatesCount(c *gin.Context) {
	_, count, err := getCandidatesWithAcceptedOffers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
	}
	c.JSON(http.StatusOK, count)
}

func GetAwaitedCandidatesCount(c *gin.Context) {
	_, count, err := getCandidatesWithAwaitedOffers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
	}
	c.JSON(http.StatusOK, count)
}

// SindhuShree KN
func AddCandidate(c *gin.Context) {
	log.Println("Are we even getting here?")
	var candidate dtomodels.Candidate
	if err := c.ShouldBindJSON(&candidate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	sqlResult, err := createNewInterviewCandidate(c, candidate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, sqlResult)
	}
}

func UpdateCandidate(c *gin.Context) {
	candidateId := c.Param("id")
	var updateCandidate dtomodels.UpdateCandidate
	if err := c.ShouldBindJSON(&updateCandidate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	sqlResult, err := updateInterviewCandidateData(c, updateCandidate, candidateId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, sqlResult)
	}
}

func getCandidateInterviewDetailsAtLevelOne(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := serve.handler.GetSelectedAndRejectedCandidatesAtLevelOne(c)
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

func getCandidateInterviewDetailsAtLevelTwo(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := serve.handler.GetSelectedAndRejectedCandidatesAtLevelTwo(c)
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

func getCandidateInterviewDetailsAtLevelThree(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := serve.handler.GetSelectedAndRejectedCandidatesAtLevelThree(c)
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

func getSelectedCandidateInterviewDetailsAtLevelOne(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := serve.handler.GetSelectedCandidatesAtLevelOne(c)
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

func getRejectedCandidateInterviewDetailsAtLevelOne(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := serve.handler.GetRejectedCandidatesAtLevelOne(c)
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

func getSelectedCandidateInterviewDetailsAtLevelTwo(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := serve.handler.GetSelectedCandidatesAtLevelTwo(c)
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

func getRejectedCandidateInterviewDetailsAtLevelTwo(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := serve.handler.GetRejectedCandidatesAtLevelTwo(c)
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

func getSelectedCandidateInterviewDetailsAtLevelThree(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := serve.handler.GetSelectedCandidatesAtDMLevel(c)
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

func getRejectedCandidateInterviewDetailsAtLevelThree(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := serve.handler.GetRejectedCandidatesAtDMLevel(c)
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

func getOnboardedCandidateInterviewDetails(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := serve.handler.GetOnboardedCandidates(c)
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

func getCandidatesWithAcceptedOffers(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := serve.handler.GetCandidatesOfferedAndAccepted(c)
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

func getCandidatesWithAwaitedOffers(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := serve.handler.GetCandidatesOfferedAndAwaited(c)
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

func createNewInterviewCandidate(c *gin.Context, candidate dtomodels.Candidate) (sql.Result, error) {
	response, err := serve.handler.CreateNewInterviewCandidate(c, candidate)
	return *response, err
}

func updateInterviewCandidateData(c *gin.Context, updatecandidate dtomodels.UpdateCandidate, candidateId string) (sql.Result, error) {
	response, err := serve.handler.UpdateInterviewCandidate(c, updatecandidate, candidateId)
	return *response, err
}
