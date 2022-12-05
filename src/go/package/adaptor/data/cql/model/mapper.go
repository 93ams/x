package model

import (
	"github.com/tilau2328/cql/package/domain/model"
	"github.com/tilau2328/cql/package/shared/data/cql/ddl"
	"strconv"
	"strings"
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
	return model.Table{
		Id: in.Id,
		TableKey: model.TableKey{
			Name:     in.Name,
			KeySpace: in.KeySpace,
		},
		Comment:                 in.Comment,
		SpeculativeRetry:        in.SpeculativeRetry,
		DefaultTTL:              in.DefaultTTL,
		Gc:                      in.Gc,
		MaxIndexInterval:        in.MaxIndexInterval,
		MinIndexInterval:        in.MinIndexInterval,
		FlushPeriod:             in.FlushPeriod,
		CrcCheckChance:          in.CrcCheckChance,
		ReadRepairChance:        in.ReadRepairChance,
		DclocalReadRepairChance: in.DclocalReadRepairChance,
		BloomFilterFpChance:     in.BloomFilterFpChance,
		Caching:                 in.Caching,
		Compression:             in.Compression,
		Compaction:              in.Compaction,
		Flags:                   in.Flags,
		Extensions:              in.Extensions,
	}
}
func FromTable(in model.Table) Table {
	return Table{
		Id:                      in.Id,
		Name:                    in.Name,
		KeySpace:                in.KeySpace,
		Comment:                 in.Comment,
		SpeculativeRetry:        in.SpeculativeRetry,
		DefaultTTL:              in.DefaultTTL,
		Gc:                      in.Gc,
		MaxIndexInterval:        in.MaxIndexInterval,
		MinIndexInterval:        in.MinIndexInterval,
		FlushPeriod:             in.FlushPeriod,
		CrcCheckChance:          in.CrcCheckChance,
		ReadRepairChance:        in.ReadRepairChance,
		DclocalReadRepairChance: in.DclocalReadRepairChance,
		BloomFilterFpChance:     in.BloomFilterFpChance,
		Caching:                 in.Caching,
		Compression:             in.Compression,
		Compaction:              in.Compaction,
		Flags:                   in.Flags,
		Extensions:              in.Extensions,
	}
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

func ToColumns(in []Column) []model.Column {
	ret := make([]model.Column, len(in))
	for i, v := range in {
		ret[i] = ToColumn(v)
	}
	return ret
}

func ToColumn(in Column) model.Column {
	return model.Column{
		ColumnKey: model.ColumnKey{
			KeySpace: in.KeySpace,
			Table:    in.Table,
			Name:     in.Name,
		},
		Order:     in.Order,
		NameBytes: in.NameBytes,
		Kind:      in.Kind,
		Pos:       in.Pos,
		Type:      in.Type,
	}
}

func FromColumn(in model.Column) Column {
	return Column{
		KeySpace:  in.KeySpace,
		Table:     in.Table,
		Name:      in.Name,
		Order:     in.Order,
		NameBytes: in.NameBytes,
		Kind:      in.Kind,
		Pos:       in.Pos,
		Type:      in.Type,
	}
}
func MapCols(c []model.Column) []ddl.Column {
	ret := make([]ddl.Column, len(c))
	for i, v := range c {
		ret[i] = NewCol(v)
	}
	return ret
}
func NewCol(column model.Column) ddl.Column {
	return ddl.Column{
		Name:    column.Name,
		Type:    column.Type,
		Primary: column.Primary,
		Static:  column.Static,
	}
}
