package postgres

import (
	"context"
	"errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5"
)

const (
	dbLink   = "link"
	colShort = "short"
	colLong  = "long"
)

// Убрать принты
// вынести константы
// разнести внутренний, стандартные и внешние пакеты

type Postgres struct {
	conn *pgx.Conn
}

type linkPair struct {
	Short string `db:"shorturl"`
	Long  string `db:"longurl"`
}

// New ...
func New(conn *pgx.Conn) *Postgres {
	return &Postgres{conn: conn}
}

func (r *Postgres) GetShortURL(long string) (string, error) {
	return "", nil
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
