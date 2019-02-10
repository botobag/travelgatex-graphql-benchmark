package artemis

import (
	"context"

	"github.com/botobag/artemis/graphql"
	"github.com/travelgateX/presenters-benchmark/pkg/access"
	"github.com/travelgateX/presenters-benchmark/pkg/common"
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
	"github.com/travelgateX/presenters-benchmark/pkg/search"
)

var dateType = &graphql.ScalarAliasConfig{
	AliasFor: graphql.String(),
}

var currencyType = &graphql.ScalarAliasConfig{
	AliasFor: graphql.String(),
}

var jsonType = &graphql.ScalarAliasConfig{
	AliasFor: graphql.String(),
}

type hotelX struct {
	options []*presenter.Option
}

func (hotel hotelX) Search(ctx context.Context) (*hotelSearch, error) {
	options := make([]*domainHotelCommon.Option, len(hotel.options))
	for i, o := range hotel.options {
		options[i] = (*domainHotelCommon.Option)(o)
	}

	return &hotelSearch{
		response: &search.HotelSearchRS{
			Options: options,
		},
	}, nil
}

var hotelXQueryType = &graphql.ObjectConfig{
	Name: "HotelXQuery",
	Fields: graphql.Fields{
		"search": {
			Type: hotelSearchType,
			Description: `Available options of an hotel for a given date and itinerary. It does not filter different classes, times or
fares. It will always retrieve all results returned by the suppliers. The availability request is very straight
forward. It only requires the criteria of search (destination, travel dates and the number of pax in each room).
But you must preload the other fields in our system by complete the fields absents.`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(hotelX).Search(ctx)
			}),
		},

		/*
			"searchStatusService": {
				Type:        graphql.NonNullOf(ServiceStatus),
				Description: "Returns status of the search service.",
			},
		*/
	},
}

var responseInterface = &graphql.InterfaceConfig{
	Name: "Response",
	Fields: graphql.Fields{
		// # Application stats
		// stats(token: String!): StatsRequest

		// # Data sent and received in the supplier’s original format.
		// auditData: AuditData

		"errors": {
			Type: graphql.ListOf(graphql.NonNullOf(errorType)),
		},

		// # Potentially harmful situations or errors that do not stop the service
		// warnings: [Warning!]
	},
}

var errorType = &graphql.ObjectConfig{
	Name:        "Error",
	Description: "Application errors",
	Fields: graphql.Fields{
		"code": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Error code",
		},
		"type": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Error type",
		},
		"description": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Error description",
		},
	},
}

type hotelSearch struct {
	request  *search.HotelSearchRQ
	response *search.HotelSearchRS

	errors   []*common.AdviseMessage
	warnings []*common.AdviseMessage

	serializedCriteria string
}

func (search *hotelSearch) Context() *string {
	request := search.request
	if request != nil && request.Settings != nil {
		return request.Settings.Context
	}
	return nil
}

func (search *hotelSearch) Options() []*hotelOptionSearch {
	response := search.response
	if response == nil || response.Options == nil {
		return nil
	}

	options := response.Options
	criteria := search.serializedCriteria
	result := make([]*hotelOptionSearch, 0, len(options))
	for _, option := range options {
		result = append(result, &hotelOptionSearch{
			option:   option,
			criteria: criteria,
		})
	}

	return result
}

func (search *hotelSearch) Errors() []*common.AdviseMessage {
	return search.errors
}

var hotelSearchType = &graphql.ObjectConfig{
	Name:        "HotelSearch",
	Description: "Results from Avail Hotel; contains all the available options for a given date and itinerary",
	Interfaces: []graphql.InterfaceTypeDefinition{
		responseInterface,
	},
	Fields: graphql.Fields{
		"context": {
			Type:        graphql.T(graphql.String()),
			Description: "Indicates the context of the response.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelSearch).Context(), nil
			}),
		},

		// # Application stats in string format
		// stats(token: String!): StatsRequest

		// # Data sent and received in the supplier's native format.
		// auditData: AuditData

		// # Request Criteria
		// requestCriteria: CriteriaSearch

		"options": {
			Type:        graphql.ListOf(graphql.NonNullOf(hotelOptionSearchType)),
			Description: "List of options returned according to the request.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelSearch).Options(), nil
			}),
		},

		"errors": {
			Type:        graphql.ListOf(graphql.NonNullOf(errorType)),
			Description: "Errors that abort services",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelSearch).Errors(), nil
			}),
		},

		// # Potentially harmful situations or errors that won't force the service to abort
		// warnings: [Warning!]
	},
}

