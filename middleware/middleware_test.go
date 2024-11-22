package middlewares

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogging(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	
	testHandler := Logger(logger, handler)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	testHandler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// assert.Greater(t, time.Since(req.Context().Value("startTime").(time.Time)), 0)
}