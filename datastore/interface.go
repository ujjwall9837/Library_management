package datastore

import (
	"gofr.dev/pkg/gofr"
	"gofr/model"
)

type Entries interface {
	// GetByID retrieves a book record based on its ID.
	GetByID(ctx *gofr.Context, id string) (*model.Entries, error)
	// Create inserts a new book record into the database.
	Create(ctx *gofr.Context, model *model.Entries) (*model.Entries, error)
	// Update updates an existing book with the provided information.
	Update(ctx *gofr.Context, model *model.Entries) (*model.Entries, error)
	// Delete removes a entry record from the database based on its ID.
	Delete(ctx *gofr.Context, id int) error
}
