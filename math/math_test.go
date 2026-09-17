package math
// For debug the 'math' package

import (
	stdmath "math" // 'import "math" as stdmath '
	"testing"
)

// Test the Matrix Multiplication. 
func TestMultiply(t *testing.T) {
	a := Matrix{Rows: 2, Cols: 3, Data: []float64{1, 2, 3, 4, 5, 6}}
	b := Matrix{Rows: 3, Cols: 2, Data: []float64{7, 8, 9, 10, 11, 12}}
	c := MultiplyMatrix(a, b)
	if c.Rows != 2 || c.Cols != 2 {
		t.Fatalf("shape got %dx%d want 2x2", c.Rows, c.Cols)
	}
	if c.At(0, 0) != 58 || c.At(0, 1) != 64 || c.At(1, 0) != 139 || c.At(1, 1) != 154 {
		t.Fatalf("mul result wrong: %v", c.Data)
	}
}
// Test the Convolution 
func TestConvTestWithOutput(t *testing.T) {
	m := CreateMatrix3dValue(1, 3, 3, func(int) Matrix {
		return Ones(3, 3)
	})
	out := ConvTestWithOutput(m, 1, 1, 1, 2, 0, false)
	if out.Wid != 2 || out.High != 2 {
		t.Fatalf("out shape got %dx%d want 2x2", out.Wid, out.High)
	}
	for i := range out.Mats[0].Data {
		if got := out.Mats[0].Data[i]; stdmath.Abs(got-4.0) > 1e-9 {
			t.Fatalf("conv value got %v want 4", got)
		}
	}
}


func TestFetchSkipOne(t *testing.T) {
	m := Matrix3d{
		Dep:  1,
		Wid:  2,
		High: 2,
		Mats: []Matrix{{Rows: 2, Cols: 2, Data: []float64{1, 2, 3, 4}}},
	}
	out := FetchSkipOne(m)
	if out.Batch != 4 {
		t.Fatalf("batch got %d want 4", out.Batch)
	}
	if out.Mats[0].Mats[0].At(0, 0) != 1 {
		t.Fatalf("sample0 got %v want 1", out.Mats[0].Mats[0].At(0, 0))
	}
	if out.Mats[3].Mats[0].At(0, 0) != 4 {
		t.Fatalf("sample3 got %v want 4", out.Mats[3].Mats[0].At(0, 0))
	}
}

func TestReshape(t *testing.T) {
	m := Matrix{Rows: 1, Cols: 4, Data: []float64{1, 2, 3, 4}}
	r := Reshape(m, 2, 2)
	if r.At(0, 0) != 1 || r.At(1, 1) != 4 {
		t.Fatalf("reshape wrong: %v", r.Data)
	}
}
