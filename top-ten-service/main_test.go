package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	router := setupRoutes()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var res []WordCountResult
	if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
		log.Fatalln(err)
	}
	assert.LessOrEqual(t, len(res), 10)
}
