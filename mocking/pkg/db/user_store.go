package db

// User represents a simple user.
type User struct {
	ID          uint
	Name, Email string
}

// UserStore is a store for getting and creating Users.
type UserStore interface {
	Get(id uint) (*User, error)
	Create(name, email string) *User
}
