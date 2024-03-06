package repository

type OrderDirection string

type Pagination struct {
	Limit  uint64
	Offset uint64
}

type OrderBy struct {
	Field     string
	Direction OrderDirection
}
