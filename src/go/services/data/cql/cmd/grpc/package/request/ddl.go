package request

import (
	"context"
	"github.com/tilau2328/cql/src/go/cmd/grpc/package/model"
	. "github.com/tilau2328/cql/src/go/package/domain/model"
	. "github.com/tilau2328/cql/src/go/package/domain/provider"
)

type DDLRequester struct {
	client model.DDLClient
}

var _ DDL = &DDLRequester{}

func NewDDLRequester(c *Client) *DDLRequester {
	return &DDLRequester{
		client: model.NewDDLClient(c.conn),
	}
}

func (r *DDLRequester) ListKeySpaces(ctx context.Context, in KeySpace) ([]KeySpace, error) {
	r.client.ListKeySpaces(ctx, model.To)
	return nil, nil
}
func (r *DDLRequester) GetKeySpace(ctx context.Context, key KeySpaceKey) (KeySpace, error) {
	return KeySpace{}, nil
}
func (r *DDLRequester) CreateKeySpace(ctx context.Context, in KeySpace) error {
	return nil
}
func (r *DDLRequester) AlterKeySpace(ctx context.Context, key KeySpaceKey, patches []Patch) error {
	return nil
}
func (r *DDLRequester) DropKeySpace(ctx context.Context, key KeySpaceKey) error {
	return nil
}
func (r *DDLRequester) ListTables(ctx context.Context, table Table) ([]Table, error) {
	return nil, nil
}
func (r *DDLRequester) GetTable(ctx context.Context, key TableKey) (Table, error) {
	return Table{}, nil
}
func (r *DDLRequester) CreateTable(ctx context.Context, table Table) error {
	return nil
}
func (r *DDLRequester) AlterTable(ctx context.Context, key TableKey, patches []Patch) error {
	return nil
}
func (r *DDLRequester) DropTable(ctx context.Context, key TableKey) error {
	return nil
}
