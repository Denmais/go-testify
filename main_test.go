package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body)

}

func TestMainHandlerWhenMissingCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=wrongCity", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expected := `wrong city value`
	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, responseRecorder.Body.String(), expected)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	req_slice := strings.Split(responseRecorder.Body.String(), ",")

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Len(t, req_slice, totalCount)
}
