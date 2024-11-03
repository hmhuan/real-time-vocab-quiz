package services

import (
	"github.com/real-time-vocab-quiz/server/dataservices"
	"github.com/real-time-vocab-quiz/server/models"
)

type QuizService struct {
	Holder *dataservices.Holder
}

var quizService *QuizService

func newQuizService(holder *dataservices.Holder) *QuizService {
	if quizService == nil {
		quizService = &QuizService{
			Holder: holder,
		}
	}
	return quizService
}

func (qs *QuizService) GetQuizByCode(code string) (models.Quiz, error) {
	// TODO: implement later
	return models.Quiz{}, nil
}
