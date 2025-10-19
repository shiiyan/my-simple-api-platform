package domain

type UserRepository interface {
	GetByID(id string) (*User, error)
}