var bookableOptionSearch = &graphql.InterfaceConfig{
	Name: "BookableOptionSearch",
	Fields: graphql.Fields{
		"supplierCode": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Supplier that offers this option.",
		},
		"accessCode": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Access code of this option.",
		},
		"id": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Indicates the id to be used on Quote as key",
		},
	},
}

type hotelOptionSearch struct {
	option   *domainHotelCommon.Option
	criteria string
}

func (search *hotelOptionSearch) SupplierCode() string {
	return search.option.Supplier
}

func (search *hotelOptionSearch) AccessCode() string {
	return search.option.Access
}

func (search *hotelOptionSearch) Market() string {
	return search.option.Market
}

func (search *hotelOptionSearch) HotelCode() string {
	return search.option.HotelCode
}

func (search *hotelOptionSearch) HotelCodeSupplier() string {
	return search.option.Id.HotelCode
}

func (search *hotelOptionSearch) HotelName() *string {
	hotelName := search.option.HotelName
	if hotelName == nil || len(*hotelName) == 0 {
		return nil
	}
	return hotelName
}

func (search *hotelOptionSearch) BoardCode() string {
	return *search.option.BoardCode
}

func (search *hotelOptionSearch) BoardCodeSupplier() string {
	return search.option.Id.BoardCode
}

func (search *hotelOptionSearch) PaymentType() domainHotelCommon.PaymentType {
	return search.option.PaymentType
}

func (search *hotelOptionSearch) Status() domainHotelCommon.StatusType {
	return search.option.Status
}

func (search *hotelOptionSearch) Occupancies() []*domainHotelCommon.Occupancy {
	occupancies := search.option.Occupancies
	result := make([]*domainHotelCommon.Occupancy, len(occupancies))
	for i := range occupancies {
		result[i] = &occupancies[i]
	}
	return result
}

func (search *hotelOptionSearch) Rooms() []*domainHotelCommon.Room {
	rooms := search.option.Rooms
	result := make([]*domainHotelCommon.Room, len(rooms))
	for i := range rooms {
		result[i] = &rooms[i]
	}
	return result
}

func (search *hotelOptionSearch) Price() *domainHotelCommon.Price {
	return &search.option.Price
}

func (search *hotelOptionSearch) Supplements() []*domainHotelCommon.Supplement {
	return search.option.Supplements
}

func (search *hotelOptionSearch) Surcharges() []*domainHotelCommon.Surcharge {
	surcharges := search.option.Surcharges
	if len(surcharges) == 0 {
		return nil
	}

	result := make([]*domainHotelCommon.Surcharge, len(surcharges))
	for i := range surcharges {
		result[i] = &surcharges[i]
	}
	return result
}

func (search *hotelOptionSearch) RateRules() []access.RateRulesType {
	rateRules := search.option.RateRules
	if len(rateRules) == 0 {
		return nil
	}
	return rateRules
}

func (search *hotelOptionSearch) CancelPolicy() *domainHotelCommon.CancelPolicy {
	return search.option.CancelPolicy
}

func (search *hotelOptionSearch) Remarks() *string {
	return search.option.Remarks
}

func (search *hotelOptionSearch) Token() string {
	return search.option.Token
}

func (search *hotelOptionSearch) ID() string {
	return search.option.OptionID
}

