package integration

import (
	"github.com/samber/lo"
	"github.com/scylladb/gocqlx/v2"
	"github.com/tilau2328/cql/package/shared/data/cql"
	"os"
	"testing"
)

var session gocqlx.Session

func TestMain(m *testing.M) {
	main := M{M: m}

	var fn func()
	session, fn = lo.Must2(cql.NewSession(cql.NewCluster(cql.Options{})))
	main.Clean(fn)

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

func (m *M) Clean(fn ...func()) { m.cleanup = append(m.cleanup, fn...) }
func (m *M) Close() {
	for _, fn := range m.cleanup {
		fn()
	}
}
