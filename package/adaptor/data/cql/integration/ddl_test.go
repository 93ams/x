package integration

import (
	"github.com/stretchr/testify/require"
	"github.com/tilau2328/cql/package/adaptor/data/cql/repo/ddl"
	"github.com/tilau2328/cql/package/domain/model"
	"testing"
)

func TestNewKeySpaceRepo(t *testing.T) {
	ksRepo := ddl.NewKeySpaceRepo(session)
	tRepo := ddl.NewTableRepo(session)
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ksRepo.Create(nil, model.KeySpace{})
			require.NoError(t, err)
			err = ksRepo.Create(nil, model.KeySpace{})
			require.NoError(t, err)
			_, err = ksRepo.Get(nil, "")
			require.NoError(t, err)
			err = ksRepo.Alter(nil, "", nil)
			require.NoError(t, err)
			err = ksRepo.Drop(nil, "")
			require.NoError(t, err)
			_, err = ksRepo.List(nil, model.KeySpace{})
			require.NoError(t, err)
			err = tRepo.Create(nil, model.Table{})
			require.NoError(t, err)
			err = tRepo.Create(nil, model.Table{})
			require.NoError(t, err)
			_, err = tRepo.Get(nil, model.TableKey{})
			require.NoError(t, err)
			err = tRepo.Alter(nil, model.TableKey{}, nil)
			require.NoError(t, err)
			err = tRepo.Drop(nil, model.TableKey{})
			require.NoError(t, err)
			_, err = tRepo.List(nil, model.Table{})
			require.NoError(t, err)
		})
	}
}
