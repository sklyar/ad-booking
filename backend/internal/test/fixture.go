package test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

// readFixtureOptions is a struct that holds options for ReadFixture.
type readFixtureOptions struct {
	// CallerSkip is the number of stack frames to ascend when determining the
	// caller's file path. It defaults to 2, which is the number of stack frames
	// between this function and the caller.
	CallerSkip int

	// RetainTrailingNewline determines whether the trailing newline is retained
	// when reading the fixture. It defaults to false.
	RetainTrailingNewline bool

	// TemplateValues holds key-value pairs to be substituted in the fixture content.
	// It defaults to nil.
	TemplateValues map[string]any
}

// ReadFixtureOption is a function that modifies readFixtureOptions.
type ReadFixtureOption func(*readFixtureOptions)

// WithCallerSkip sets the number of stack frames to ascend when determining the
// caller's file path.
func WithCallerSkip(skip int) ReadFixtureOption {
	return func(opts *readFixtureOptions) {
		opts.CallerSkip = skip
	}
}

// WithRetainTrailingNewline determines whether the trailing newline is retained
// when reading the fixture.
func WithRetainTrailingNewline(retain bool) ReadFixtureOption {
	return func(opts *readFixtureOptions) {
		opts.RetainTrailingNewline = retain
	}
}

// WithTemplateValues configures template substitutions for fixture content processing.
func WithTemplateValues(values map[string]any) ReadFixtureOption {
	return func(opts *readFixtureOptions) {
		opts.TemplateValues = values
	}
}

// ReadFixture reads a test data file and returns its content as bytes.
// The file location is relative to the caller's file path.
func ReadFixture(t *testing.T, name string, opts ...ReadFixtureOption) []byte {
	t.Helper()

	options := readFixtureOptions{
		CallerSkip: 2,
	}
	for _, opt := range opts {
		opt(&options)
	}

	return readFixture(t, name, options)
}

// ReadTypedFixture reads a JSON test data file and unmarshals it into a specified type.
// The file location is relative to the caller's file path.
func ReadTypedFixture[T any](t *testing.T, name string, opts ...ReadFixtureOption) *T {
	t.Helper()

	options := readFixtureOptions{
		CallerSkip: 2,
	}
	for _, opt := range opts {
		opt(&options)
	}

	b := readFixture(t, name, options)

	var res T
	err := json.Unmarshal(b, &res)
	require.NoError(t, err)

	return &res
}

func ReadProtoFixture(t *testing.T, name string, out json.Unmarshaler, opts ...ReadFixtureOption) {
	t.Helper()

	options := readFixtureOptions{
		CallerSkip: 2,
	}
	for _, opt := range opts {
		opt(&options)
	}

	b := readFixture(t, name, options)

	err := out.UnmarshalJSON(b)
	require.NoError(t, err)
}

// readFixture reads and optionally processes a fixture file using text/template for substitutions.
func readFixture(t *testing.T, name string, opts readFixtureOptions) []byte {
	_, curFile, _, ok := runtime.Caller(opts.CallerSkip)
	require.True(t, ok)

	path := filepath.Join(filepath.Dir(curFile), "testdata", name)

	b, err := os.ReadFile(path)
	require.NoError(t, err)

	if !opts.RetainTrailingNewline {
		// Remove trailing newline if present
		if len(b) > 0 && b[len(b)-1] == '\n' {
			b = b[:len(b)-1]
		}
	}

	if len(opts.TemplateValues) > 0 {
		tpl, err := template.New(path).Parse(string(b))
		require.NoError(t, err, "failed to parse template")

		tpl.Option("missingkey=error")

		var buf bytes.Buffer
		err = tpl.Execute(&buf, opts.TemplateValues)
		require.NoError(t, err, "failed to execute template")

		b = buf.Bytes()
	}

	return b
}
