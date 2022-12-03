package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"gql/package/model"
)

// CreateKeySpace is the resolver for the createKeySpace field.
func (r *mutationResolver) CreateKeySpace(ctx context.Context, in model.NewKeyspace) (*model.KeySpace, error) {
	panic(fmt.Errorf("not implemented: CreateKeySpace - createKeySpace"))
}

// AlterKeySpace is the resolver for the alterKeySpace field.
func (r *mutationResolver) AlterKeySpace(ctx context.Context, in model.NewKeyspace) (*model.KeySpace, error) {
	panic(fmt.Errorf("not implemented: AlterKeySpace - alterKeySpace"))
}

// DropKeySpace is the resolver for the dropKeySpace field.
func (r *mutationResolver) DropKeySpace(ctx context.Context, in string) (bool, error) {
	panic(fmt.Errorf("not implemented: DropKeySpace - dropKeySpace"))
}

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

// Keyspace is the resolver for the keyspace field.
func (r *queryResolver) Keyspace(ctx context.Context, name *string) (*model.KeySpace, error) {
	panic(fmt.Errorf("not implemented: Keyspace - keyspace"))
}

// Keyspaces is the resolver for the keyspaces field.
func (r *queryResolver) Keyspaces(ctx context.Context) ([]*model.KeySpace, error) {
	panic(fmt.Errorf("not implemented: Keyspaces - keyspaces"))
}
