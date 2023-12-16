package datastore

import (
	"database/sql"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"gofr/model"
)

type student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Class string `json:"class"`
}

func New() *student {
	return &student{}
}

func (s *student) GetByID(ctx *gofr.Context, id string) (*model.Student, error) {
	var resp model.Student

	err := ctx.DB().QueryRowContext(ctx, " SELECT id,name,age,class FROM student where id= ?", id).
		Scan(&resp.ID, &resp.Name, &resp.Age, &resp.Class)
	switch err {
	case sql.ErrNoRows:
		return &model.Student{}, errors.EntityNotFound{Entity: "student", ID: id}
	case nil:
		return &resp, nil
	default:
		return &model.Student{}, err
	}
}

//func (s *student) Create(ctx *gofr.Context, student *model.Student) (*model.Student, error) {
//	var resp model.Student
//
//	result, err := ctx.DB().ExecContext(ctx, "INSERT INTO student (id,name, age, class) VALUES ( ? , ? , ? , ? )",
//		student.ID, student.Name, student.Age, student.Class)
//
//	if err != nil {
//		return &model.Student{}, errors.DB{Err: err}
//	}
//
//	// Get the last inserted ID
//	lastInsertID, err := result.LastInsertId()
//	if err != nil {
//		return &model.Student{}, errors.DB{Err: err}
//	}
//
//	// Set the ID in the response
//	resp.ID = int(lastInsertID)
//	resp.ID = student.ID
//	resp.Name = student.Name
//	resp.Age = student.Age
//	resp.Class = student.Class
//
//	return &resp, nil
//}

func (s *student) Create(ctx *gofr.Context, student *model.Student) (*model.Student, error) {
	var resp model.Student

	result, err := ctx.DB().ExecContext(ctx, "INSERT INTO student (name, age, class) VALUES (?, ?, ?)",
		student.Name, student.Age, student.Class)

	if err != nil {
		return &model.Student{}, errors.DB{Err: err}
	}

	// Get the last inserted ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return &model.Student{}, errors.DB{Err: err}
	}

	// Set the ID in the response
	resp.ID = int(lastInsertID)
	resp.Name = student.Name
	resp.Age = student.Age
	resp.Class = student.Class

	return &resp, nil
}
func (s *student) Update(ctx *gofr.Context, student *model.Student) (*model.Student, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE student SET name= ? ,age= ? ,class= ?  WHERE id= ?",
		student.Name, student.Age, student.Class, student.ID)
	if err != nil {
		return &model.Student{}, errors.DB{Err: err}
	}

	return student, nil
}

func (s *student) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM student where id= ?", id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
