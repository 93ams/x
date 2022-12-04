package integration

import (
	"context"
	"github.com/stretchr/testify/require"
	. "github.com/tilau2328/cql/package/adaptor/data/cql/repo/ddl"
	. "github.com/tilau2328/cql/package/domain/model"
	. "github.com/tilau2328/cql/package/shared/patch"
	"strconv"
	"testing"
	"time"
)

func TestKeySpaceRepo(t *testing.T) {
	ksRepo := NewKeySpaceRepo(session)
	tRepo := NewTableRepo(session)
	r := int(time.Now().Unix())
	k1, k2 := "test"+strconv.Itoa(r), "test"+strconv.Itoa(r+1)
	tests := []struct {
		name        string
		ksCreate    []KeySpace
		ksGet       map[KeySpaceKey]KeySpace
		ksAlter     map[KeySpaceKey][]Patch
		ksDrop      []KeySpaceKey
		ksList      []KeySpace
		tableCreate []Table
		tableGet    map[TableKey]Table
		tableAlter  map[TableKey][]Patch
		tableDrop   []TableKey
		tableList   []Table
	}{
		{
			name: "",
			ksCreate: []KeySpace{
				{
					KeySpaceKey: KeySpaceKey(k1),
					Replication: NewSimpleReplication(1),
					Durable:     true,
				},
				{
					KeySpaceKey: KeySpaceKey(k2),
					Replication: NewSimpleReplication(2),
				},
			},
			ksGet: map[KeySpaceKey]KeySpace{
				KeySpaceKey(k1): {
					KeySpaceKey: KeySpaceKey(k1),
					Replication: NewSimpleReplication(1),
					Durable:     true,
				},
				KeySpaceKey(k2): {
					KeySpaceKey: KeySpaceKey(k2),
					Replication: NewSimpleReplication(2),
				},
			},
			ksAlter: map[KeySpaceKey][]Patch{},
			ksDrop:  []KeySpaceKey{KeySpaceKey(k1)},
			ksList: []KeySpace{
				{
					KeySpaceKey: KeySpaceKey(k2),
					Replication: NewSimpleReplication(2),
				},
			},
			tableCreate: []Table{},
			tableGet:    map[TableKey]Table{},
			tableAlter:  map[TableKey][]Patch{},
			tableDrop:   []TableKey{},
			tableList:   []Table{},
		},
		{
			name: "other",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			t.Cleanup(func() {
				res, err := ksRepo.List(ctx, KeySpace{})
				require.NoError(t, err)
				for _, ks := range FilterSystem(res) {
					require.NoError(t, ksRepo.Drop(ctx, ks.KeySpaceKey))
				}
			})
			for _, v := range tt.ksCreate {
				require.NoError(t, ksRepo.Create(ctx, v))
			}
			for k, v := range tt.ksGet {
				res, err := ksRepo.Get(ctx, k)
				require.NoError(t, err)
				require.Equal(t, v, res)
			}
			for k, v := range tt.ksAlter {
				require.NoError(t, ksRepo.Alter(ctx, k, v))
			}
			for _, k := range tt.ksDrop {
				require.NoError(t, ksRepo.Drop(ctx, k))
			}
			res, err := ksRepo.List(ctx, KeySpace{})
			require.NoError(t, err)
			require.Equal(t, tt.ksList, FilterSystem(res))
			for _, v := range tt.tableCreate {
				require.NoError(t, tRepo.Create(ctx, v))
			}
			for k, v := range tt.tableGet {
				res, err := tRepo.Get(ctx, k)
				require.NoError(t, err)
				require.Equal(t, v, res)
			}
			for k, v := range tt.tableAlter {
				require.NoError(t, tRepo.Alter(ctx, k, v))
			}
			for _, k := range tt.tableDrop {
				require.NoError(t, tRepo.Drop(ctx, k))
			}
			res2, err := tRepo.List(ctx, Table{})
			require.NoError(t, err)
			require.Equal(t, tt.tableList, res2)
		})
	}
}