var hotelOptionSearchType = &graphql.ObjectConfig{
	Name:        "HotelOptionSearch",
	Description: "An option includes hotel information, meal plan, total price, conditions and room description",
	Interfaces: []graphql.InterfaceTypeDefinition{
		bookableOptionSearch,
	},
	Fields: graphql.Fields{
		"supplierCode": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Supplier that offers this option.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).SupplierCode(), nil
			}),
		},
		"accessCode": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Access code of this option.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).AccessCode(), nil
			}),
		},

		"market": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Market of this option.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).Market(), nil
			}),
		},

		"hotelCode": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Code of the hotel in the context selected.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).HotelCode(), nil
			}),
		},

		"hotelCodeSupplier": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Supplier's hotel code.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).HotelCodeSupplier(), nil
			}),
		},

		"hotelName": {
			Type:        graphql.T(graphql.String()),
			Description: "Name of the hotel.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).HotelName(), nil
			}),
		},

		"boardCode": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Code of the board in the context selected.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).BoardCode(), nil
			}),
		},

		"boardCodeSupplier": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Supplier's board code.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).BoardCodeSupplier(), nil
			}),
		},

		"paymentType": {
			Type:        graphql.NonNullOf(paymentTypeEnum),
			Description: "Indicates the payment type of the option returned. Possible options: Merchant, Direct, Card Booking, Card check in and Mixed.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).PaymentType(), nil
			}),
		},

		"status": {
			Type:        graphql.NonNullOf(statusTypeEnum),
			Description: "The possible values in status in response are Available (OK) or On Request (RQ).",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).Status(), nil
			}),
		},

		"occupancies": {
			Type:        graphql.ListOf(graphql.NonNullOf(occupancyType)),
			Description: "List of occupancies for the request",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).Occupancies(), nil
			}),
		},

		"rooms": {
			Type:        graphql.ListOf(graphql.NonNullOf(roomType)),
			Description: "List of rooms of the option returned.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).Rooms(), nil
			}),
		},

		"price": {
			Type:        graphql.NonNullOf(priceType),
			Description: "Specifies the prices (Gross, Net and Amount) of the option returned.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).Price(), nil
			}),
		},

		"supplements": {
			Type:        graphql.ListOf(graphql.NonNullOf(supplementType)),
			Description: "List of supplements of the option returned.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).Supplements(), nil
			}),
		},

		"surcharges": {
			Type:        graphql.ListOf(graphql.NonNullOf(surchargeType)),
			Description: "List of surcharges of the option returned.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).Surcharges(), nil
			}),
		},

		"rateRules": {
			Type:        graphql.ListOf(graphql.NonNullOf(rateRulesTypeEnum)),
			Description: "Specifies rate rules of the option returned.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).RateRules(), nil
			}),
		},

		"cancelPolicy": {
			Type:        cancelPolicyType,
			Description: "Specifies cancel policies of the option returned.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).CancelPolicy(), nil
			}),
		},

		"remarks": {
			Type:        graphql.T(graphql.String()),
			Description: "Additional information about the option.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).Remarks(), nil
			}),
		},

		"addOns": {
			Type:        addOnsType,
			Description: "Additional information about the option",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return nil, nil
			}),
		},

		"token": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Token for Deep Link",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).Token(), nil
			}),
		},

		"id": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Access code of this option.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*hotelOptionSearch).ID(), nil
			}),
		},
	},
}

var paymentTypeEnum = &graphql.EnumConfig{
	Name:        "PaymentType",
	Description: "Options payment type",
	Values: graphql.EnumValueDefinitionMap{
		"MERCHANT": {
			Description: "The payment is managed by the supplier.",
		},
		"DIRECT": {
			Description: "The payment is made straight to the actual payee, without sending it through an intermediary or a third party.",
		},
		"CARD_BOOKING": {
			Description: "The payment is managed by the supplier. The payment is effectuated at the time of booking.",
		},
		"CARD_CHECK_IN": {
			Description: "The payment is managed by the supplier. The payment is effectuated at check in in the hotel.",
		},
	},
}

var statusTypeEnum = &graphql.EnumConfig{
	Name:        "StatusType",
	Description: "Indicartes options status",
	Values: graphql.EnumValueDefinitionMap{
		"OK": {
			Description: "The status of the avail is available",
		},
		"RQ": {
			Description: "The status of the avail is On request",
		},
	},
}

var occupancyType = &graphql.ObjectConfig{
	Name:        "Occupancy",
	Description: "Information about occupancy.",
	Fields: graphql.Fields{
		"id": {
			Type:        graphql.NonNullOfType(graphql.Int()),
			Description: "Unique ID room in this option.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Occupancy).Id, nil
			}),
		},
		"paxes": {
			Type:        graphql.NonNullOf(graphql.ListOf(graphql.NonNullOf(paxType))),
			Description: "List of pax of this occupancy.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				paxes := source.(*domainHotelCommon.Occupancy).Paxes
				result := make([]*domainHotelCommon.Pax, len(paxes))
				for i := range paxes {
					result[i] = &paxes[i]
				}
				return result, nil
			}),
		},
	},
}

var paxType = &graphql.ObjectConfig{
	Name:        "Pax",
	Description: "Specifies the age pax. The range of what is considered an adult, infant or baby is particular to each supplier.",
	Fields: graphql.Fields{
		"age": {
			Type:        graphql.NonNullOfType(graphql.Int()),
			Description: "Specifies the age pax.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Pax).Age, nil
			}),
		},
	},
}

