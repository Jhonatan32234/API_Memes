package users

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(user *User) error {
	
	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id, created_at
	`
	return r.db.QueryRow(
		query,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreatedAt)
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	u := &User{}
	err := r.db.QueryRow(
		"SELECT id, email, password, created_at FROM users WHERE email=$1",
		email,
	).Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt)

	if err != nil {
		return nil, err
	}
	return u, nil
}


func (r *Repository) FindAll() ([]User, error) {
	rows, err := r.db.Query(`
		SELECT id, email, created_at
		FROM users
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *Repository) FindByID(id int64) (*User, error) {
	u := &User{}
	err := r.db.QueryRow(`
		SELECT id, email, created_at
		FROM users
		WHERE id = $1
	`, id).Scan(&u.ID, &u.Email, &u.CreatedAt)

	if err != nil {
		return nil, err
	}

	return u, nil
}
