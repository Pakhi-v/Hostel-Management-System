package datastore

import (
	"github.com/Pakhi-v/Hostel-Management-System/model"

	"gofr.dev/pkg/gofr"
)

type Student interface {
	// GetByID retrieves a student record based on its ID.
	GetByID(ctx *gofr.Context, id int) (*model.Student, error)
	// Create inserts a new student record into the database.
	Create(ctx *gofr.Context, model *model.Student) (*model.Student, error)
	// Update updates an existing student record with the provided information.
	Update(ctx *gofr.Context, model *model.Student) (*model.Student, error)
	// Delete removes a Student record from the database based on its ID.
	Delete(ctx *gofr.Context, id int) error
}