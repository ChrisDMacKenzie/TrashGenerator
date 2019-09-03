package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

var numbers = []string{"x", "y", "t", "n"}
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

type operator interface {
	setDefaults()
	setSecondaryOps()
	setValues()
	compute(x, y, t float64) float64
	print() string
}

// TotalOperations is the total number of distinct operation types
const TotalOperations = 10

func getRandomOperators(n int) []operator {
	var seq []operator
	for i := 0; i < n; i++ {
		idx := r.Intn(100) % TotalOperations
		switch idx {
		case 0:
			op := new(cosine)
			op.setDefaults()
			seq = append(seq, op)
		case 1:
			op := new(sine)
			op.setDefaults()
			seq = append(seq, op)
		case 2:
			op := new(subtract)
			op.setDefaults()
			seq = append(seq, op)
		case 3:
			op := new(multiply)
			op.setDefaults()
			seq = append(seq, op)
		case 4:
			op := new(absolute)
			op.setDefaults()
			seq = append(seq, op)
		case 5:
			op := new(squareRoot)
			op.setDefaults()
			seq = append(seq, op)
		case 6:
			op := new(floor)
			op.setDefaults()
			seq = append(seq, op)
		case 7:
			op := new(max)
			op.setDefaults()
			seq = append(seq, op)
		case 8:
			op := new(min)
			op.setDefaults()
			seq = append(seq, op)
		case 9:
			op := new(atan2)
			op.setDefaults()
			seq = append(seq, op)
		}
	}

	return seq
}

func getRandomValues(allVals []string, n int) []string {
	var selected []string
	for i := 0; i < n; i++ {
		selected = append(selected, allVals[r.Intn(100)%len(allVals)])
	}
	return selected
}

func setComputeVals(names []string, x, y, t, n float64) []float64 {
	var values []float64
	for _, name := range names {
		switch name {
		case "x":
			values = append(values, x)
		case "y":
			values = append(values, y)
		case "t":
			values = append(values, t)
		case "n":
			values = append(values, n)
		}
	}
	return values
}

func generateRandomFloat() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return float64(r.Intn(10)) + r.Float64()
}

type cosine struct {
	name         string
	values       []string
	numOps       int
	secondaryOps []operator
	rand         float64
	primary      *bool
}

func (c *cosine) setDefaults() {
	t := true
	c.primary = &t
	c.numOps = 1
	c.values = *(new([]string))
	c.secondaryOps = *(new([]operator))
	c.rand = 0
	c.name = "COSINE"
}

func (c *cosine) setValues() {
	f := false
	c.primary = &f
	rvals := getRandomValues(numbers, 1)
	copyvals := make([]string, len(rvals))
	copy(copyvals, rvals)
	c.values = copyvals
	c.rand = generateRandomFloat()
}

func (c *cosine) setSecondaryOps() {
	c.secondaryOps = getRandomOperators(c.numOps)
	for _, op := range c.secondaryOps {
		// op.setDefaults()
		op.setValues()
	}
}

func (c *cosine) compute(x, y, t float64) float64 {
	if *c.primary == true {
		v := c.secondaryOps[0].compute(x, y, t)
		return math.Cos(v)
	}
	vals := setComputeVals(c.values, x, y, t, c.rand)
	return math.Cos(vals[0])
}

func (c *cosine) print() string {
	// fmt.Println(*c.primary, "cosine")
	var e string
	if *c.primary == true {
		inner := c.secondaryOps[0].print()
		e = fmt.Sprintf("cos(%s)", inner)
	} else {
		switch c.values[0] {
		case "n":
			e = fmt.Sprintf("cos(%f)", c.rand)
		default:
			e = fmt.Sprintf("cos(%s)", c.values[0])
		}
	}
	return e
}

type sine struct {
	name         string
	values       []string
	numOps       int
	secondaryOps []operator
	rand         float64
	primary      *bool
}

func (s *sine) setDefaults() {
	t := true
	s.primary = &t
	s.numOps = 1
	s.values = *(new([]string))
	s.secondaryOps = *(new([]operator))
	s.rand = 0
	s.name = "SINE"
}

func (s *sine) setValues() {
	f := false
	s.primary = &f
	rvals := getRandomValues(numbers, 1)
	copyvals := make([]string, len(rvals))
	copy(copyvals, rvals)
	s.values = copyvals
	s.rand = generateRandomFloat()
}

func (s *sine) setSecondaryOps() {
	s.secondaryOps = getRandomOperators(s.numOps)
	for _, op := range s.secondaryOps {
		// op.setDefaults()
		op.setValues()
	}
}

func (s *sine) compute(x, y, t float64) float64 {
	if *s.primary == true {
		v := s.secondaryOps[0].compute(x, y, t)
		return math.Sin(v)
	}
	vals := setComputeVals(s.values, x, y, t, s.rand)
	return math.Sin(vals[0])
}

