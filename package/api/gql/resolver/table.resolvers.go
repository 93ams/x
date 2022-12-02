package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tilau2328/cql/package/api/gql/model"
)

// CreateTable is the resolver for the createTable field.
func (r *mutationResolver) CreateTable(ctx context.Context, in model.NewTable) (*model.Table, error) {
	panic(fmt.Errorf("not implemented: CreateTable - createTable"))
}

// AlterTable is the resolver for the alterTable field.
func (r *mutationResolver) AlterTable(ctx context.Context, in model.NewTable) (*model.Table, error) {
	panic(fmt.Errorf("not implemented: AlterTable - alterTable"))
}

// DropTable is the resolver for the dropTable field.
func (r *mutationResolver) DropTable(ctx context.Context, in string) (bool, error) {
	panic(fmt.Errorf("not implemented: DropTable - dropTable"))
}

// Table is the resolver for the table field.
func (r *queryResolver) Table(ctx context.Context, name string) (*model.Table, error) {
	panic(fmt.Errorf("not implemented: Table - table"))
}

// Tables is the resolver for the tables field.
func (r *queryResolver) Tables(ctx context.Context, keyspace string) ([]*model.Table, error) {
	panic(fmt.Errorf("not implemented: Tables - tables"))
}
