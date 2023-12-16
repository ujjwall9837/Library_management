package handler

import (
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gofr/datastore"
	"gofr/model"
	"strconv"
)

type handler struct {
	store datastore.Entries
}

func New(s datastore.Entries) handler {
	return handler{store: s}
}
func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	// ctx.PathParam() returns the path parameter from HTTP request.
	id := ctx.PathParam("id")
	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	if _, err := validateID(id); err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.store.GetByID(ctx, id)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "entry",
			ID:     id,
		}
	}
	return resp, nil
}
func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var student model.Entries

	// ctx.Bind() binds the incoming data from the HTTP request to a provided interface (i).
	if err := ctx.Bind(&student); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	resp, err := h.store.Create(ctx, &student)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	id, err := validateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}
	var entries model.Entries
	if err = ctx.Bind(&entries); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	entries.BookId = id
	resp, err := h.store.Update(ctx, &entries)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	id, err := validateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}
	if err := h.store.Delete(ctx, id); err != nil {
		return nil, err
	}
	return "Deleted successfully", nil
}
func validateID(id string) (int, error) {
	res, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return res, err
}
