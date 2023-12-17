package datastore

import (
	"database/sql"
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"Hostel-Management-System/model"
)

type hospi struct{}

func New() *host {
	return &host{}
}

func (s *host) GetByID(ctx *gofr.Context, id int) (*model.Student, error) {
	var resp model.Student
	strId := strconv.Itoa(id)
	err := ctx.DB().QueryRowContext(ctx, " SELECT StudentID, name, gender, roomNumber, course FROM Student where StudentID=?", strId).
		Scan(&resp.StudentID, &resp.Name, &resp.Gender, &resp.RoomNumber, &resp.course)
	switch err {
	case sql.ErrNoRows:
		strId := strconv.Itoa(id)
		return &model.Student{}, errors.EntityNotFound{Entity: "Students", ID: strId}
	case nil:
		return &resp, nil
	default:
		return &model.Student{}, err
	}
}

func (s *host) Create(ctx *gofr.Context, student *model.Student) (*model.Student, error) {
	var resp model.Student

	res, err := ctx.DB().ExecContext(ctx, "INSERT INTO Students (name, gender, roomNumber, course) VALUES (?,?,?,?)", student.Name, student.Gender, student.RoomNumber, student.course)

	if err != nil {
		return &model.Student{}, errors.DB{Err: err}
	}
	id, err := res.LastInsertId()
	if err != nil {
		return &model.Student{}, errors.DB{Err: err}
	}
	err = ctx.DB().QueryRowContext(ctx, "SELECT * FROM Students WHERE StudentID = ?", id).Scan(&resp.StudentID, &resp.Name, &resp.Gender, &resp.RoomNumber, &resp.course)
	if err != nil {
		return &model.Student{}, errors.DB{Err: err}
	}

	return &resp, nil
}

func (s *host) Update(ctx *gofr.Context, student *model.Student) (*model.Student, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE Students SET name=?, gender=?, roomNumber=?, course=? WHERE StudentID=?",
		student.Name, student.Gender, student.RoomNumber, student.course, student.StudentID)
	if err != nil {
		return &model.Student{}, errors.DB{Err: err}
	}

	return student, nil
}

func (s *host) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM Students where StudentID=?", id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
