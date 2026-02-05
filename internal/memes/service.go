package memes

import (
	"database/sql"
)

type Service struct {
	repo *Repository
}

func NewService(db *sql.DB) *Service {
	return &Service{
		repo: NewRepository(db),
	}
}

func (s *Service) Create(dto CreateMemeDTO) (*Meme, error) {
    meme := &Meme{
        Title:     dto.Title,
        ImageData: dto.ImageData,
        AuthorID:  dto.AuthorID,
    }

    if err := s.repo.Create(meme); err != nil {
        return nil, err
    }
    return meme, nil
}

func (s *Service) GetAll() ([]Meme, error) {
    return s.repo.FindAll()
}