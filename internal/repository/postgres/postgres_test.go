package postgres

import (
	"context"
	"regexp"
	"testing"

	"github.com/pashagolub/pgxmock/v2"
)

func TestPostgres_AddLinksPair(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close(context.Background())

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"link\" (\"long\", \"short\") VALUES ('long', 'short')"))
	p := New(mock)
	p.AddLinksPair("short", "long")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostgres_GetLongURL(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close(context.Background())

	mock.ExpectQuery(regexp.QuoteMeta("SELECT \"long\" FROM \"link\" WHERE (\"short\" = 'short')"))
	p := New(mock)
	p.GetLongURL("short")
	if err != nil {
		t.Errorf("error '%s' was not expected, while inserting a row", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
