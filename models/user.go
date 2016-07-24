package models

// User user model
type User struct {
	ID       int      `json:"id" db:"id"`
	Name     string   `json:"name" db:"name"`
	Password string   `json:"pasword" db:"password"`
	Address  string   `json:"address" db:"address"`
	RoleID   int      `db:"role_id"`
	UserRole UserRole `json:"userRole"`
}
