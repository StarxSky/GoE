package autodiff

import (
	stdmath "math"
	"testing"
)

func TestSinGrad(t *testing.T) {
	x := NewNode(0.5)
	y := Sin(x)
	if got := y.Gradient(x); stdmath.Abs(got-stdmath.Cos(0.5)) > 1e-9 {
		t.Fatalf("d(sin)/dx got %v want %v", got, stdmath.Cos(0.5))
	}
}

func TestLnGrad(t *testing.T) {
	x := NewNode(2.0)
	y := Ln(x)
	if got := y.Gradient(x); stdmath.Abs(got-0.5) > 1e-9 {
		t.Fatalf("d(ln)/dx got %v want %v", got, 0.5)
	}
}

func TestMulGrad(t *testing.T) {
	a := NewNode(3.0)
	b := NewNode(4.0)
	y := Mul(a, b)
	if ga, gb := y.Gradient(a), y.Gradient(b); ga != 4.0 || gb != 3.0 {
		t.Fatalf("mul grads got (%v,%v) want (4,3)", ga, gb)
	}
}

func TestPowGrad(t *testing.T) {
	x := NewNode(3.0)
	y := Pow(x, NewNode(2.0))
	if got := y.Gradient(x); stdmath.Abs(got-6.0) > 1e-9 {
		t.Fatalf("d(x^2)/dx got %v want 6", got)
	}
}
