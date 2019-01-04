package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/mjetpax/80sMixtapeAPI/config"
	"github.com/stretchr/testify/assert"
)

func TestCassette(t *testing.T) {

	assert := assert.New(t)

	// Define tests
	tests := []struct {
		description  string
		url          string
		expectedCode int
	}{
		{
			description:  "Casssette - default",
			url:          "/cassette/C60",
			expectedCode: http.StatusOK,
		},
	}

	// Mock the db
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Expected to open mock database connection, received error, %s", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// Create mock db data
	rows := sqlmock.NewRows([]string{"id", "title", "artist", "year", "video", "duration_label", "duration"}).
		AddRow(1, "Call Me", "Blondie", 1980, "https://music.youtube.com/watch?v=StKVS0eI85I", "2:15", 135)

	// Define mock db query / response
	mock.ExpectQuery("SELECT (.+) FROM songs WHERE id > (.+) ORDER BY random() Limit (.+)").
		WithArgs(25).
		WillReturnRows(rows)

	// Define services to pass mock db to handler.
	var services config.Services
	services = config.Services{DB: sqlxDB}

	// Set up router
	router := httprouter.New()
	router.Handle(http.MethodGet, "/cassette/*cassette_type", GetCassette(&services))

	// Loop through tests.
	for _, test := range tests {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, test.url, nil)
		router.ServeHTTP(w, req)
		assert.Equal(test.expectedCode, w.Code, test.description)
	}
}
