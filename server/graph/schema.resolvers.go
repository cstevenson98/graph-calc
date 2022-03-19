package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/cstevenson98/graph-calc/server/pkg/govaluate"

	"github.com/cstevenson98/graph-calc/server/graph/generated"
	"github.com/cstevenson98/graph-calc/server/graph/model"
)

func (r *queryResolver) Plot2d(ctx context.Context, input model.Expression) ([][]*float64, error) {
	expression, err := govaluate.NewEvaluableExpression(input.Body)
	if err != nil {
		return nil, fmt.Errorf("invalid expression: %v", err)
	}
	_, _ = expression.Evaluate(nil)
	// result is now set to "true", the bool value.
	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }