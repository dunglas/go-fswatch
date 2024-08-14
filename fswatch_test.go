package fswatch_test

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/dunglas/go-fswatch"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateFile(t *testing.T) {
	tmp, err := os.MkdirTemp("", "fswatch")
	require.NoError(t, err)

	tmp, _ = filepath.EvalSymlinks(tmp)
	foo := path.Join(tmp, "foo")
	bar := path.Join(tmp, "bar")

	var (
		wg sync.WaitGroup
		i  int
	)
	wg.Add(2)

	s, err := fswatch.NewSession(
		[]string{tmp},
		func(e []fswatch.Event) {
			assert.NotEmpty(t, e)
			assert.LessOrEqual(t, e[0].Time, time.Now())

			path, _ := filepath.EvalSymlinks(e[0].Path)

			switch i {
			case 0:
				assert.Equal(t, tmp, path)
				assert.Contains(t, e[0].Types, fswatch.IsDir)

			case 1:
				assert.Equal(t, foo, path)
				assert.Contains(t, e[0].Types, fswatch.Created)
			}

			i++
			wg.Done()
		},
		fswatch.WithMonitorType(fswatch.SystemDefaultMonitor),
		fswatch.WithAllowOverflow(true),
		fswatch.WithLatency(1),
		fswatch.WithRecursive(true),
		fswatch.WithDirectoryOnly(true),
		fswatch.WithFollowSymlinks(true),
		fswatch.WithEventTypeFilters([]fswatch.EventType{fswatch.Created, fswatch.Updated, fswatch.IsDir, fswatch.IsFile}),
		fswatch.WithFilters([]fswatch.Filter{{Text: "bar$", FilterType: fswatch.FilterExclude, CaseSensitive: false, Extended: false}}),
		fswatch.WithProperties(map[string]string{"foo": "bar"}),
	)
	require.NoError(t, err)

	wg.Add(1)
	go func() {
		require.NoError(t, s.Start())

		wg.Done()
	}()

	time.Sleep(5 * time.Second)

	fooFile, err := os.Create(foo)
	require.NoError(t, err)
	require.NoError(t, fooFile.Close())

	barFile, err := os.Create(bar)
	require.NoError(t, err)
	require.NoError(t, barFile.Close())

	require.NoError(t, s.Stop())

	time.Sleep(3 * time.Second)
	require.NoError(t, s.Destroy())

	wg.Wait()
	time.Sleep(2 * time.Second)
}

func Example() {
	s, err := fswatch.NewSession([]string{"/tmp"}, func(e []fswatch.Event) {
		fmt.Printf("%s", filepath.Base(e[0].Path))
	})
	if err != nil {
		panic(err)
	}

	// Start() is blocking, it must be called in a dedicated goroutine
	go func() {
		if err := s.Start(); err != nil {
			panic(err)
		}
	}()

	// Give some time to the monitor to start, this is a limitation of the underlying C library
	time.Sleep(5 * time.Second)

	f, err := os.Create("/tmp/foo.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := s.Stop(); err != nil {
		panic(err)
	}
	// Give some time to the monitor to stop, this is a limitation of the underlying C library
	time.Sleep(3 * time.Second)

	if err := s.Destroy(); err != nil {
		panic(err)
	}

	// Output:
	// foo.txt
}
