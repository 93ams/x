package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tilau2328/cql/package/api/gql/model"
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

// Keyspace is the resolver for the keyspace field.
func (r *queryResolver) Keyspace(ctx context.Context, name *string) (*model.KeySpace, error) {
	panic(fmt.Errorf("not implemented: Keyspace - keyspace"))
}

// Keyspaces is the resolver for the keyspaces field.
func (r *queryResolver) Keyspaces(ctx context.Context) ([]*model.KeySpace, error) {
	panic(fmt.Errorf("not implemented: Keyspaces - keyspaces"))
}
