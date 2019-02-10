package artemis

import (
	"context"
	"net/http"

	"github.com/botobag/artemis/graphql"
	"github.com/botobag/artemis/graphql/handler"

	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
)

// Candidate implements presenter.CandidateHandlerFunc.
type Candidate struct{}

var _ presenter.CandidateHandlerFunc = (*Candidate)(nil)

// HandlerFunc implements presenter.CandidateHandlerFunc.
func (Candidate) HandlerFunc(options []*presenter.Option) (http.HandlerFunc, error) {
	query, err := graphql.NewObject(&graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"hotelX": {
				Type: hotelXQueryType,
				Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
					return hotelX{
						options: options,
					}, nil
				}),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	schema, err := graphql.NewSchema(&graphql.SchemaConfig{
		Query: query,
	})
	if err != nil {
		return nil, err
	}

	handler, err := handler.New(schema)
	if err != nil {
		return nil, err
	}

	return handler.ServeHTTP, nil
}

// UnmarshalOptions implements presenter.CandidateHandlerFunc.
func (Candidate) UnmarshalOptions(b []byte) ([]*presenter.Option, error) {
	return presenter.JSONUnmarshalOptions(b)
}
