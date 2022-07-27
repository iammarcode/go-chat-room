package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/iammarcode/go-chat-room/initializer"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {

	router := initializer.R

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v1/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "miles")
}
