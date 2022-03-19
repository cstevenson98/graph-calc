package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math"

	"github.com/cstevenson98/graph-calc/server/graph/generated"
	"github.com/cstevenson98/graph-calc/server/graph/model"
	"github.com/cstevenson98/graph-calc/server/pkg/govaluate"
)

func (r *queryResolver) Plot2d(ctx context.Context, expression model.Expression, rangeArg model.XRange) ([]*float64, error) {
	exp, err := govaluate.NewEvaluableExpression(expression.Body)
	if err != nil {
		return nil, fmt.Errorf("invalid expression: %v", err)
	}

	parameters := make(map[string]interface{}, 8)
	var out []*float64

	for j := 0; j < rangeArg.N + 1; j++ {
		localX := rangeArg.Xmin + (float64)(j)*(rangeArg.Xmax-rangeArg.Xmin)/(float64)(rangeArg.N)
		parameters["x"] = localX

		result, err := exp.Evaluate(parameters)
		if err != nil {
			return nil, err
		}

		switch i := result.(type) {
		case float64:
			localY := i
			out = append(out, &localX)
			out = append(out, &localY)
		default:
			nan := math.NaN()
			out = append(out, &localX)
			out = append(out, &nan)
		}
	}

	return out, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
