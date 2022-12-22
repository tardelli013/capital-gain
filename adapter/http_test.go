package adapter

import (
	"bytes"
	"desafio-nu/core/usecases"
	"desafio-nu/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthz(t *testing.T) {
	useCase := usecases.NewOperationUseCase()
	router := SetupRouter(useCase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/healthz", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "running", w.Body.String())
}

func TestCalculateCapitalGain(t *testing.T) {
	useCase := usecases.NewOperationUseCase()
	router := SetupRouter(useCase)

	input, _ := json.Marshal(mocks.MockCase7())
	body := bytes.NewReader(input)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/calc", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `[{"tax":0},{"tax":0},{"tax":0},{"tax":0},{"tax":3000},{"tax":0},{"tax":0},{"tax":3700},{"tax":0}]`, w.Body.String())
}

func TestCalculateCapitalGainBadRequest(t *testing.T) {
	useCase := usecases.NewOperationUseCase()
	router := SetupRouter(useCase)

	input, _ := json.Marshal("invalid input")
	body := bytes.NewReader(input)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/calc", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCalculateCapitalGainInternalServerError(t *testing.T) {
	useCase := mocks.NewOperationUseCaseMock()
	router := SetupRouter(useCase)

	input, _ := json.Marshal(mocks.MockCase7())
	body := bytes.NewReader(input)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/calc", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
