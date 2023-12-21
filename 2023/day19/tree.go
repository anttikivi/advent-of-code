package main

type interval struct {
	Min, Max int
}

type object map[byte]int

type constraint map[byte]interval

type expr interface {
	isExpr()
	apply(obj object) int
	propagate(constr constraint) int
}

type cond interface {
	isCond()
	apply(obj object) bool
	propagate(constr constraint) (constraint, constraint)
}

type ifThenElse struct {
	cond  cond
	then  expr
	else_ expr
}

type accepted struct{}
type rejected struct{}

type lt struct {
	category byte
	value    int
}
type gt struct {
	category byte
	value    int
}

func (i interval) Len() int {
	return i.Max - i.Min + 1
}

func Empty() interval {
	return interval{0, -1}
}

func (i interval) Intersection(j interval) interval {
	if i.Max < j.Min || j.Max < i.Min {
		return Empty()
	}
	a := max(i.Min, j.Min)
	b := min(i.Max, j.Max)
	return interval{a, b}
}

func (_ ifThenElse) isExpr() {}
func (_ accepted) isExpr()   {}
func (_ rejected) isExpr()   {}
func (_ lt) isCond()         {}
func (_ gt) isCond()         {}

func (c lt) apply(obj object) bool {
	return obj[c.category] < c.value
}
func (c gt) apply(obj object) bool {
	return obj[c.category] > c.value
}
func (e accepted) apply(obj object) int {
	return obj['x'] + obj['m'] + obj['a'] + obj['s']
}
func (e rejected) apply(obj object) int {
	return 0
}

func (e ifThenElse) apply(obj object) int {
	if e.cond.apply(obj) {
		return e.then.apply(obj)
	} else {
		return e.else_.apply(obj)
	}
}

func (c constraint) copy() constraint {
	return constraint{'x': c['x'], 'm': c['m'], 'a': c['a'], 's': c['s']}
}

func (c lt) propagate(constr constraint) (constraint, constraint) {
	pos := constr.copy()
	neg := constr.copy()
	name := c.category
	pos[name] = constr[name].Intersection(interval{1, c.value - 1})
	neg[name] = constr[name].Intersection(interval{c.value, 4000})
	return pos, neg
}

func (c gt) propagate(constr constraint) (constraint, constraint) {
	pos := constr.copy()
	neg := constr.copy()
	name := c.category
	pos[name] = constr[name].Intersection(interval{c.value + 1, 4000})
	neg[name] = constr[name].Intersection(interval{1, c.value})
	return pos, neg
}

func (e accepted) propagate(constr constraint) int {
	res := 1
	for _, v := range constr {
		res *= v.Len()
	}
	return res
}

func (e rejected) propagate(constr constraint) int {
	return 0
}

func (e ifThenElse) propagate(constr constraint) int {
	pos, neg := e.cond.propagate(constr)
	return e.then.propagate(pos) + e.else_.propagate(neg)
}
