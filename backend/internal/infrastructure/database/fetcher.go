package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/sklyar/go-transact/txsql"
)

// model is an interface that should be implemented by all models.
type model[EntityT any] interface {
	Convert() *EntityT
}

// scannerFunc is a function that scans a row into a model.
type scannerFunc[ModelT any] func(row txsql.Row, item *ModelT) error

// Fetcher is a struct that fetches data from the database.
type Fetcher[EntityT any, ModelT model[EntityT]] struct {
	db   txsql.DB
	scan scannerFunc[ModelT]
}

// NewFetcher creates a new Fetcher.
func NewFetcher[EntityT any, ModelT model[EntityT]](db txsql.DB, scannerFn scannerFunc[ModelT]) *Fetcher[EntityT, ModelT] {
	return &Fetcher[EntityT, ModelT]{db: db, scan: scannerFn}
}

// Row fetches a single row from the database.
func (f *Fetcher[EntityT, ModelT]) Row(ctx context.Context, builder sq.SelectBuilder) (*EntityT, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	var m ModelT
	row := f.db.QueryRow(ctx, query, args...)
	if err := f.scan(row, &m); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("scan row: %w", err)
	}

	return m.Convert(), nil
}

// Rows fetches multiple rows from the database.
func (f *Fetcher[EntityT, ModelT]) Rows(ctx context.Context, builder sq.SelectBuilder) ([]EntityT, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	rows, err := f.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("execute query: %w", err)
	}
	defer rows.Close()

	var items []EntityT
	for rows.Next() {
		var m ModelT
		if err := f.scan(rows, &m); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		items = append(items, *m.Convert())
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate rows: %w", err)
	}

	return items, nil
}
