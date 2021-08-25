package main

import (
	"sort"
)

type service struct {
	rank        Rank
	Layout      [][]int
	sumOfGroups int
	row         []int
}

func newService() *service {
	return &service{}
}

func (s *service) seatPeople(rankNo int, groups []int) {
	s.rank.Number = rankNo
	sort.Ints(groups)
	s.creatRow(groups)

}

func (s *service) creatRow(groups []int) {
	testint := 0
	for i := 0; i < len(groups); i++ {

		if s.sumOfGroups+groups[i] > cap(s.Layout[i]) {
			s.Layout[testint] = s.row
			testint++
			s.sumOfGroups = 0
			s.row = nil
		}
		if s.sumOfGroups < cap(s.Layout[i]) && s.sumOfGroups+groups[i] <= cap(s.Layout[i]) {
			g := s.creatGroup(groups[i])
			s.row = append(s.row, g...)
			s.sumOfGroups += groups[i]
		}

	}
}

func (s *service) creatGroup(group int) []int {
	var groupView []int
	for i := 0; i < group; i++ {

		groupView = append(groupView, group)

	}
	return groupView
}
