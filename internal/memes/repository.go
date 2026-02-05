package memes

import "database/sql"

type Repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) Create(m *Meme) error {
    query := `INSERT INTO memes (title, image_data, author_id) 
              VALUES ($1, $2, $3) RETURNING id, created_at`
    return r.db.QueryRow(query, m.Title, m.ImageData, m.AuthorID).Scan(&m.ID, &m.CreatedAt)
}

func (r *Repository) FindAll() ([]Meme, error) {
    query := `
        SELECT m.id, m.title, m.image_data, m.author_id, u.email, m.created_at 
        FROM memes m 
        JOIN users u ON m.author_id = u.id 
        ORDER BY m.created_at DESC`
    
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []Meme
    for rows.Next() {
        var m Meme
        if err := rows.Scan(&m.ID, &m.Title, &m.ImageData, &m.AuthorID, &m.AuthorName, &m.CreatedAt); err != nil {
            return nil, err
        }
        list = append(list, m)
    }
    return list, nil
}