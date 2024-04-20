package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//  Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое:

func TestMainHandlerWhenOk(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/cafe?city=moscow&count=1",
		nil,
	)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body)

}

// Город, который передаётся в параметре city, не поддерживается.
// Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.

func TestMainHandlerWhenCityWrong(t *testing.T) {
	var wrongCity = "wrong city value"

	req := httptest.NewRequest(
		http.MethodGet,
		"/cafe?city=tula&count=1",
		nil,
	)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, responseRecorder.Body.String(), wrongCity)
}

// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	var totalCount = 4

	req := httptest.NewRequest(
		http.MethodGet,
		"/cafe?city=moscow&count=50",
		nil,
	)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Len(t, list, totalCount)
}
