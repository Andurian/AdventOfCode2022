package main

import "fmt"

type Fraction struct {
	numerator, denominator int
}

func (f Fraction) String() string {
	return fmt.Sprintf("(%d/%d)", f.numerator, f.denominator)
}

func simplify(f Fraction) Fraction {
	if f.numerator%f.denominator == 0 {
		return Fraction{f.numerator / f.denominator, 1}
	}
	return f
}

func TryAsInt(f Fraction) int {
	f = simplify(f)
	if f.denominator == 1 {
		return f.numerator
	}
	panic("Cannot simplify")
}

func Add(left, right Fraction) Fraction {
	numerator := left.numerator*right.denominator + right.numerator*left.denominator
	denominator := left.denominator * right.denominator
	return simplify(Fraction{numerator, denominator})
}

func Subtract(left, right Fraction) Fraction {
	numerator := left.numerator*right.denominator - right.numerator*left.denominator
	denominator := left.denominator * right.denominator
	return simplify(Fraction{numerator, denominator})
}

func Multiply(left, right Fraction) Fraction {
	numerator := left.numerator * right.numerator
	denominator := left.denominator * right.denominator
	return simplify(Fraction{numerator, denominator})
}

func Divide(left, right Fraction) Fraction {
	numerator := left.numerator * right.denominator
	denominator := left.denominator * right.numerator
	return simplify(Fraction{numerator, denominator})
}

type Term interface {
	Degree() int
	Get(degree int) Fraction

	Add(Term) Term
	SubtractRight(Term) Term
	SubtractLeft(Term) Term
	Multiply(Term) Term
	DivideRight(Term) Term
	DivideLeft(Term) Term
}

type TermDeg0 struct {
	value Fraction
}

func NewTermDeg0FromInt(v int) *TermDeg0 {
	return &TermDeg0{Fraction{v, 1}}
}

func (t *TermDeg0) String() string {
	return fmt.Sprintf("%v", t.value)
}

func (t *TermDeg0) Degree() int { return 0 }
func (t *TermDeg0) Get(degree int) Fraction {
	if degree != 0 {
		panic("Invalid degree")
	}
	return t.value
}

func (t *TermDeg0) Add(other Term) Term {
	if other.Degree() > t.Degree() {
		return other.Add(t)
	}
	return &TermDeg0{Add(t.value, other.Get(0))}
}

func (t *TermDeg0) SubtractRight(other Term) Term {
	if other.Degree() > t.Degree() {
		return other.SubtractLeft(t)
	}
	return &TermDeg0{Subtract(t.value, other.Get(0))}
}

func (t *TermDeg0) SubtractLeft(other Term) Term {
	if other.Degree() > t.Degree() {
		return other.SubtractRight(t)
	}
	return &TermDeg0{Subtract(other.Get(0), t.value)}
}

func (t *TermDeg0) Multiply(other Term) Term {
	if other.Degree() > t.Degree() {
		return other.Multiply(t)
	}
	return &TermDeg0{Multiply(t.value, other.Get(0))}
}

func (t *TermDeg0) DivideRight(other Term) Term {
	if other.Degree() > t.Degree() {
		return other.DivideLeft(t)
	}
	return &TermDeg0{Divide(t.value, other.Get(0))}
}

func (t *TermDeg0) DivideLeft(other Term) Term {
	if other.Degree() > t.Degree() {
		return other.DivideRight(t)
	}
	return &TermDeg0{Divide(other.Get(0), t.value)}
}

type TermDeg1 struct {
	valueX, value Fraction
}

func NewTermDeg1FromInt(vX, v int) *TermDeg1 {
	return &TermDeg1{Fraction{vX, 1}, Fraction{v, 1}}
}

func (t *TermDeg1) String() string {
	return fmt.Sprintf("[%v*x + %v]", t.valueX, t.value)
}

func (t *TermDeg1) Degree() int { return 1 }
func (t *TermDeg1) Get(degree int) Fraction {
	if degree != 0 {
		panic("Invalid degree")
	}
	return t.value
}

func (t *TermDeg1) Add(other Term) Term {
	if other.Degree() > t.Degree() {
		panic("degrees > 1 not supported")
	}
	if other.Degree() == 1 {
		return &TermDeg1{Add(t.valueX, other.Get(1)), Add(t.value, other.Get(0))}
	}
	return &TermDeg1{t.valueX, Add(t.value, other.Get(0))}
}

func (t *TermDeg1) SubtractRight(other Term) Term {
	if other.Degree() > t.Degree() {
		panic("degrees > 1 not supported")
	}
	if other.Degree() == 1 {
		return &TermDeg1{Subtract(t.valueX, other.Get(1)), Subtract(t.value, other.Get(0))}
	}
	return &TermDeg1{t.valueX, Subtract(t.value, other.Get(0))}
}

func (t *TermDeg1) SubtractLeft(other Term) Term {
	if other.Degree() > t.Degree() {
		panic("degrees > 1 not supported")
	}
	if other.Degree() == 1 {
		return &TermDeg1{Subtract(other.Get(1), t.valueX), Subtract(other.Get(0), t.value)}
	}
	return &TermDeg1{Subtract(Fraction{0, 1}, t.valueX), Subtract(other.Get(0), t.value)}
}

func (t *TermDeg1) Multiply(other Term) Term {
	if other.Degree() > t.Degree() {
		panic("degrees > 1 not supported")
	}
	if other.Degree() == 1 {
		panic("Cannot mulitply up to degrees > 1")
	}
	return &TermDeg1{Multiply(t.valueX, other.Get(0)), Multiply(t.value, other.Get(0))}
}

func (t *TermDeg1) DivideRight(other Term) Term {
	if other.Degree() > t.Degree() {
		panic("degrees > 1 not supported")
	}
	if other.Degree() == 1 {
		panic("Cannot divide down to deg 0")
	}
	return &TermDeg1{Divide(t.valueX, other.Get(0)), Divide(t.value, other.Get(0))}
}

func (t *TermDeg1) DivideLeft(other Term) Term {
	if other.Degree() > t.Degree() {
		panic("degrees > 1 not supported")
	}
	panic("Should not happen")
}

func AddTerm(left, right Term) Term {
	return left.Add(right)
}

func SubtractTerm(left, right Term) Term {
	return left.SubtractRight(right)
}

func MultiplyTerm(left, right Term) Term {
	return left.Multiply(right)
}

func DivideTerm(left, right Term) Term {
	return left.DivideRight(right)
}

func Solve(t *TermDeg1) Fraction {
	vX := t.valueX
	v := Subtract(Fraction{0, 1}, t.value)
	return Divide(v, vX)
}
