package handler

import (
	"strconv"

	"github.com/Pakhi-v/Hostel-Management-System/datastore"
	"github.com/Pakhi-v/Hostel-Management-System/model"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	store datastore.Student
}

func New(s datastore.Student) handler {
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
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}
	response, err := h.store.GetByID(ctx, idInt)

	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "Patients",
			ID:     id,
		}
	}
	return response, nil
}

var patient model.student

// ctx.Bind() binds the incoming data from the HTTP request to a provided interface (i).
func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	if err := ctx.Bind(&Student); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.store.Create(ctx, &Student)
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

	var student model.Student
	if err = ctx.Bind(&student); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	student.StudenttID = id

	resp, err := h.store.Update(ctx, &student)
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
