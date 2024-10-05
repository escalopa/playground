package storage

import (
	"database/sql"

	"github.com/escalopa/SQLite/model"
)

type UserRepository struct {
	DB *sql.DB
}

func (repo *UserRepository) GetAllUsers() ([]model.User, error) {
	rows, err := repo.DB.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepository) CreateUser(name string) (model.User, error) {
	result, err := repo.DB.Exec("INSERT INTO users (name) VALUES (?)", name)
	if err != nil {
		return model.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.User{}, err
	}

	return model.User{ID: int(id), Name: name}, nil
}