var roomType = &graphql.ObjectConfig{
	Name:        "Room",
	Description: "Contains the room information of the option returned.",
	Fields: graphql.Fields{
		"occupancyRefId": {
			Type:        graphql.NonNullOfType(graphql.Int()),
			Description: "ID reference to the occupancy",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Room).OccupancyRefID, nil
			}),
		},
		"code": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Indicates the room code",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Room).Code, nil
			}),
		},
		"description": {
			Type:        graphql.T(graphql.String()),
			Description: "Description about the room",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Room).Description, nil
			}),
		},
		"refundable": {
			Type:        graphql.T(graphql.Boolean()),
			Description: "Identifies if the room is refundable or not.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Room).Refundable, nil
			}),
		},
		"units": {
			Type:        graphql.T(graphql.Int()),
			Description: "Identifies if the room is refundable or not.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Room).Units, nil
			}),
		},
		"roomPrice": {
			Type:        graphql.NonNullOf(roomPriceType),
			Description: "Specifies the room price.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return &source.(*domainHotelCommon.Room).RoomPrice, nil
			}),
		},
		"beds": {
			Type:        graphql.ListOf(graphql.NonNullOf(bedType)),
			Description: "List of beds.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Room).Beds, nil
			}),
		},
		"ratePlans": {
			Type:        graphql.ListOf(graphql.NonNullOf(ratePlanType)),
			Description: "Daily break downs rate plan.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Room).RatePlans, nil
			}),
		},
		"promotions": {
			Type:        graphql.ListOf(graphql.NonNullOf(promotionType)),
			Description: "Daily break downs promotions.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Room).Promotions, nil
			}),
		},
	},
}

var roomPriceType = &graphql.ObjectConfig{
	Name:        "RoomPrice",
	Description: "Specifies the room price.",
	Fields: graphql.Fields{
		"price": {
			Type:        graphql.NonNullOf(priceType),
			Description: "Total price for all days.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return &source.(*domainHotelCommon.RoomPrice).Price, nil
			}),
		},
		"breakdown": {
			Type:        graphql.ListOf(graphql.NonNullOf(priceBreakdownType)),
			Description: "Daily break downs price.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.RoomPrice).Breakdown, nil
			}),
		},
	},
}

var bedType = &graphql.ObjectConfig{
	Name:        "Bed",
	Description: "Contains information about a bed.",
	Fields: graphql.Fields{
		"type": {
			Type:        graphql.T(graphql.String()),
			Description: "Specified the bed type",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Bed).Type, nil
			}),
		},
		"description": {
			Type:        graphql.T(graphql.String()),
			Description: "Description about the bed",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Bed).Description, nil
			}),
		},
		"count": {
			Type:        graphql.T(graphql.Int()),
			Description: "Indicates number of beds in a room",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Bed).Count, nil
			}),
		},
		"shared": {
			Type:        graphql.T(graphql.Boolean()),
			Description: "Specifies if the bed is shared or not",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Bed).Shared, nil
			}),
		},
	},
}

var ratePlanType = &graphql.ObjectConfig{
	Name:        "RatePlan",
	Description: "Information about the rate of the option returned.",
	Fields: graphql.Fields{
		"code": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Specifies the rate code.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.RatePlan).Code, nil
			}),
		},
		"name": {
			Type:        graphql.T(graphql.String()),
			Description: "Specifies the rate name.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.RatePlan).Name, nil
			}),
		},
		"effectiveDate": {
			Type:        dateType,
			Description: "Start date in which the rate becomes effective.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.RatePlan).EffectiveDate, nil
			}),
		},
		"expireDate": {
			Type:        dateType,
			Description: "Expire date of the rate.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.RatePlan).ExpireDate, nil
			}),
		},
	},
}

var promotionType = &graphql.ObjectConfig{
	Name:        "Promotion",
	Description: "Information about room promotions(offers).",
	Fields: graphql.Fields{
		"code": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Specifies the promotion code.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Promotion).Code, nil
			}),
		},
		"name": {
			Type:        graphql.T(graphql.String()),
			Description: "Specifies the promotion name.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Promotion).Name, nil
			}),
		},
		"effectiveDate": {
			Type:        dateType,
			Description: "Promotion effective date.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Promotion).EffectiveDate, nil
			}),
		},
		"expireDate": {
			Type:        dateType,
			Description: "Promotion expire date.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Promotion).ExpireDate, nil
			}),
		},
	},
}

var priceBreakdownType = &graphql.ObjectConfig{
	Name:        "PriceBreakdown",
	Description: "Information about daily price.",
	Fields: graphql.Fields{
		"effectiveDate": {
			Type:        dateType,
			Description: "Start date in which the price becomes effective.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.PriceBreakDown).EffectiveDate, nil
			}),
		},
		"expireDate": {
			Type:        dateType,
			Description: "Expire date of price.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.PriceBreakDown).ExpireDate, nil
			}),
		},
		"price": {
			Type:        graphql.NonNullOf(priceType),
			Description: "Specifies the daily price.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return &source.(*domainHotelCommon.PriceBreakDown).Price, nil
			}),
		},
	},
}

