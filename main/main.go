package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/ArpitChinmay/interview/handlers"
	dtomodels "github.com/ArpitChinmay/interview/main/dtoModels"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var interviewHandler *handlers.InterviewHandler

func main() {
	router := gin.Default()
	router.GET("/db", GetSelectedAndRejectedCandidates)
	router.GET("/db/:level/", GetSepecificCandidateDetails)
	router.Run(":5000")

}

func init() {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/int_db_data?parseTime=true")

	if err != nil {
		log.Println("could not connect to the database...")
		log.Fatal(err)
	}

	log.Println("Connected to database...")
	DB = db
	interviewHandler = new(handlers.InterviewHandler)
	log.Println("created interviewHandler object", interviewHandler)
}

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
		c.JSON(http.StatusNotImplemented, gin.H{"error:": "feature not implemented by me..."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong value for level param..."})
	}
}

func GetSepecificCandidateDetails(c *gin.Context) {
	level, err := strconv.ParseInt(c.Param("level"), 0, 32)
	selected, err2 := strconv.ParseBool(c.Query("selected"))
	count, err3 := strconv.ParseBool(c.Query("count"))
	if err != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "problem reading url params..."})
	}

	if level == 1 {
		if selected {
			if count {
				_, datacount, err := getSelectedCandidateInterviewDetailsAtLevelOne(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, datacount)
			}
			detailsOfCandidatesDTO, _, err := getSelectedCandidateInterviewDetailsAtLevelOne(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
			}
			c.JSON(http.StatusOK, detailsOfCandidatesDTO)
		} else {
			_, datacount, err := getRejectedCandidateInterviewDetailsAtLevelOne(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
			}
			c.JSON(http.StatusOK, datacount)
			detailsOfCandidatesDTO, _, err := getRejectedCandidateInterviewDetailsAtLevelOne(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
			}
			c.JSON(http.StatusOK, detailsOfCandidatesDTO)
		}
	} else if level == 2 {
		if selected {
			if count {
				_, datacount, err := getSelectedCandidateInterviewDetailsAtLevelTwo(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
				}
				c.JSON(http.StatusOK, datacount)
			}
			detailsOfCandidatesDTO, _, err := getSelectedCandidateInterviewDetailsAtLevelTwo(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
			}
			c.JSON(http.StatusOK, detailsOfCandidatesDTO)
		} else {
			detailsOfCandidatesDTO, _, err := getRejectedCandidateInterviewDetailsAtLevelTwo(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
			}
			c.JSON(http.StatusOK, detailsOfCandidatesDTO)
		}
	} else if level == 3 {
		if selected {
			detailsOfCandidatesDTO, _, err := getSelectedCandidateInterviewDetailsAtLevelThree(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
			}
			c.JSON(http.StatusOK, detailsOfCandidatesDTO)
		} else {
			detailsOfCandidatesDTO, _, err := getRejectedCandidateInterviewDetailsAtLevelThree(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
			}
			c.JSON(http.StatusOK, detailsOfCandidatesDTO)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong request params..."})
	}
}

func getCandidateInterviewDetailsAtLevelOne(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetSelectedAndRejectedCandidatesAtLevelOne(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getCandidateInterviewDetailsAtLevelTwo(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetSelectedAndRejectedCandidatesAtLevelTwo(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getSelectedCandidateInterviewDetailsAtLevelOne(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetSelectedCandidatesAtLevelOne(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getRejectedCandidateInterviewDetailsAtLevelOne(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetRejectedCandidatesAtLevelOne(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getSelectedCandidateInterviewDetailsAtLevelTwo(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetSelectedCandidatesAtLevelTwo(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getRejectedCandidateInterviewDetailsAtLevelTwo(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetRejectedCandidatesAtLevelTwo(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getSelectedCandidateInterviewDetailsAtLevelThree(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetSelectedCandidatesAtDMLevel(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}

func getRejectedCandidateInterviewDetailsAtLevelThree(c *gin.Context) ([]dtomodels.InterviewDTO, int, error) {
	detailsOfCandidatesDTO := []dtomodels.InterviewDTO{}
	DetailsOfAllCandidates, count, err := interviewHandler.GetRejectedCandidatesAtDMLevel(c, DB)
	if err != nil {
		return detailsOfCandidatesDTO, 0, err
	}

	for _, candidate := range DetailsOfAllCandidates {
		candidateDTO := dtomodels.InterviewDTO{}
		result := candidateDTO.MapInterviewDetails(&candidate)
		detailsOfCandidatesDTO = append(detailsOfCandidatesDTO, result)
	}
	return detailsOfCandidatesDTO, count, nil
}
