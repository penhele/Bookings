package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/penhele/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error(fmt.Sprintf("types is not *chi.Mux, type is %T", v))
	}
}