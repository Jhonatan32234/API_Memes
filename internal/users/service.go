package users

import (
	"database/sql"
	"fmt"

	"api_memes/internal/shared"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(db *sql.DB) *Service {
	return &Service{
		repo: NewRepository(db),
	}
}

func (s *Service) Create(dto CreateUserDTO) (*User, error) {
	if err := ValidateCreateUser(dto); err != nil {
		return nil, shared.ErrValidation
	}

	_, err := s.repo.FindByEmail(dto.Email)
	if err == nil {
		return nil, shared.ErrDuplicate
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)

	user := &User{
		Email:    dto.Email,
		Password: string(hashed),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetAll() ([]User, error) {
	return s.repo.FindAll()
}


func (s *Service) Login(email, password string) (*User, error) {
    user, err := s.repo.FindByEmail(email)
    if err != nil {
        return nil, shared.ErrNotFound
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, fmt.Errorf("invalid credentials")
    }

    return user, nil
}

func (s *Service) GetByID(id int64) (*User, error) {
	return s.repo.FindByID(id)
}

func toUserResponse(u *User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}
