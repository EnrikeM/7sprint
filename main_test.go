package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

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

	assert.Equal(t, responseRecorder.Code, http.StatusOK)
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

	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
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

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Equal(t, len(list), totalCount)
}