func (s *sine) print() string {
	// fmt.Println(*s.primary, "sine")
	var e string
	if *s.primary == true {
		inner := s.secondaryOps[0].print()
		e = fmt.Sprintf("sin(%s)", inner)
	} else {
		switch s.values[0] {
		case "n":
			e = fmt.Sprintf("cos(%f)", s.rand)
		default:
			e = fmt.Sprintf("cos(%s)", s.values[0])
		}
	}
	return e
}

type subtract struct {
	name         string
	values       []string
	numOps       int
	secondaryOps []operator
	rand         float64
	primary      *bool
}

func (s *subtract) setDefaults() {
	t := true
	s.primary = &t
	s.numOps = 2
	s.values = *(new([]string))
	s.secondaryOps = *(new([]operator))
	s.rand = 0
	s.name = "SUBTRACT"
}

func (s *subtract) setValues() {
	f := false
	s.primary = &f
	rvals := getRandomValues(numbers, 2)
	copyvals := make([]string, len(rvals))
	copy(copyvals, rvals)
	s.values = copyvals
	s.rand = generateRandomFloat()
}

func (s *subtract) setSecondaryOps() {
	s.secondaryOps = getRandomOperators(s.numOps)
	for _, op := range s.secondaryOps {
		op.setDefaults()
		// fmt.Printf("%+v\n", op)
		op.setValues()
	}
}

func (s *subtract) compute(x, y, t float64) float64 {
	if *s.primary == true {
		v1 := s.secondaryOps[0].compute(x, y, t)
		v2 := s.secondaryOps[1].compute(x, y, t)
		return v1 - v2
	}
	vals := setComputeVals(s.values, x, y, t, s.rand)
	return vals[0] - vals[1]
}

func (s *subtract) print() string {
	// fmt.Println(*s.primary, "subtract")
	var e []string
	if *s.primary == true {
		for _, op := range s.secondaryOps {
			e = append(e, op.print())
		}
	} else {
		for _, v := range s.values {
			switch v {
			case "n":
				e = append(e, fmt.Sprintf("%f", s.rand))
			default:
				e = append(e, fmt.Sprintf("%s", v))
			}

		}
	}
	return "(" + strings.Join(e, "-") + ")"
}

type multiply struct {
	name         string
	values       []string
	numOps       int
	secondaryOps []operator
	rand         float64
	primary      *bool
}

func (m *multiply) setDefaults() {
	t := true
	m.primary = &t
	m.numOps = 2
	m.values = *(new([]string))
	m.secondaryOps = *(new([]operator))
	m.rand = 0
	m.name = "MULTIPLY"
}

func (m *multiply) setValues() {
	f := false
	m.primary = &f
	rvals := getRandomValues(numbers, 2)
	copyvals := make([]string, len(rvals))
	copy(copyvals, rvals)
	m.values = copyvals
	m.rand = generateRandomFloat()
}

func (m *multiply) setSecondaryOps() {
	m.secondaryOps = getRandomOperators(m.numOps)
	for _, op := range m.secondaryOps {
		// op.setDefaults()
		// fmt.Printf("%+v\n", op)
		op.setValues()
	}
}

func (m *multiply) compute(x, y, t float64) float64 {
	if *m.primary == true {
		v1 := m.secondaryOps[0].compute(x, y, t)
		v2 := m.secondaryOps[1].compute(x, y, t)
		return v1 * v2
	}
	vals := setComputeVals(m.values, x, y, t, m.rand)
	return vals[0] * vals[1]
}

func (m *multiply) print() string {
	// fmt.Println(*m.primary, "multiply")
	var e []string
	if *m.primary == true {
		for _, op := range m.secondaryOps {
			e = append(e, op.print())
		}
	} else {
		for _, v := range m.values {
			switch v {
			case "n":
				e = append(e, fmt.Sprintf("%f", m.rand))
			default:
				e = append(e, fmt.Sprintf("%s", v))
			}
		}
	}
	return "(" + strings.Join(e, "*") + ")"
}

type absolute struct {
	name         string
	values       []string
	numOps       int
	secondaryOps []operator
	rand         float64
	primary      *bool
}

func (a *absolute) setDefaults() {
	t := true
	a.primary = &t
	a.numOps = 1
	a.values = *(new([]string))
	a.secondaryOps = *(new([]operator))
	a.rand = 0
	a.name = "ABSOLUTE"
}

func (a *absolute) setValues() {
	f := false
	a.primary = &f
	rvals := getRandomValues(numbers, 1)
	copyvals := make([]string, len(rvals))
	copy(copyvals, rvals)
	a.values = copyvals
	a.rand = generateRandomFloat()
}

