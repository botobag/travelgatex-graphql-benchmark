package gqlgensm

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
)

type Candidate struct{}

var _ presenter.CandidateHandlerFunc = (*Candidate)(nil)

func (Candidate) HandlerFunc(options []*presenter.Option) (http.HandlerFunc, error) {
	soptions := make([]*domainHotelCommon.Option, len(options))
	for i, o := range options {
		opt := (domainHotelCommon.Option)(*o)
		soptions[i] = &opt
	}
	return handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{soptions}})), nil
}

func (Candidate) UnmarshalOptions(b []byte) ([]*presenter.Option, error) {
	return presenter.JSONUnmarshalOptions(b)
}
