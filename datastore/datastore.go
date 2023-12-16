package datastore

import (
	"database/sql"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gofr/model"
	"time"
)

type entries struct {
	BookId     int       `json:"bookId"`
	BookName   string    `json:"book_name"`
	AuthorName int       `json:"author_name"`
	IssuedTo   string    `json:"issued_to"`
	IssuedDate time.Time `json:"issued_date"`
}

func New() *entries {
	return &entries{}
}
func (s *entries) GetByID(ctx *gofr.Context, id string) (*model.Entries, error) {
	var resp model.Entries

	err := ctx.DB().QueryRowContext(ctx, " SELECT bookId,book_name,author_name,issued_to,issued_date FROM entries where bookid= ?", id).
		Scan(&resp.BookId, &resp.BookName, &resp.AuthorName, &resp.IssuedTo, &resp.IssuedDate)
	switch err {
	case sql.ErrNoRows:
		return &model.Entries{}, errors.EntityNotFound{Entity: "entries", ID: id}
	case nil:
		return &resp, nil
	default:
		return &model.Entries{}, err
	}
}
func (s *entries) Create(ctx *gofr.Context, entries *model.Entries) (*model.Entries, error) {
	var resp model.Entries

	// Use entries.IssuedDate if provided, otherwise use time.Now()
	issuedDate := entries.IssuedDate
	if issuedDate.IsZero() {
		issuedDate = time.Now()
	}
	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO entries (bookId, book_name, author_name, issued_to, issued_date) VALUES (?,?,?,?,?)",
		entries.BookId, entries.BookName, entries.AuthorName, entries.IssuedTo, issuedDate)
	if err != nil {
		return &model.Entries{}, errors.DB{Err: err}
	}
	// Remove the line since lastInsertID is not being used
	// lastInsertID, err := result.LastInsertId()
	// if err != nil {
	// 	return &model.Entries{}, errors.DB{Err: err}
	// }
	// Set the ID in the response using entries.BookId
	resp.BookId = entries.BookId
	resp.BookName = entries.BookName
	resp.AuthorName = entries.AuthorName
	resp.IssuedTo = entries.IssuedTo
	resp.IssuedDate = issuedDate
	return &resp, nil
}
func (s *entries) Update(ctx *gofr.Context, entries *model.Entries) (*model.Entries, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE entries SET  book_name= ? ,author_name= ? , issued_to = ?    WHERE bookid= ?",
		entries.BookName, entries.AuthorName, entries.IssuedTo, entries.BookId)
	if err != nil {
		return &model.Entries{}, errors.DB{Err: err}
	}
	return entries, nil
}
func (s *entries) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM entries where bookId= ?", id)
	if err != nil {
		return errors.DB{Err: err}
	}
	return nil
}