func (a *absolute) setSecondaryOps() {
	a.secondaryOps = getRandomOperators(a.numOps)
	for _, op := range a.secondaryOps {
		// op.setDefaults()
		// fmt.Printf("%+v\n", op)
		op.setValues()
	}
}

func (a *absolute) compute(x, y, t float64) float64 {
	if *a.primary == true {
		v := a.secondaryOps[0].compute(x, y, t)
		return math.Abs(v)
	}
	vals := setComputeVals(a.values, x, y, t, a.rand)
	return math.Abs(vals[0])
}

func (a *absolute) print() string {
	// fmt.Println(*a.primary, "absolute")
	var e string
	if *a.primary == true {
		inner := a.secondaryOps[0].print()
		e = fmt.Sprintf("abs(%s)", inner)
	} else {
		switch a.values[0] {
		case "n":
			e = fmt.Sprintf("abs(%f)", a.rand)
		default:
			e = fmt.Sprintf("abs(%s)", a.values[0])
		}
	}
	return e
}

type squareRoot struct {
	name         string
	values       []string
	numOps       int
	secondaryOps []operator
	rand         float64
	primary      *bool
}

func (s *squareRoot) setDefaults() {
	t := true
	s.primary = &t
	s.numOps = 1
	s.values = *(new([]string))
	s.secondaryOps = *(new([]operator))
	s.rand = 0
	s.name = "SQRT"
}

func (s *squareRoot) setValues() {
	f := false
	s.primary = &f
	rvals := getRandomValues(numbers, 1)
	copyvals := make([]string, len(rvals))
	copy(copyvals, rvals)
	s.values = copyvals
	s.rand = generateRandomFloat()
}

func (s *squareRoot) setSecondaryOps() {
	s.secondaryOps = getRandomOperators(s.numOps)
	for _, op := range s.secondaryOps {
		// op.setDefaults()
		// fmt.Printf("%+v\n", op)
		op.setValues()
	}
}

func (s *squareRoot) compute(x, y, t float64) float64 {
	if *s.primary == true {
		v := s.secondaryOps[0].compute(x, y, t)
		return math.Sqrt(v)
	}
	vals := setComputeVals(s.values, x, y, t, s.rand)
	return math.Sqrt(vals[0])
}

func (s *squareRoot) print() string {
	// fmt.Println(*s.primary, "sqrt")
	var e string
	if *s.primary == true {
		inner := s.secondaryOps[0].print()
		e = fmt.Sprintf("sqrt(%s)", inner)
	} else {
		switch s.values[0] {
		case "n":
			e = fmt.Sprintf("sqrt(%f)", s.rand)
		default:
			e = fmt.Sprintf("sqrt(%s)", s.values[0])
		}
	}
	return e
}

type floor struct {
	name         string
	values       []string
	numOps       int
	secondaryOps []operator
	rand         float64
	primary      *bool
}

func (f *floor) setDefaults() {
	t := true
	f.primary = &t
	f.numOps = 1
	f.values = *(new([]string))
	f.secondaryOps = *(new([]operator))
	f.rand = 0
	f.name = "FLOOR"
}

func (f *floor) setValues() {
	fa := false
	f.primary = &fa
	rvals := getRandomValues(numbers, 1)
	copyvals := make([]string, len(rvals))
	copy(copyvals, rvals)
	f.values = copyvals
	f.rand = generateRandomFloat()
}

func (f *floor) setSecondaryOps() {
	f.secondaryOps = getRandomOperators(f.numOps)
	for _, op := range f.secondaryOps {
		// op.setDefaults()
		// fmt.Printf("%+v\n", op)
		op.setValues()
	}
}

func (f *floor) compute(x, y, t float64) float64 {
	if *f.primary == true {
		v := f.secondaryOps[0].compute(x, y, t)
		return math.Floor(v)
	}
	vals := setComputeVals(f.values, x, y, t, f.rand)
	return math.Floor(vals[0])
}

func (f *floor) print() string {
	// fmt.Println(*f.primary, "floor")
	var e string
	if *f.primary == true {
		inner := f.secondaryOps[0].print()
		e = fmt.Sprintf("floor(%s)", inner)
	} else {
		switch f.values[0] {
		case "n":
			e = fmt.Sprintf("floor(%f)", f.rand)
		default:
			e = fmt.Sprintf("floor(%s)", f.values[0])
		}
	}
	return e
}

type max struct {
	name         string
	values       []string
	numOps       int
	secondaryOps []operator
	rand         float64
	primary      *bool
}

func (m *max) setDefaults() {
	t := true
	m.primary = &t
	m.numOps = 2
	m.values = *(new([]string))
	m.secondaryOps = *(new([]operator))
	m.rand = 0
	m.name = "MAX"
}

func (m *max) setValues() {
	f := false
	m.primary = &f
	rvals := getRandomValues(numbers, 2)
	copyvals := make([]string, len(rvals))
	copy(copyvals, rvals)
	m.values = copyvals
	m.rand = generateRandomFloat()
}

