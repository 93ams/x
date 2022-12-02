package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tilau2328/cql/package/api/gql/exec"
)

// Insert is the resolver for the insert field.
func (r *mutationResolver) Insert(ctx context.Context, query *string) (*string, error) {
	panic(fmt.Errorf("not implemented: Insert - insert"))
}

// Update is the resolver for the update field.
func (r *mutationResolver) Update(ctx context.Context, query *string) (*string, error) {
	panic(fmt.Errorf("not implemented: Update - update"))
}

// Delete is the resolver for the delete field.
func (r *mutationResolver) Delete(ctx context.Context, query *string) (*string, error) {
	panic(fmt.Errorf("not implemented: Delete - delete"))
}

// Select is the resolver for the select field.
func (r *queryResolver) Select(ctx context.Context, query *string) (*string, error) {
	panic(fmt.Errorf("not implemented: Select - select"))
}

// Mutation returns exec.MutationResolver implementation.
func (r *Resolver) Mutation() exec.MutationResolver { return &mutationResolver{r} }

// Query returns exec.QueryResolver implementation.
func (r *Resolver) Query() exec.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
