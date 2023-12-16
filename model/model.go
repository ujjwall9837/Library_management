package model

import "time"

type Entries struct {
	BookId     int       `json:"bookId"`
	BookName   string    `json:"book_name"`
	AuthorName string    `json:"author_name"`
	IssuedTo   string    `json:"issued_to"`
	IssuedDate time.Time `json:"issued_date"`
}
