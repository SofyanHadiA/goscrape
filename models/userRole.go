package models

// UserRole user role model
type UserRole struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
