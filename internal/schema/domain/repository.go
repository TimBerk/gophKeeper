package domain

// Repository интерфейс для работы с версией
type Repository interface {
	Current() (int64, error)
}
