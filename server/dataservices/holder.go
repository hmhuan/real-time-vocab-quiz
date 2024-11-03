package dataservices

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/real-time-vocab-quiz/server/models"
	"github.com/vnteamopen/dataservicex"
)

type Holder struct {
	Users      *Users
	Quizzes    *Quizzes
	QuizAudits *QuizAudits
}

var holder *Holder

func NewHolder(dbConn *sqlx.DB) *Holder {
	dialect := goqu.Dialect("postgres")
	users := Users{
		DataServices: dataservicex.DataServices[models.User](
			dataservicex.NewDataServices(
				dbConn,
				dataservicex.WithDialect[models.User](dialect),
			),
		),
	}

	quizzes := Quizzes{
		DataServices: dataservicex.DataServices[models.Quiz](
			dataservicex.NewDataServices(
				dbConn,
				dataservicex.WithDialect[models.Quiz](dialect),
			),
		),
	}

	quizAudits := QuizAudits{
		DataServices: dataservicex.DataServices[models.QuizAudit](
			dataservicex.NewDataServices(
				dbConn,
				dataservicex.WithDialect[models.QuizAudit](dialect),
			),
		),
	}
	holder := &Holder{
		Users:      &users,
		Quizzes:    &quizzes,
		QuizAudits: &quizAudits,
	}
	return holder
}
