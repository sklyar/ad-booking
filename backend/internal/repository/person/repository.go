package person

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/sklyar/ad-booking/backend/internal/entity"
	"github.com/sklyar/ad-booking/backend/internal/infrastructure/database"
	"github.com/sklyar/ad-booking/backend/internal/repository"
)

const tableName = "contact_person"

var (
	columns          = []string{"id", "name", "vk_id", "created_at", "updated_at"}
	orderableColumns = []string{"id", "name", "created_at"}
)

type Storage struct {
	db         database.Handler
	sqlBuilder sq.StatementBuilderType
	fetcher    *database.Fetcher[entity.ContactPerson, model]
}

func New(dbHandler database.Handler) *Storage {
	return &Storage{
		db:         dbHandler,
		sqlBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		fetcher:    database.NewFetcher[entity.ContactPerson, model](dbHandler, scan),
	}
}

func (s *Storage) Create(ctx context.Context, person *entity.ContactPerson) error {
	m := newModel(person)
	m.CreatedAt = time.Now().UTC()
	m.UpdateAt = m.CreatedAt

	query, args, err := s.sqlBuilder.
		Insert(tableName).
		Columns(columns[1:]...).
		Values(m.Name, m.VKID, m.CreatedAt, m.UpdateAt).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	if err := s.db.QueryRow(ctx, query, args...).Scan(&m.ID); err != nil {
		return fmt.Errorf("execute query: %w", err)
	}

	*person = *m.Convert()

	return nil
}

func (s *Storage) Update(ctx context.Context, person *entity.ContactPerson) error {
	m := newModel(person)

	query, args, err := s.sqlBuilder.
		Update(tableName).
		Set("name", m.Name).
		Set("vk_id", m.VKID).
		Set("updated_at", m.UpdateAt).
		Where(sq.Eq{"id": m.ID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	if _, err := s.db.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("execute query: %w", err)
	}

	return nil
}

func (s *Storage) Get(ctx context.Context, id uint64) (*entity.ContactPerson, error) {
	return s.fetcher.Row(
		ctx,
		s.selectBuilder().Where(sq.Eq{"id": id}),
	)
}

func (s *Storage) Filter(ctx context.Context, filter repository.ContactPersonFilter) ([]entity.ContactPerson, error) {
	if err := filter.Validate(orderableColumns); err != nil {
		return nil, fmt.Errorf("validate filter: %w", err)
	}

	b := s.selectBuilder()

	if filter.Name != nil {
		b = b.Where(sq.Like{"name": filter.Name})
	}
	if filter.VKID != nil {
		b = b.Where(sq.Like{"vk_id": filter.VKID})
	}

	b = b.OrderBy(filter.Sorting.Field + " " + string(filter.Sorting.Direction))
	if filter.Pagination.Limit > 0 {
		b = b.Limit(uint64(filter.Pagination.Limit))
	}
	if filter.Pagination.LastID > 0 {
		switch filter.Sorting.Direction {
		case repository.OrderDirectionAsc:
			b = b.Where(sq.Gt{"id": filter.Pagination.LastID})
		case repository.OrderDirectionDesc:
			b = b.Where(sq.Lt{"id": filter.Pagination.LastID})
		}
	}

	return s.fetcher.Rows(
		ctx,
		b,
	)
}

func (s *Storage) selectBuilder() sq.SelectBuilder {
	return s.sqlBuilder.
		Select(columns...).
		From(tableName).
		Where(sq.Eq{"deleted_at": nil})
}

func scan(row database.Row, item *model) error {
	return row.Scan(
		&item.ID,
		&item.Name,
		&item.VKID,
		&item.CreatedAt,
		&item.UpdateAt,
	)
}
