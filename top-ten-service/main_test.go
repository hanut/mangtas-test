package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	router := setupRoutes()
	testString := `
		ben benny tested ben
		benny benny ben benny
		tested tested benny ben
	`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(testString))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var res []WordCountResult
	if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, 3, len(res))
	assert.Equal(t, "benny", res[0].Word)
	assert.Equal(t, uint16(5), res[0].Count)
	assert.Equal(t, "ben", res[1].Word)
	assert.Equal(t, uint16(4), res[1].Count)
	assert.Equal(t, "tested", res[2].Word)
	assert.Equal(t, uint16(3), res[2].Count)
}
