package main

import (
	"fmt"
	"github.com/edenofjinx/go-bookings/internal/config"
	"github.com/go-chi/chi"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)
	switch value := mux.(type) {
	case *chi.Mux:
		//do nothing; test passed
		default:
			t.Error(fmt.Sprintf("type is not *chi.Mux, type is %T", value))
	}
}
