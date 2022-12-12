package integration

import (
	"github.com/gocql/gocql"
	"github.com/samber/lo"
	"github.com/scylladb/gocqlx/v2"
	"os"
	"testing"
)

var session gocqlx.Session

func TestMain(m *testing.M) {
	main := M{M: m}

	var fn func()
	session, fn = lo.Must2(NewSession(NewCluster(Options{Consistency: gocql.All})))
	main.Defer(fn)

	os.Exit(main.Run())
}

type M struct {
	*testing.M
	cleanup []func()
}

func (m *M) Run() int {
	defer m.Close()
	return m.M.Run()
}

func (m *M) Defer(fn ...func()) { m.cleanup = append(m.cleanup, fn...) }
func (m *M) Close() {
	for _, fn := range m.cleanup {
		fn()
	}
}
