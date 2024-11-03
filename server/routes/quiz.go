package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/real-time-vocab-quiz/server/controllers"
)

func newQuizRouter(g *gin.RouterGroup, controller *controllers.QuizController) {
	group := g.Group("/quizzes")
	group.GET("/:code", controller.JoinQuizByCode)

	// group.GET("/:code/reserve", reserveHandler)
}
