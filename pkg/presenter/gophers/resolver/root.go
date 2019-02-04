package graphResolver

import "github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"

type QueryResolver struct {
	Options []*domainHotelCommon.Option
}
func (r *QueryResolver) HotelX() *HotelXQueryResolver {

	return &HotelXQueryResolver{r.Options}
}

func (r *QueryResolver) Search() *SearchResolver {
	panic("not impl")
}
