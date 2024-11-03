package models

import (
	"time"

	"github.com/lib/pq"
)

type Question struct {
	ID                 int64          `db:"id" goqu:"skipupdate"`
	Title              string         `db:"title"`
	Content            string         `db:"content"`
	BackgroundImageUrl string         `db:"background_image_url"`
	QuestionType       string         `db:"question_type"`
	Options            pq.StringArray `db:"options"`
	CorrectAnswers     int            `db:"correct_answers"`
	Points             int            `db:"points"`
	Explanation        *string        `db:"explanation"`
	TimeLimit          int            `db:"time_limit"`
	CreatedAt          time.Time      `db:"created_at"`
	UpdatedAt          time.Time      `db:"updated_at"`
}
