package models

type Role string

const (
	Admin  = Role("admin")
	Common = Role("user")
)

type User struct {
	ID      int    `db:"id"`
	Email   string `db:"email"`
	Pwdhash string `db:"pwdhash"`
	Role    Role   `db:"role"`
}
