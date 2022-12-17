package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tilau2328/x/src/go/cmd/gql/package/exec"
)

// Noop is the resolver for the noop field.
func (r *mutationResolver) Noop(ctx context.Context) (*bool, error) {
	panic(fmt.Errorf("not implemented: Noop - noop"))
}

// Version is the resolver for the version field.
func (r *queryResolver) Version(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: Version - version"))
}

// Mutation returns exec.MutationResolver implementation.
func (r *Resolver) Mutation() exec.MutationResolver { return &mutationResolver{r} }

// Query returns exec.QueryResolver implementation.
func (r *Resolver) Query() exec.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
