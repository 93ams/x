package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gql/package/exec"
	"gql/package/model"
)

// CreateTable is the resolver for the createTable field.
func (r *mutationResolver) CreateTable(ctx context.Context, in model.NewTable) (*model.Table, error) {
	panic(fmt.Errorf("not implemented: CreateTable - createTable"))
}

// DropTable is the resolver for the dropTable field.
func (r *mutationResolver) DropTable(ctx context.Context, in string) (bool, error) {
	panic(fmt.Errorf("not implemented: DropTable - dropTable"))
}

// Keyspaces is the resolver for the keyspaces field.
func (r *queryResolver) Keyspaces(ctx context.Context) ([]*model.KeySpace, error) {
	panic(fmt.Errorf("not implemented: Keyspaces - keyspaces"))
}

// Tables is the resolver for the tables field.
func (r *queryResolver) Tables(ctx context.Context, keyspace string) ([]*model.Table, error) {
	panic(fmt.Errorf("not implemented: Tables - tables"))
}

// Mutation returns exec.MutationResolver implementation.
func (r *Resolver) Mutation() exec.MutationResolver { return &mutationResolver{r} }

// Query returns exec.QueryResolver implementation.
func (r *Resolver) Query() exec.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
