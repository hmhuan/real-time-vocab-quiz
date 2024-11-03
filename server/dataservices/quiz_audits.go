package dataservices

import (
	"github.com/real-time-vocab-quiz/server/models"
	"github.com/vnteamopen/dataservicex"
)

type QuizAudits struct {
	dataservicex.DataServices[models.QuizAudit]
}
