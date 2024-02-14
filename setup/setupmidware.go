package setup

import (
	"github.com/ArpitChinmay/interview/middleware"
)

type MidwareGenerator struct {
	midware middleware.InterviewMiddleware
}

func SetupMiddleware(repo *RepositoryGenerator) *MidwareGenerator {
	interviewMidware := middleware.InitializeInterviewMiddleware(repo.InterviewRepository)
	return &MidwareGenerator{midware: interviewMidware}
}