var priceType = &graphql.ObjectConfig{
	Name: "Price",
	Description: `Price indicates the value of the room/option.
Supplements and/or surcharges can be included into the price, and will be verified with nodes Supplements/Surcharges.`,
	Interfaces: []graphql.InterfaceTypeDefinition{
		priceable,
	},
	Fields: graphql.Fields{
		"currency": {
			Type: graphql.NonNullOf(currencyType),
			Description: `Currency code indicating which currency should be paid.
This information is mandatory.`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Price).Currency, nil
			}),
		},
		"binding": {
			Type: graphql.NonNullOfType(graphql.Boolean()),
			Description: `It indicates if the price indicated in the gross must be respected.
That is, the customer can not sell the room / option at a price lower than that established by the supplier.
This information is mandatory.`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Price).Binding, nil
			}),
		},
		"net": {
			Type: graphql.NonNullOfType(graphql.Float()),
			Description: `Indicates the net price that the customer must pay to the supplier.
This information is mandatory.`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Price).Net, nil
			}),
		},
		"gross": {
			Type:        graphql.T(graphql.Float()),
			Description: `Indicates the retail price that the supplier sells to the customer.`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Price).Gross, nil
			}),
		},
		"exchange": {
			Type: graphql.NonNullOf(exchangeType),
			Description: `Provides information about the currency of original, and its rate applied over the results returned by the Supplier.
This information is mandatory.`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return &source.(*domainHotelCommon.Price).Exchange, nil
			}),
		},
		"markups": {
			Type:        graphql.ListOf(graphql.NonNullOf(markupType)),
			Description: `Informs markup applied over supplier price. `,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Price).Markups, nil
			}),
		},
	},
}

var priceable = &graphql.InterfaceConfig{
	Name: "Priceable",
	Fields: graphql.Fields{
		"currency": {
			Type:        graphql.NonNullOf(currencyType),
			Description: "Specifies the currency.",
		},
		"binding": {
			Type:        graphql.NonNullOfType(graphql.Boolean()),
			Description: "Is binding.",
		},
		"net": {
			Type:        graphql.NonNullOfType(graphql.Float()),
			Description: "Specifies the import net.",
		},
		"gross": {
			Type:        graphql.T(graphql.Float()),
			Description: "Specifies the import gross.",
		},
		"exchange": {
			Type:        graphql.NonNullOf(exchangeType),
			Description: "Specifies the exchange.",
		},
	},
}

var exchangeType = &graphql.ObjectConfig{
	Name:        "Exchange",
	Description: "Provides information about the currency of original, and its rate applied over the results returned by the Supplier.",
	Fields: graphql.Fields{
		"currency": {
			Type:        graphql.NonNullOf(currencyType),
			Description: "Provide information about the currency of origin",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Exchange).Currency, nil
			}),
		},
		"rate": {
			Type:        graphql.NonNullOfType(graphql.Float()),
			Description: "Provides information about the rate applied over results",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Exchange).Rate, nil
			}),
		},
	},
}

var markupType = &graphql.ObjectConfig{
	Name:        "Markup",
	Description: "Informs markup applied over supplier price.",
	Interfaces: []graphql.InterfaceTypeDefinition{
		priceable,
	},
	Fields: graphql.Fields{
		"channel": {
			Type:        graphql.T(graphql.String()),
			Description: "channel of markup application.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Markup).Channel, nil
			}),
		},
		"currency": {
			Type: graphql.NonNullOf(currencyType),
			Description: `Currency code indicating which currency should be paid.
This information is mandatory.`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Markup).Currency, nil
			}),
		},
		"binding": {
			Type: graphql.NonNullOfType(graphql.Boolean()),
			Description: `It indicates if the price indicated in the gross must be respected.
That is, the customer can not sell the room / option at a price lower than that established by the supplier.
This information is mandatory.`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Markup).Binding, nil
			}),
		},
		"net": {
			Type: graphql.NonNullOfType(graphql.Float()),
			Description: `Indicates the net price that the customer must pay to the supplier.
This information is mandatory.`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Markup).Net, nil
			}),
		},
		"gross": {
			Type:        graphql.T(graphql.Float()),
			Description: `Indicates the retail price that the supplier sells to the customer.`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Markup).Gross, nil
			}),
		},
		"exchange": {
			Type: graphql.NonNullOf(exchangeType),
			Description: `Provides information about the currency of original, and its rate applied over the results returned by the Supplier.
This information is mandatory.`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return &source.(*domainHotelCommon.Markup).Exchange, nil
			}),
		},
		"rules": {
			Type:        graphql.NonNullOf(graphql.ListOf(graphql.NonNullOf(ruleType))),
			Description: "Breakdown of the applied rules for a markup",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				rules := source.(*domainHotelCommon.Markup).Rules
				result := make([]*domainHotelCommon.Rule, len(rules))
				for i := range rules {
					result[i] = &rules[i]
				}
				return result, nil
			}),
		},
	},
}

