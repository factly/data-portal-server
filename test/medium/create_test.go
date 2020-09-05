package medium

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/data-portal-server/action"
	"github.com/factly/data-portal-server/test"
	"github.com/gavv/httpexpect"
	"gopkg.in/h2non/gock.v1"
)

func TestCreateMedium(t *testing.T) {
	// Setup DB
	mock := test.SetupMockDB()

	// Setup HttpExpect
	router := action.RegisterRoutes()
	server := httptest.NewServer(router)
	defer server.Close()

	test.MeiliGock()
	gock.New(server.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()

	e := httpexpect.New(t, server.URL)

	t.Run("create medium", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "dp_medium"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, Medium["name"], Medium["slug"], Medium["type"], Medium["title"], Medium["description"], Medium["caption"], Medium["alt_text"], Medium["file_size"], Medium["url"], Medium["dimensions"]).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
		mock.ExpectCommit()

		e.POST(basePath).
			WithJSON(Medium).
			Expect().
			Status(http.StatusCreated).
			JSON().
			Object().
			ContainsMap(Medium)

		test.ExpectationsMet(t, mock)
	})

	t.Run("unprocessable medium body", func(t *testing.T) {
		e.POST(basePath).
			WithJSON(invalidMedium).
			Expect().
			Status(http.StatusUnprocessableEntity)
	})

	t.Run("empty medium body", func(t *testing.T) {
		e.POST(basePath).
			Expect().
			Status(http.StatusUnprocessableEntity)
	})

	t.Run("creating medium fails", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "dp_medium"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, Medium["name"], Medium["slug"], Medium["type"], Medium["title"], Medium["description"], Medium["caption"], Medium["alt_text"], Medium["file_size"], Medium["url"], Medium["dimensions"]).
			WillReturnError(errors.New("cannot create medium"))
		mock.ExpectRollback()

		e.POST(basePath).
			WithJSON(Medium).
			Expect().
			Status(http.StatusInternalServerError)

		test.ExpectationsMet(t, mock)
	})

	t.Run("create medium when meili is down", func(t *testing.T) {
		gock.Off()
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "dp_medium"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, Medium["name"], Medium["slug"], Medium["type"], Medium["title"], Medium["description"], Medium["caption"], Medium["alt_text"], Medium["file_size"], Medium["url"], Medium["dimensions"]).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
		mock.ExpectRollback()

		e.POST(basePath).
			WithJSON(Medium).
			Expect().
			Status(http.StatusInternalServerError)
		test.ExpectationsMet(t, mock)
	})

}
