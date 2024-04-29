package postgres

import (
	"database/sql"
	"fmt"

	"github.com/ramzyrsr/domain/user"
)

type UserRepositoryPG struct {
	db *sql.DB
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewUserRepositoryPG(db *sql.DB) *UserRepositoryPG {
	return &UserRepositoryPG{db: db}
}

func (ur *UserRepositoryPG) Create(user *user.User) error {
	var count int
	_ = ur.db.QueryRow("SELECT count(email) FROM users WHERE email = $1", user.Email).Scan(&count)
	if count != 0 {
		errorMessage := fmt.Errorf("email already used. Please check your email")
		return errorMessage
	}

	_, err := ur.db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

func (ur *UserRepositoryPG) GetByID(id int) (*user.User, error) {
	user := &user.User{}
	err := ur.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepositoryPG) Update(user *user.User) error {
	_, err := ur.db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, user.ID)
	return err
}

func (ur *UserRepositoryPG) Delete(id int) error {
	_, err := ur.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (ur *UserRepositoryPG) GetByEmail(user *user.User) (*user.User, error) {
	err := ur.db.QueryRow("SELECT id, name, email FROM users WHERE email = $1", user.Email).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
