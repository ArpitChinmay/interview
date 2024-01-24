package src

import (
	"database/sql"
	"interview-dashboard/handlers"
	dtomodels "interview-dashboard/main/dtoModels"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var interviewHandler *handlers.InterviewHandler

func main() {
	router := gin.Default()
	router.GET("/db/", L1)
	//router.GET("/db/:level/", GetSepecificCandidateDetails)
}

func init() {
	db, err := sql.Open("mysql", "root:admin@localhost:3306")

	if err != nil {
		log.Println("could not connect to the database...")
		log.Fatal(err)
	}

	log.Println("Connected to database...")
	DB = db
	interviewHandler = new(handlers.InterviewHandler)
	log.Println("created interviewHandler object", interviewHandler)
}

func L1(c *gin.Context) {
	detailsOfCandidatesDTO, _, err := getCandidateInterviewDetailsAtLevelOne(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "some error occurred..."})
	}
	c.JSON(http.StatusOK, detailsOfCandidatesDTO)
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
