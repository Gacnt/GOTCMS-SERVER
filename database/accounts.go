package otdb

import (
	"database/sql"
)

// Account is the database struct to hold the account information
type Account struct {
	ID        uint64         `db:"id"`
	Name      string         `db:"name"`
	Password  string         `db:"password"`
	Secret    sql.NullString `db:"secret"`
	Type      int            `db:"type"`
	PremDays  int            `db:"premdays"`
	LastDay   int            `db:"lastday"`
	Email     string         `db:"email"`
	Creation  string         `db:"creation"`
	RoleLevel int            `db:"role_level"`
}
