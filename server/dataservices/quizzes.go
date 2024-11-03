package dataservices

import (
	"github.com/real-time-vocab-quiz/server/models"
	"github.com/vnteamopen/dataservicex"
)

type Quizzes struct {
	dataservicex.DataServices[models.Quiz]
}
