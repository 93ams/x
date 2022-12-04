package model

import (
	"github.com/tilau2328/cql/package/domain/model"
	"strconv"
	"strings"
)

type (
	Replication map[string]string
	KeySpace    struct {
		Name        string      `db:"keyspace_name"`
		Durable     bool        `db:"durable_writes"`
		Replication Replication `db:"replication"`
	}
	Table struct {
		Id                      string            `db:"id"`
		Name                    string            `db:"table_name"`
		Keyspace                string            `db:"keyspace_name"`
		Comment                 string            `db:"comment"`
		SpeculativeRetry        string            `db:"speculative_retry"`
		DefaultTTL              int               `db:"default_time_to_live"`
		Gc                      int               `db:"gc_grace_seconds"`
		MaxIndexInterval        int               `db:"max_index_interval"`
		MinIndexInterval        int               `db:"min_index_interval"`
		FlushPeriod             int               `db:"memtable_flush_period_in_ms"`
		CrcCheckChance          float64           `db:"crc_check_chance"`
		ReadRepairChance        float64           `db:"read_repair_chance"`
		DclocalReadRepairChance float64           `db:"dclocal_read_repair_chance"`
		BloomFilterFpChance     float64           `db:"bloom_filter_fp_chance"`
		Caching                 map[string]string `db:"caching"`
		Compression             map[string]string `db:"compression"`
		Compaction              map[string]string `db:"compaction"`
		Flags                   []string          `db:"flags"`
		Extensions              map[string][]byte `db:"extensions"`
	}
)

func ToKeySpaces(in []KeySpace) []model.KeySpace {
	ret := make([]model.KeySpace, len(in))
	for i, v := range in {
		ret[i] = ToKeySpace(v)
	}
	return ret
}
func ToKeySpace(in KeySpace) model.KeySpace {
	return model.KeySpace{
		Replication: ToReplication(in.Replication),
		KeySpaceKey: model.KeySpaceKey(in.Name),
		Durable:     in.Durable,
	}
}
func FromKeySpace(in model.KeySpace) KeySpace {
	return KeySpace{
		Replication: FromReplication(in.Replication),
		Name:        string(in.KeySpaceKey),
		Durable:     in.Durable,
	}
}
func ToTables(in []Table) []model.Table {
	ret := make([]model.Table, len(in))
	for i, v := range in {
		ret[i] = ToTable(v)
	}
	return ret
}
func ToTable(in Table) model.Table {
	return model.Table{}
}
func FromTable(in model.Table) Table {
	return Table{}
}
func ToReplication(in Replication) model.Replication {
	ret := model.Replication{}
	for k, v := range in {
		switch k {
		case "class":
			ret[k] = strings.TrimPrefix(v, "org.apache.cassandra.locator.")
		default:
			ret[k], _ = strconv.Atoi(v)
		}
	}
	return ret
}
func FromReplication(in model.Replication) Replication {
	ret := Replication{}
	for k, v := range in {
		switch k {
		case "class":
			if i, ok := v.(string); ok {
				ret[k] = i
			}
		default:
			if i, ok := v.(int); ok {
				ret[k] = strconv.Itoa(i)
			}
		}
	}
	return ret
}
