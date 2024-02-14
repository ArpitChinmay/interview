package setup

import (
	"github.com/ArpitChinmay/interview/handlers"
)

type HandlerGenerator struct {
	InterviewHandler handlers.InterviewHandler
}

func SetupHandler(midware *MidwareGenerator) *HandlerGenerator {
	interviewHandler := handlers.InitializeHandler(midware.midware)
	return &HandlerGenerator{InterviewHandler: interviewHandler}
}
