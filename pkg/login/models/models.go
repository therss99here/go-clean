package login

import (
	"time"
)

type user struct {
	Id        int64      `db:"id" auto:"true" json:"id"`
	Email     string     `db:"email" json:"email"`
	Password  string     `db:"password" json:"password"`
	FirstName string     `db:"first_name" json:"first_name"`
	LastName  *string    `db:"last_name" json:"last_name"`
	PhoneId   *uuid.UUID `db:"phone_id" json:"phone_id"`
	UpdatedAt time.Time  `db:"updated_at json:"updated_at`
	CreatedAt time.Time  `db:"created_at json:"created_at`
}
