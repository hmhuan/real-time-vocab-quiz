package models

import "time"

type QuizUserSession struct {
	ID        int64     `db:"id" goqu:"skipupdate"`
	QuizID    int64     `db:"quiz_id"`
	UserID    int64     `db:"user_id"`
	Score     int       `db:"score"`
	JoinTime  time.Time `db:"join_time"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (QuizUserSession) TableName() string {
	return "quiz_user_sessions"
}

func (QuizUserSession) IDColumnName() string {
	return "id"
}
