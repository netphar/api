package main_test

import (
	"github.com/ad05bzag/rest-api"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a main.App

const (
	user     = "postgres"
	password = "Sts56ttKaq"
	dbname   = "drugcomb"
)

func TestMain(m *testing.M) {
	a = main.App{}
	a.Initialize(user, password, dbname)

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS doses
(
id INTEGER,
druga VARCHAR(256),
drugb VARCHAR(256),
dosea DOUBLE PRECISION,
doseb DOUBLE PRECISION,
response DOUBLE PRECISION,
dss DOUBLE PRECISION,
synergy_hsa DOUBLE PRECISION,
cell_line VARCHAR(256),
CONSTRAINT doses_pkey PRIMARY KEY (id)
)`

func clearTable() {
	a.DB.Exec("DELETE FROM doses")
	a.DB.Exec("ALTER SEQUENCE doses_id_seq RESTART WITH 1")
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/doses", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