func (m *max) setSecondaryOps() {
	m.secondaryOps = getRandomOperators(m.numOps)
	for _, op := range m.secondaryOps {
		// op.setDefaults()
		// fmt.Printf("%+v\n", op)
		op.setValues()
	}
}

func (m *max) compute(x, y, t float64) float64 {
	if *m.primary == true {
		v1 := m.secondaryOps[0].compute(x, y, t)
		v2 := m.secondaryOps[1].compute(x, y, t)
		return math.Max(v1, v2)
	}
	vals := setComputeVals(m.values, x, y, t, m.rand)
	return math.Max(vals[0], vals[1])
}

func (m *max) print() string {
	// fmt.Println(*m.primary, "max")
	var e []string
	if *m.primary == true {
		for _, op := range m.secondaryOps {
			e = append(e, op.print())
		}
	} else {
		for _, v := range m.values {
			switch v {
			case "n":
				e = append(e, fmt.Sprintf("%f", m.rand))
			default:
				e = append(e, fmt.Sprintf("%s", v))
			}

		}
	}
	return "max(" + strings.Join(e, ",") + ")"
}

type min struct {
	name         string
	values       []string
	numOps       int
	secondaryOps []operator
	rand         float64
	primary      *bool
}

func (m *min) setDefaults() {
	t := true
	m.primary = &t
	m.numOps = 2
	m.values = *(new([]string))
	m.secondaryOps = *(new([]operator))
	m.rand = 0
	m.name = "MIN"
}

func (m *min) setValues() {
	f := false
	m.primary = &f
	rvals := getRandomValues(numbers, 2)
	copyvals := make([]string, len(rvals))
	copy(copyvals, rvals)
	m.values = copyvals
	m.rand = generateRandomFloat()
}

func (m *min) setSecondaryOps() {
	m.secondaryOps = getRandomOperators(m.numOps)
	for _, op := range m.secondaryOps {
		// op.setDefaults()
		// fmt.Printf("%+v\n", op)
		op.setValues()
	}
}

func (m *min) compute(x, y, t float64) float64 {
	if *m.primary == true {
		v1 := m.secondaryOps[0].compute(x, y, t)
		v2 := m.secondaryOps[1].compute(x, y, t)
		return math.Max(v1, v2)
	}
	vals := setComputeVals(m.values, x, y, t, m.rand)
	return math.Min(vals[0], vals[1])
}

func (m *min) print() string {
	// fmt.Println(*m.primary, "min")
	var e []string
	if *m.primary == true {
		for _, op := range m.secondaryOps {
			e = append(e, op.print())
		}
	} else {
		for _, v := range m.values {
			switch v {
			case "n":
				e = append(e, fmt.Sprintf("%f", m.rand))
			default:
				e = append(e, fmt.Sprintf("%s", v))
			}

		}
	}
	return "min(" + strings.Join(e, ",") + ")"
}

type atan2 struct {
	name         string
	values       []string
	numOps       int
	secondaryOps []operator
	rand         float64
	primary      *bool
}

func (a *atan2) setDefaults() {
	t := true
	a.primary = &t
	a.numOps = 2
	a.values = *(new([]string))
	a.secondaryOps = *(new([]operator))
	a.rand = 0
	a.name = "ATAN2"
}

func (a *atan2) setValues() {
	f := false
	a.primary = &f
	rvals := getRandomValues(numbers, 2)
	copyvals := make([]string, len(rvals))
	copy(copyvals, rvals)
	a.values = copyvals
	a.rand = generateRandomFloat()
}

func (a *atan2) setSecondaryOps() {
	a.secondaryOps = getRandomOperators(a.numOps)
	for _, op := range a.secondaryOps {
		// op.setDefaults()
		// fmt.Printf("%+v\n", op)
		op.setValues()
	}
}

func (a *atan2) compute(x, y, t float64) float64 {
	if *a.primary == true {
		v1 := a.secondaryOps[0].compute(x, y, t)
		v2 := a.secondaryOps[1].compute(x, y, t)
		return math.Atan2(v1, v2)
	}
	vals := setComputeVals(a.values, x, y, t, a.rand)
	return math.Min(vals[0], vals[1])
}

func (a *atan2) print() string {
	// fmt.Println(*a.primary, "atan2")
	var e []string
	if *a.primary == true {
		for _, op := range a.secondaryOps {
			e = append(e, op.print())
		}
	} else {
		for _, v := range a.values {
			switch v {
			case "n":
				e = append(e, fmt.Sprintf("%f", a.rand))
			default:
				e = append(e, fmt.Sprintf("%s", v))
			}

		}
	}
	return "atan2(" + strings.Join(e, ",") + ")"
}