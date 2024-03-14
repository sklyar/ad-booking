package test

func ToPtr[T any](v T) *T {
	return &v
}
