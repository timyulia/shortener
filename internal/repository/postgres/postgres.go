package postgres

import (
	"context"
	"errors"

	"github.com/doug-martin/goqu/v9"
	pgx "github.com/jackc/pgx/v5"
	pgconn "github.com/jackc/pgx/v5/pgconn"
)

const (
	dbLink   = "link"
	colShort = "short"
	colLong  = "long"
)

type PgxIface interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
}

// Убрать принты
// вынести константы
// разнести внутренний, стандартные и внешние пакеты

type Postgres struct {
	//conn *pgx.Conn
	conn PgxIface
}

type linkPair struct {
	Short string `db:"short"`
	Long  string `db:"long"`
}

// New ...
func New(conn PgxIface) *Postgres {
	return &Postgres{conn: conn}
}

func (r *Postgres) GetLongURL(short string) (string, error) {
	selectSQL, _, _ := goqu.From(dbLink).
		Select(colLong).
		Where(
			goqu.Ex{colShort: short}).ToSQL()
	var long string

	err := r.conn.QueryRow(context.Background(), selectSQL).Scan(&long)

	if errors.Is(err, pgx.ErrNoRows) {
		return "", nil
	}

	return long, err
}

func (r *Postgres) AddLinksPair(short, long string) error {
	insertSQL, _, _ := goqu.Insert(dbLink).Rows(
		linkPair{short, long},
	).ToSQL()

	_, err := r.conn.Exec(context.Background(), insertSQL)

	return err
}
