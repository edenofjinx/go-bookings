package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var testHandler myHandler
	handler := NoSurf(&testHandler)
	switch value := handler.(type) {
	case http.Handler:
	// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is: %T", value))
	}
}

func TestSessionLoad(t *testing.T) {
	var testHandler myHandler
	handler := SessionLoad(&testHandler)
	switch value := handler.(type) {
	case http.Handler:
	// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is: %T", value))
	}
}
