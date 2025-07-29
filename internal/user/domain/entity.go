package domain

import "golang.org/x/crypto/bcrypt"

type ID string

// User - сущность для работы с данными пользователя
type User struct {
	ID       ID
	Username string
	Hash     []byte
}

// New - инициализация нового пользователя
func New(id ID, u, p string) (*User, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{ID: id, Username: u, Hash: h}, nil
}

// Check - проверка пароля пользователя
func (u *User) Check(pw string) bool {
	return bcrypt.CompareHashAndPassword(u.Hash, []byte(pw)) == nil
}
