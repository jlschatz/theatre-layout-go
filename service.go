package main

import (
	"fmt"
	"sort"
)

type service struct {
	rank        Rank
	Layout      [][]int
	sumOfGroups int
	usersInRow  []int
}

func newService() *service {
	srv := &service{}
	srv.Layout = make([][]int, 10)
	for i := range srv.Layout {
		srv.Layout[i] = make([]int, 0, 10)
	}
	return srv
}

func (s *service) seatPeople(rankNo int, groups []int) Layout {
	s.rank.Number = rankNo
	sort.Ints(groups)
	s.creatRow(groups)

	var testResponse Layout
	var testSection Section
	var testRank Rank

	testSection.Name = "hall"
	testRank.Number = s.rank.Number
	testRank.Rows = s.Layout

	testSection.Ranks = append(testSection.Ranks, testRank)
	testResponse.Sections = append(testResponse.Sections, testSection)

	return testResponse
}

func (s *service) creatRow(groups []int) {
	row := 0
	for i := 0; i < len(groups); i++ {
		if s.sumOfGroups+groups[i] > cap(s.Layout[i]) {
			s.Layout[row] = s.usersInRow
			row++
			s.sumOfGroups = 0
			s.usersInRow = nil
		}
		if s.sumOfGroups < cap(s.Layout[i]) && s.sumOfGroups+groups[i] <= cap(s.Layout[i]) {
			g := s.creatGroup(groups[i])
			s.usersInRow = append(s.usersInRow, g...)
			s.sumOfGroups += groups[i]
		}
		if i == len(groups)-1 {
			s.Layout[row] = s.usersInRow
			fmt.Println("All groups placed")
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
