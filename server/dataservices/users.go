package dataservices

import (
	"github.com/real-time-vocab-quiz/server/models"
	"github.com/vnteamopen/dataservicex"
)

type Users struct {
	dataservicex.DataServices[models.User]
}
