package storage

import (
	"log"
	"user/models"

	"github.com/jmoiron/sqlx"
)

type User struct {
	DB *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{
		DB: db,
	}
}

func (u *User) CreateUser(req *models.UserRequest) (*models.UserResponse, error) {
    query := `
        INSERT INTO users (name, email, age) 
        VALUES ($1, $2, $3)
        RETURNING id, name, email, age;
    `

    row := u.DB.QueryRow(query, req.Name, req.Email, req.Age)

    var res models.UserResponse

    err := row.Scan(&res.Id, &res.Name, &res.Email, &res.Age)
    if err != nil {
        log.Printf("Failed to create user: %v", err)
        return nil, err
    }

    return &res, nil
}

func (u *User) UpdateUserById(id int, req *models.UserRequest) (*models.UserResponse, error) {
	query := `
		UPDATE users
		SET name = $1, email = $2, age = $3
		WHERE id = $4
		RETURNING id, name, email, age
	`

	row := u.DB.QueryRow(query, req.Name, req.Email, req.Age, id)

	var res models.UserResponse

	err := row.Scan(&res.Id, &res.Name, &res.Email, &res.Age)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (u *User) DeleteUserByID(id int) (*models.UserResponse, error) {
	query := `
		DELETE FROM users
		WHERE id = $1
		RETURNING id, name, email, age
	`

	row := u.DB.QueryRow(query, id)

	var res models.UserResponse

	err := row.Scan(&res.Id, &res.Name, &res.Email, &res.Age)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (u *User) GetUserByID(id int) (*models.UserResponse, error){
	query := `
		SELECT id, name, email, age 
		FROM users
		WHERE id = $1
	`

	row := u.DB.QueryRow(query, id)

	var res models.UserResponse

	err := row.Scan(&res.Id, &res.Name, &res.Email, &res.Age)
	if err != nil {
		return nil, err
	}

	return &res, err
}
