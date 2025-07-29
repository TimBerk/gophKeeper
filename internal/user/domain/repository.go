package domain

// Repository - интерфейс для работы с данными пользователя
type Repository interface {
	Save(*User) error
	ByUsername(string) (*User, error)
	ByID(ID) (*User, error)
}