var ruleType = &graphql.ObjectConfig{
	Name: "Rule",
	Fields: graphql.Fields{
		"id": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "rule identifier",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return &source.(*domainHotelCommon.Rule).Id, nil
			}),
		},
		"name": {
			Type:        graphql.T(graphql.String()),
			Description: "rule name",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return &source.(*domainHotelCommon.Rule).Name, nil
			}),
		},
		"type": {
			Type:        graphql.NonNullOf(markupRuleTypeEnum),
			Description: "type of the value",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return &source.(*domainHotelCommon.Rule).Type, nil
			}),
		},
		"value": {
			Type:        graphql.NonNullOfType(graphql.Float()),
			Description: "value applied by this rule",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return &source.(*domainHotelCommon.Rule).Value, nil
			}),
		},
	},
}

var markupRuleTypeEnum = &graphql.EnumConfig{
	Name:        "markupRuleType",
	Description: "Indicates what type of value is the markup, by percentage or is an import.",
	Values: graphql.EnumValueDefinitionMap{
		"PERCENT": {
			Description: "Indicates the percentage applied by a rule.",
		},
		"IMPORT": {
			Description: "Indicates the exact amount applied by a rule.",
		},
	},
}

var supplementType = &graphql.ObjectConfig{
	Name:        "Supplement",
	Description: "Supplement that it can be or its already added to the option returned. Contains all the information about the supplement.",
	Fields: graphql.Fields{
		"code": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Specifies the supplement code.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Supplement).Code, nil
			}),
		},
		"name": {
			Type:        graphql.T(graphql.String()),
			Description: "Specifies the supplement name.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Supplement).Name, nil
			}),
		},
		"description": {
			Type:        graphql.T(graphql.String()),
			Description: "Specifies the supplement description.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Supplement).Description, nil
			}),
		},
		"supplementType": {
			Type:        graphql.NonNullOf(supplementTypeEnum),
			Description: "Indicates the supplement type. Possible types: Fee, Ski_pass, Lessons, Meals, Equipment, Ticket, Transfers, Gla, Activity or Null.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Supplement).SupplementType, nil
			}),
		},
		"chargeType": {
			Type: graphql.NonNullOf(chargeTypeEnum),
			Description: `Indicates the charge types. We need to know whether the supplements have to be paid when the consumer gets to the hotel or beforehand.
Possible charge types: Include or Exclude.
when include: this supplement is mandatory and included in the option's price
when exclude: this supplement is not included in the option's price`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Supplement).ChargeType, nil
			}),
		},
		"mandatory": {
			Type: graphql.NonNullOfType(graphql.Boolean()),
			Description: `Indicates if the supplement is mandatory or not. If mandatory, this supplement will be applied to this option
if the chargeType is excluded the customer will have to pay it directly at the hotel`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Supplement).Mandatory, nil
			}),
		},
		"durationType": {
			Type:        durationTypeEnum,
			Description: "Specifies the duration type. Possible duration types: Range (specified dates) or Open. This field is mandatory for PDI.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Supplement).DurationType, nil
			}),
		},
		"quantity": {
			Type:        graphql.T(graphql.Int()),
			Description: "Indicates the quantity of field in the element \"unit\".",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Supplement).Quantity, nil
			}),
		},
		"unit": {
			Type:        unitTimeTypeEnum,
			Description: "Indicates the unit type. Possible unit types: Day or Hour.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Supplement).Unit, nil
			}),
		},
		"effectiveDate": {
			Type:        dateType,
			Description: "Indicates the effective date of the supplement.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Supplement).EffectiveDate, nil
			}),
		},
		"expireDate": {
			Type:        dateType,
			Description: "Indicates the expire date of the supplement.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Supplement).ExpireDate, nil
			}),
		},
		"resort": {
			Type:        resortType,
			Description: "Contains information about the resort",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return &source.(*domainHotelCommon.Supplement).Resort, nil
			}),
		},
		"price": {
			Type:        priceType,
			Description: "Indicates the supplement price.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return &source.(*domainHotelCommon.Supplement).Price, nil
			}),
		},
	},
}

