package main

type Layout struct {
	Sections []Section `json:"sections"`
}

type Section struct {
	Name  string `json:"sectionname"`
	Ranks []Rank `json:"ranks"`
}

type Rank struct {
	Number int     `json:"ranknumber"`
	Rows   [][]int `json:"rows"`
}

type Resposne struct {
	Layout `json:"layout"`
}
