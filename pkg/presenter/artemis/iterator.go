package artemis

import (
	"github.com/botobag/artemis/graphql/executor"
	"github.com/botobag/artemis/iterator"
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

// SearchOptionList implements executor.Iterable to loop over []*domainHotelCommon.Option as an
// list of *hotelOptionSearch
type SearchOptionList struct {
	options  []*domainHotelCommon.Option
	criteria string
}

// Iterator implements executor.Iterable.
func (list SearchOptionList) Iterator() executor.Iterator {
	return &SearchOptionListIter{
		list: list,
	}
}

// Size implements executor.SizedIterable.
func (list SearchOptionList) Size() int {
	return len(list.options)
}

// SearchOptionListIter implements executor.Iterator which defines iterator for an SearchOptionList.
type SearchOptionListIter struct {
	list    SearchOptionList
	nextIdx int
}

// Next implements executor.Iterator.
func (iter *SearchOptionListIter) Next() (interface{}, error) {
	var (
		list    = iter.list
		nextIdx = iter.nextIdx
	)
	if nextIdx >= list.Size() {
		return nil, iterator.Done
	}

	option := list.options[nextIdx]
	iter.nextIdx++
	return &hotelOptionSearch{
		option:   option,
		criteria: iter.list.criteria,
	}, nil
}

// OccupancyList implements executor.Iterable to loop over []domainHotelCommon.Occupancy as an list
// of *domainHotelCommon.Occupancy.
type OccupancyList []domainHotelCommon.Occupancy

// Iterator implements executor.Iterable.
func (list OccupancyList) Iterator() executor.Iterator {
	return &OccupancyListIter{
		list: list,
	}
}

// Size implements executor.SizedIterable.
func (list OccupancyList) Size() int {
	return len(list)
}

// OccupancyListIter implements executor.Iterator which defines iterator for an OccupancyList.
type OccupancyListIter struct {
	list    OccupancyList
	nextIdx int
}

// Next implements executor.Iterator.
func (iter *OccupancyListIter) Next() (interface{}, error) {
	var (
		list    = iter.list
		nextIdx = iter.nextIdx
	)
	if nextIdx >= list.Size() {
		return nil, iterator.Done
	}
	occupancy := &list[nextIdx]
	iter.nextIdx++
	return occupancy, nil
}

// RoomList implements executor.Iterable to loop over []domainHotelCommon.Room as an list of
// *domainHotelCommon.Room.
type RoomList []domainHotelCommon.Room

// Iterator implements executor.Iterable.
func (list RoomList) Iterator() executor.Iterator {
	return &RoomListIter{
		list: list,
	}
}

// Size implements executor.SizedIterable.
func (list RoomList) Size() int {
	return len(list)
}

// RoomListIter implements executor.Iterator which defines iterator for an RoomList.
type RoomListIter struct {
	list    RoomList
	nextIdx int
}

// Next implements executor.Iterator.
func (iter *RoomListIter) Next() (interface{}, error) {
	var (
		list    = iter.list
		nextIdx = iter.nextIdx
	)
	if nextIdx >= list.Size() {
		return nil, iterator.Done
	}
	room := &list[nextIdx]
	iter.nextIdx++
	return room, nil
}

// SurchargeList implements executor.Iterable to loop over []domainHotelCommon.Surcharge as an list
// of *domainHotelCommon.Surcharge.
type SurchargeList []domainHotelCommon.Surcharge

// Iterator implements executor.Iterable.
func (list SurchargeList) Iterator() executor.Iterator {
	return &SurchargeListIter{
		list: list,
	}
}

// Size implements executor.SizedIterable.
func (list SurchargeList) Size() int {
	return len(list)
}

// SurchargeListIter implements executor.Iterator which defines iterator for an SurchargeList.
type SurchargeListIter struct {
	list    SurchargeList
	nextIdx int
}

// Next implements executor.Iterator.
func (iter *SurchargeListIter) Next() (interface{}, error) {
	var (
		list    = iter.list
		nextIdx = iter.nextIdx
	)
	if nextIdx >= list.Size() {
		return nil, iterator.Done
	}
	surcharge := &list[nextIdx]
	iter.nextIdx++
	return surcharge, nil
}

// PaxList implements executor.Iterable to loop over []domainHotelCommon.Pax as an list of
// *domainHotelCommon.Pax.
type PaxList []domainHotelCommon.Pax

// Iterator implements executor.Iterable.
func (list PaxList) Iterator() executor.Iterator {
	return &PaxListIter{
		list: list,
	}
}

// Size implements executor.SizedIterable.
func (list PaxList) Size() int {
	return len(list)
}

// PaxListIter implements executor.Iterator which defines iterator for an PaxList.
type PaxListIter struct {
	list    PaxList
	nextIdx int
}

// Next implements executor.Iterator.
func (iter *PaxListIter) Next() (interface{}, error) {
	var (
		list    = iter.list
		nextIdx = iter.nextIdx
	)
	if nextIdx >= list.Size() {
		return nil, iterator.Done
	}
	pax := &list[nextIdx]
	iter.nextIdx++
	return pax, nil
}

// RuleList implements executor.Iterable to loop over []domainHotelCommon.Rule as an list
// of *domainHotelCommon.Rule.
type RuleList []domainHotelCommon.Rule

// Iterator implements executor.Iterable.
func (list RuleList) Iterator() executor.Iterator {
	return &RuleListIter{
		list: list,
	}
}

// Size implements executor.SizedIterable.
func (list RuleList) Size() int {
	return len(list)
}

// RuleListIter implements executor.Iterator which defines iterator for an RuleList.
type RuleListIter struct {
	list    RuleList
	nextIdx int
}

// Next implements executor.Iterator.
func (iter *RuleListIter) Next() (interface{}, error) {
	var (
		list    = iter.list
		nextIdx = iter.nextIdx
	)
	if nextIdx >= list.Size() {
		return nil, iterator.Done
	}
	rule := &list[nextIdx]
	iter.nextIdx++
	return rule, nil
}

// CancelPenaltyList implements executor.Iterable to loop over []domainHotelCommon.CancelPenalty as
// an list of *domainHotelCommon.CancelPenalty.
type CancelPenaltyList []domainHotelCommon.CancelPenalty

// Iterator implements executor.Iterable.
func (list CancelPenaltyList) Iterator() executor.Iterator {
	return &CancelPenaltyListIter{
		list: list,
	}
}

// Size implements executor.SizedIterable.
func (list CancelPenaltyList) Size() int {
	return len(list)
}

// CancelPenaltyListIter implements executor.Iterator which defines iterator for an CancelPenaltyList.
type CancelPenaltyListIter struct {
	list    CancelPenaltyList
	nextIdx int
}

// Next implements executor.Iterator.
func (iter *CancelPenaltyListIter) Next() (interface{}, error) {
	var (
		list    = iter.list
		nextIdx = iter.nextIdx
	)
	if nextIdx >= list.Size() {
		return nil, iterator.Done
	}
	cancelPenalty := &list[nextIdx]
	iter.nextIdx++
	return cancelPenalty, nil
}