var supplementTypeEnum = &graphql.EnumConfig{
	Name:        "SupplementType",
	Description: "Supplement Type",
	Values: graphql.EnumValueDefinitionMap{
		"SKI_PASS": {
			Description: "A ticket or pass authorizing the holder to ski in a certain place or resort.",
		},
		"LESSONS": {
			Description: "Lessons of any type that the costumer can take.",
		},
		"MEALS": {
			Description: "Supplement of a determined meal plan.",
		},
		"EQUIPMENT": {
			Description: "Extra equipment for a specific purpose.",
		},
		"TICKET": {
			Description: "Admission to some service.",
		},
		"TRANSFERS": {
			Description: "Transfers used by the costumer.",
		},
		"GALA": {
			Description: "Gala: A festive occasion, celebration or special entertainment.",
		},
		"ACTIVITY": {
			Description: " Activities that the costumer can do.",
		},
		"PERCENT": {
			Description: "Indicates the percentage applied by a rule.",
		},
		"IMPORT": {
			Description: "Indicates the exact amount applied by a rule.",
		},
	},
}

var chargeTypeEnum = &graphql.EnumConfig{
	Name:        "ChargeType",
	Description: "Charge Type",
	Values: graphql.EnumValueDefinitionMap{
		"INCLUDE": {
			Description: "The charge is included.",
		},
		"EXCLUDE": {
			Description: "The charge is excluded.",
		},
	},
}

var durationTypeEnum = &graphql.EnumConfig{
	Name:        "DurationType",
	Description: "Duration Type",
	Values: graphql.EnumValueDefinitionMap{
		"RANGE": {
			Description: "Date range is set.",
		},
		"OPEN": {
			Description: "Not restricted by date.",
		},
	},
}

var unitTimeTypeEnum = &graphql.EnumConfig{
	Name:        "UnitTimeType",
	Description: "Unit Time Type",
	Values: graphql.EnumValueDefinitionMap{
		"DAY": {
			Description: "Day",
		},
		"HOUR": {
			Description: "Hour",
		},
	},
}

var resortType = &graphql.ObjectConfig{
	Name:        "Resort",
	Description: "Contains information about the Resort.",
	Fields: graphql.Fields{
		"code": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Specifies the resort code.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Resort).Code, nil
			}),
		},
		"name": {
			Type:        graphql.T(graphql.String()),
			Description: "Specifies the resort name.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Resort).Name, nil
			}),
		},
		"description": {
			Type:        graphql.T(graphql.String()),
			Description: "Specifies the resort description.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Resort).Description, nil
			}),
		},
	},
}

var surchargeType = &graphql.ObjectConfig{
	Name:        "Surcharge",
	Description: "Surcharge that it can be or it is already added to the option returned. Contains all the information about the surcharge.",
	Fields: graphql.Fields{
		"chargeType": {
			Type: graphql.NonNullOf(chargeTypeEnum),
			Description: `Indicates the charge types. We need to know whether the supplements have to be paid when the consumer gets to the hotel or beforehand.
Possible charge types: Include or Exclude.
when include: this supplement is mandatory and included in the option's price
when exclude: this supplement is not included in the option's price`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Surcharge).ChargeType, nil
			}),
		},
		"mandatory": {
			Type: graphql.NonNullOfType(graphql.Boolean()),
			Description: `Indicates if the supplement is mandatory or not. If mandatory, this supplement will be applied to this option
if the chargeType is excluded the customer will have to pay it directly at the hotel`,
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Surcharge).Mandatory, nil
			}),
		},
		"price": {
			Type:        graphql.NonNullOf(priceType),
			Description: "Indicates the surcharge price.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return &source.(*domainHotelCommon.Surcharge).Price, nil
			}),
		},
		"description": {
			Type:        graphql.T(graphql.String()),
			Description: "Specifies the surcharge description.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.Surcharge).Description, nil
			}),
		},
	},
}

