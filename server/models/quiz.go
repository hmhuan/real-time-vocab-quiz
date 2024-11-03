package models

import (
	"time"

	"github.com/lib/pq"
)

type Quiz struct {
	ID          int64         `db:"id" goqu:"skipupdate"`
	Code        string        `db:"code"`
	QuestionIDs pq.Int64Array `db:"question_ids"`
	CreatedAt   time.Time     `db:"created_at"`
	UpdatedAt   time.Time     `db:"updated_at"`
}

func (Quiz) TableName() string {
	return "quizzes"
}

func (Quiz) IDColumnName() string {
	return "id"
}
