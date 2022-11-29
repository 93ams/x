package cql

import (
	"context"
	"github.com/samber/lo"
	"github.com/scylladb/gocqlx/v2"
	"github.com/tilau2328/cql/internal/models"
	"github.com/tilau2328/cql/package/cmd/pretty"
	"os"
	"strings"
)

type Table struct {
	Id                      string
	KeyspaceName            string
	TableName               string
	Comment                 string
	SpeculativeRetry        string
	DefaultTimeToLive       int
	GcGraceSeconds          int
	MaxIndexInterval        int
	MinIndexInterval        int
	MemtableFlushPeriodInMs int
	CrcCheckChance          float64
	ReadRepairChance        float64
	DclocalReadRepairChance float64
	BloomFilterFpChance     float64
	Caching                 map[string]string
	Compression             map[string]string
	Compaction              map[string]string
	Flags                   []string
	Extensions              map[string][]byte
}

func ListTables(session gocqlx.Session, props string) error {
	var ret []Table
	if err := models.Tables.SelectQuery(session).BindStruct(Table{
		KeyspaceName: props,
	}).SelectRelease(&ret); err != nil {
		return err
	}
	pretty.NewTable(
		pretty.Header("#", "Id", "KeyspaceName", "TableName", "Comment"),
		pretty.Rows(lo.Map(ret, func(v Table, i int) []any {
			return []any{i, v.Id, v.KeyspaceName, v.TableName, v.Comment}
		})...),
	).Write(os.Stdout)
	return nil
}

func CreateTable(session gocqlx.Session, name string, props []string) error {
	return SafeExec(context.Background(), session.Session, `CREATE table IF NOT EXISTS `+
		name+` (`+strings.Join(props, ",")+`)`)
}
func AlterTable(session gocqlx.Session, name string, props []string) error {
	return SafeExec(context.Background(), session.Session, `ALTER table `+
		name+` (`+strings.Join(props, ",")+`)`)
}
func DropTable(session gocqlx.Session, props string) error {
	return SafeExec(context.Background(), session.Session, `DROP table IF EXISTS `+props)
}
