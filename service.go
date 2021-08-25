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
	return &service{}
}

func (s *service) seatPeople(rankNo int, groups []int) {
	s.rank.Number = rankNo
	sort.Ints(groups)
	s.creatRow(groups)
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
