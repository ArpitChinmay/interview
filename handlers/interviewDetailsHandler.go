package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"interview-dashboard/db/datareader"
	"interview-dashboard/models"

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
