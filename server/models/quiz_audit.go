package models

import "time"

type QuizAudit struct {
	ID          int64     `db:"id" goqu:"skipupdate"`
	QuizId      int64     `db:"quiz_id"`
	EventType   string    `db:"event_type"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

func (QuizAudit) TableName() string {
	return "quizz_audits"
}

func (QuizAudit) IDColumnName() string {
	return "id"
}
