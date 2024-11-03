package controllers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/real-time-vocab-quiz/server/services"
)

type Controllers struct {
	QuizController *QuizController
}

func NewControllers(serviceList *services.Services) *Controllers {
	return &Controllers{
		QuizController: newQuizController(serviceList.QuizService),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO: temporarily for development
		return true
	},
}
