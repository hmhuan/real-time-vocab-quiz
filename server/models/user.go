package models

import "time"

type User struct {
	ID int64 `db:"id" goqu:"skipupdate"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (User) IDColumnName() string {
	return "id"
}
