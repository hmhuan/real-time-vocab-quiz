package services

import "github.com/real-time-vocab-quiz/server/dataservices"

type Services struct {
	QuizService *QuizService
}

func NewServices(holder *dataservices.Holder) *Services {
	return &Services{
		QuizService: newQuizService(holder),
	}
}
