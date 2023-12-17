package main

import (
	"github.com/Pakhi-v/Hostel-Management-System/datastore"
	"github.com/Pakhi-v/Hostel-Management-System/handler"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	s := datastore.New()
	h := handler.New(s)

	app.GET("/student/{id}", h.GetByID)
	app.POST("/student", h.Create)
	app.PUT("/student/{id}", h.Update)
	app.DELETE("/student/{id}", h.Delete)

	// starting the server on a custom port
	app.Server.HTTP.Port = 8000
	app.Start()
}
