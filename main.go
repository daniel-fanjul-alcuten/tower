package main

import (
	"fmt"
)

type ShapeId string

type Shape struct {
	id     ShapeId
	amount int
}

type StyleName string

type Style struct {
	name   StyleName
	shapes []Shape
}

type Selection map[StyleName]Shape

func (s Selection) Amount() (t int) {
	for _, v := range s {
		t += v.amount
	}
	return
}

type Group map[ShapeId]int

func (g Group) Amount() (t int) {
	for _, v := range g {
		t += v
	}
	return
}

func main() {

	f := Style{"f", []Shape{Shape{"h", 200}, Shape{"d", 800}, Shape{"v", 1600}, Shape{"s", 3200}}}
	s := Style{"s", []Shape{Shape{"d", 400}, Shape{"v", 800}, Shape{"h", 2400}, Shape{"s", 4800}}}
	e := Style{"e", []Shape{Shape{"v", 1600}, Shape{"d", 3200}, Shape{"h", 4000}}}
	r := Style{"r", []Shape{Shape{"s", 2000}, Shape{"d", 4000}}}
	d := Style{"d", []Shape{Shape{"d", 3200}, Shape{"v", 6000}}}

	f.shapes = f.shapes[1:]
	s.shapes = s.shapes[1:]

	for _, max := range []Group{
		Group{"h": 1, "d": 1},
		Group{"h": 1, "d": 1, "v": 1, "s": 1},
		Group{"h": 1, "d": 1, "v": 2, "s": 1},
	} {
		search(max, f, s, e, r, d)
	}
}

func search(max Group, styles ...Style) {
	var best Selection
	searchRec(styles, Selection{}, Group{}, max, &best)
	fmt.Printf("%v: %v (%v)\n", max, best, best.Amount())
}

func searchRec(styles []Style, selection Selection, group, max Group, best *Selection) {
	if group.Amount() == max.Amount() {
		selectionAmount := selection.Amount()
		bestAmount := best.Amount()
		if selectionAmount > bestAmount {
			tmp := Selection{}
			for k, v := range selection {
				tmp[k] = v
			}
			*best = tmp
		}
		return
	}
	if len(styles) == 0 {
		return
	}
	tail := styles[1:]
	searchRec(tail, selection, group, max, best)
	style := styles[0]
	name := style.name
	for _, shape := range style.shapes {
		id := shape.id
		if group[id] < max[id] {
			selection[name] = shape
			group[id] += 1
			searchRec(tail, selection, group, max, best)
			group[id] -= 1
		}
	}
	delete(selection, name)
}
