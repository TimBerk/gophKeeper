package domain

// Repository интерфейс для работы с записями
type Repository interface {
	Save(*Secret) error
	List(string) ([]*Secret, error)
}
