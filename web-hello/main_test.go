package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSayHelloName(t *testing.T) {

	// make a request to send to our handler
	r, err := http.NewRequest("GET", "/?name=poop", nil)
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
	// make a test http recorder
	w := httptest.NewRecorder()

	sayhelloName(w, r)

	expectedResult := "Hello poop!"
	actualResult := w.Body.String()

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
