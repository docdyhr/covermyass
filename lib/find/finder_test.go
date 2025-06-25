package find

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/sundowndev/covermyass/v2/lib/filter"
	"os"
	"path/filepath"
	"testing"
)

func createFile(t *testing.T, path string) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("data"), 0o644); err != nil {
		t.Fatal(err)
	}
}

func pathsFromResults(results []FileInfo) []string {
	out := make([]string, 0, len(results))
	for _, r := range results {
		out = append(out, r.Path())
	}
	return out
}

func TestFinder_Run(t *testing.T) {
	ctx := context.Background()
	tmp := t.TempDir()

	createFile(t, filepath.Join(tmp, "abs", "file.log"))
	createFile(t, filepath.Join(tmp, "abs", "sub", "inner.log"))
	createFile(t, filepath.Join(tmp, "abs", "match.txt"))

	eng := filter.NewEngine()
	// exclude match.txt
	assert.NoError(t, eng.AddRule("/abs/match.txt"))

	f := New(os.DirFS(tmp), eng)

	t.Run("absolute pattern", func(t *testing.T) {
		res, err := f.Run(ctx, []string{"/abs/*.log"})
		assert.NoError(t, err)
		assert.ElementsMatch(t, []string{"/abs/file.log"}, pathsFromResults(res))
	})

	t.Run("relative pattern", func(t *testing.T) {
		res, err := f.Run(ctx, []string{"abs/sub/*.log"})
		assert.NoError(t, err)
		assert.ElementsMatch(t, []string{"/abs/sub/inner.log"}, pathsFromResults(res))
	})

	t.Run("filtered path", func(t *testing.T) {
		res, err := f.Run(ctx, []string{"/abs/match.txt"})
		assert.NoError(t, err)
		assert.Empty(t, res)
	})
}