var rateRulesTypeEnum = &graphql.EnumConfig{
	Name:        "RateRulesType",
	Description: "Rate Rules",
	Values: graphql.EnumValueDefinitionMap{
		"PACKAGE": {
			Description: "The product can't be sold separately from another product attached to it, such as a flight.",
		},
		"OLDER55": {
			Description: "Options that can only be sold to people who are 55 and older.",
		},
		"OLDER60": {
			Description: "Options that can only be sold to people who are 60 and older.",
		},
		"OLDER65": {
			Description: "Options that can only be sold to people who are 65 and older.",
		},
		"CANARY_RESIDENT": {
			Description: "The rate CanaryResident is applicable to Canary Islands residents only.",
		},
		"BALEARIC_RESIDENT": {
			Description: "The rate BalearicResident is applicable to Balearic Islands residents only.",
		},
		"LARGE_FAMILY": {
			Description: "The rate largeFamily is applied to large families and is determined by each supplier",
		},
		"HONEYMOON": {
			Description: "The rate honeymoon is applied to those who just got married and is determined by each supplier.",
		},
		"PUBLIC_SERVANT": {
			Description: "The rate publicServant is applicable to public servants only.",
		},
		"UNEMPLOYED": {
			Description: "The rate unemployed is applied to those without work.",
		},
		"NORMAL": {
			Description: "The rate normal refers to options without RateRule",
		},
		"NON_REFUNDABLE": {
			Description: "The rate non refundable is applied to non refundable options",
		},
	},
}

var cancelPolicyType = &graphql.ObjectConfig{
	Name:        "CancelPolicy",
	Description: "Information about a policy cancellation.",
	Fields: graphql.Fields{
		"refundable": {
			Type:        graphql.NonNullOfType(graphql.Boolean()),
			Description: "Indicates if the option is refundable or non-refundable",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.CancelPolicy).Refundable, nil
			}),
		},
		"cancelPenalties": {
			Type:        graphql.ListOf(graphql.NonNullOf(cancelPenaltyType)),
			Description: "Specifies the resort name.",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				cancelPenalties := source.(*domainHotelCommon.CancelPolicy).CancelPenalties
				result := make([]*domainHotelCommon.CancelPenalty, len(cancelPenalties))
				for i := range cancelPenalties {
					result[i] = &cancelPenalties[i]
				}
				return result, nil
			}),
		},
	},
}

var cancelPenaltyType = &graphql.ObjectConfig{
	Name:        "CancelPenalty",
	Description: "Contains information for cancellation penalities.",
	Fields: graphql.Fields{
		"hoursBefore": {
			Type:        graphql.NonNullOfType(graphql.Int()),
			Description: "Cancellation fees applicable X number of hours before the check-in date",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.CancelPenalty).HoursBefore, nil
			}),
		},
		"penaltyType": {
			Type:        graphql.NonNullOf(cancelPenaltyTypeEnum),
			Description: "Type of penalty; this can be Nights, Percent or Import",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.CancelPenalty).Type, nil
			}),
		},
		"currency": {
			Type:        graphql.NonNullOf(currencyType),
			Description: "Currency used in the cancellation policy",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.CancelPenalty).Currency, nil
			}),
		},
		"value": {
			Type:        graphql.NonNullOfType(graphql.Float()),
			Description: "Currency used in the cancellation policy",
			Resolver: graphql.FieldResolverFunc(func(ctx context.Context, source interface{}, info graphql.ResolveInfo) (interface{}, error) {
				return source.(*domainHotelCommon.CancelPenalty).Value, nil
			}),
		},
	},
}
var cancelPenaltyTypeEnum = &graphql.EnumConfig{
	Name:        "CancelPenaltyType",
	Description: "Options type",
	Values: graphql.EnumValueDefinitionMap{
		"NIGHTS": {
			Description: "Indicates the number of nights to be penalized.",
		},
		"PERCENT": {
			Description: "Indicates the percentage to pay based on the option price.",
		},
		"IMPORT": {
			Description: "Indicates the exact amount payable.",
		},
	},
}

var addOnsType = &graphql.ObjectConfig{
	Name:        "AddOns",
	Description: "Additional information about the option",
	Fields: graphql.Fields{
		"distribute": {
			Type:        jsonType,
			Description: "Extra information from the distribution layer",
			Deprecation: &graphql.Deprecation{
				Reason: "deprecated from 2018-05-21. You can find it in distribution AddOn",
			},
		},
		"distribution": {
			Type:        graphql.ListOf(graphql.NonNullOf(addOnType)),
			Description: "Extra information from the distribution layer",
		},
	},
}

var addOnType = &graphql.ObjectConfig{
	Name:        "AddOn",
	Description: "Additional information about the option",
	Fields: graphql.Fields{
		"key": {
			Type:        graphql.NonNullOfType(graphql.String()),
			Description: "Contains keyword/ID to identify the AddOn.",
		},
		"value": {
			Type:        graphql.NonNullOf(jsonType),
			Description: "Contains AddOn values.",
		},
	},
}
