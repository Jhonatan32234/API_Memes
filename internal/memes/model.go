package memes

import "time"

type Meme struct {
    ID         int64     `json:"id"`
    Title      string    `json:"title"`
    ImageData  string    `json:"image_data"`
    AuthorID   int64     `json:"author_id"`
    AuthorName string    `json:"author_name"`
    CreatedAt  time.Time `json:"created_at"`
}

type CreateMemeDTO struct {
    Title     string `json:"title"`
    ImageData string `json:"image_data"`
    AuthorID  int64  `json:"author_id"`
}