package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestSongs(t *testing.T) {

	assert := assert.New(t)

	tests := []struct {
		description  string
		url          string
		expectedCode int
	}{
		{
			description:  "Songs - list catalog, pg. 1",
			url:          "/songs/0/20",
			expectedCode: http.StatusOK,
		},
		{
			description:  "Songs - list catalog, pg. 2",
			url:          "/songs/20/20",
			expectedCode: http.StatusOK,
		},
		{
			description:  "Songs - list catalog, no pg specified",
			url:          "/songs/",
			expectedCode: http.StatusNotFound,
		},
	}

	// mockDB, mock, err := sqlmock.New()
	// defer mockDB.Close()
	// sqlxDB = sqlx.NewDb(mockDB,"sqlmock")

	router := httprouter.New()

	for _, test := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.url, nil)
		router.ServeHTTP(w, req)
		assert.Equal(test.expectedCode, w.Code, test.description)
	}
}
