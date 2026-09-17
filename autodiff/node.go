package autodiff

import (
	stdmath "math"
)

type MonadicOperation struct {
	Value float64
	Grad  float64
}

type DyadicOperation struct {
	Value float64
	Grad1 float64
	Grad2 float64
}

type PolyadicOperation struct {
	Value float64
	Grad  []float64
}

type Node struct {
	value float64
	uid   int64
}

func NewNode(value float64) *Node {
	return &Node{value: value, uid: NextUID()}
}

func (n *Node) Value() float64 {
	return n.value
}

func (n *Node) UID() int64 {
	return n.uid
}

func (n *Node) gradientRecursive(g *Graph, currentUID, stopUID int64) float64 {
	if currentUID == stopUID {
		return 1.0
	}
	var sum float64
	if g.IsConnected(currentUID) {
		for _, edge := range g.Get(currentUID) {
			sum += edge.Grad * n.gradientRecursive(g, edge.Uid, stopUID)
		}
	}
	return sum
}

func (n *Node) Gradient(target *Node) float64 {
	return n.gradientRecursive(GetGraph(), n.uid, target.uid)
}

func (n *Node) GradientVector(targets []*Node) []float64 {
	grad := make([]float64, len(targets))
	for i, t := range targets {
		grad[i] = n.gradientRecursive(GetGraph(), n.uid, t.uid)
	}
	return grad
}

func (n *Node) GradientMatrix(targets [][]*Node) [][]float64 {
	grad := make([][]float64, len(targets))
	for i, row := range targets {
		grad[i] = make([]float64, len(row))
		for j, t := range row {
			grad[i][j] = n.gradientRecursive(GetGraph(), n.uid, t.uid)
		}
	}
	return grad
}

func monadicOp(n *Node, fun func(float64) MonadicOperation) *Node {
	res := fun(n.value)
	result := NewNode(res.Value)
	GetGraph().Connect(result.uid, res.Grad, n.uid)
	return result
}

func dyadicOp(l, r *Node, fun func(float64, float64) DyadicOperation) *Node {
	res := fun(l.value, r.value)
	result := NewNode(res.Value)
	g := GetGraph()
	g.Connect(result.uid, res.Grad1, l.uid)
	g.Connect(result.uid, res.Grad2, r.uid)
	return result
}

func polyadicOp(nodes []*Node, fun func([]float64) PolyadicOperation) *Node {
	values := make([]float64, len(nodes))
	for i, n := range nodes {
		values[i] = n.value
	}
	res := fun(values)
	result := NewNode(res.Value)
	g := GetGraph()
	for i, n := range nodes {
		g.Connect(result.uid, res.Grad[i], n.uid)
	}
	return result
}

func Add(l, r *Node) *Node {
	return dyadicOp(l, r, func(a, b float64) DyadicOperation {
		return DyadicOperation{Value: a + b, Grad1: 1, Grad2: 1}
	})
}

func Sub(l, r *Node) *Node {
	return dyadicOp(l, r, func(a, b float64) DyadicOperation {
		return DyadicOperation{Value: a - b, Grad1: 1, Grad2: -1}
	})
}

func Mul(l, r *Node) *Node {
	return dyadicOp(l, r, func(a, b float64) DyadicOperation {
		return DyadicOperation{Value: a * b, Grad1: b, Grad2: a}
	})
}

func Div(l, r *Node) *Node {
	return dyadicOp(l, r, func(a, b float64) DyadicOperation {
		return DyadicOperation{Value: a / b, Grad1: 1 / b, Grad2: -a / (b * b)}
	})
}

func Sin(n *Node) *Node {
	return monadicOp(n, func(v float64) MonadicOperation {
		return MonadicOperation{Value: stdmath.Sin(v), Grad: stdmath.Cos(v)}
	})
}

func Cos(n *Node) *Node {
	return monadicOp(n, func(v float64) MonadicOperation {
		return MonadicOperation{Value: stdmath.Cos(v), Grad: -stdmath.Sin(v)}
	})
}

func Tan(n *Node) *Node {
	return monadicOp(n, func(v float64) MonadicOperation {
		return MonadicOperation{Value: stdmath.Tan(v), Grad: 1 / stdmath.Pow(stdmath.Cos(v), 2)}
	})
}

func Sinh(n *Node) *Node {
	return monadicOp(n, func(v float64) MonadicOperation {
		return MonadicOperation{Value: stdmath.Sinh(v), Grad: stdmath.Cosh(v)}
	})
}

func Cosh(n *Node) *Node {
	return monadicOp(n, func(v float64) MonadicOperation {
		return MonadicOperation{Value: stdmath.Cosh(v), Grad: stdmath.Sinh(v)}
	})
}

func Asin(n *Node) *Node {
	return monadicOp(n, func(v float64) MonadicOperation {
		return MonadicOperation{Value: stdmath.Asin(v), Grad: 1 / stdmath.Sqrt(1-v*v)}
	})
}

func Acos(n *Node) *Node {
	return monadicOp(n, func(v float64) MonadicOperation {
		return MonadicOperation{Value: stdmath.Acos(v), Grad: -1 / stdmath.Sqrt(1-v*v)}
	})
}

func Atan(n *Node) *Node {
	return monadicOp(n, func(v float64) MonadicOperation {
		return MonadicOperation{Value: stdmath.Atan(v), Grad: 1 / (1 + v*v)}
	})
}

func Tanh(n *Node) *Node {
	return monadicOp(n, func(v float64) MonadicOperation {
		t := stdmath.Tanh(v)
		return MonadicOperation{Value: t, Grad: 1 - t*t}
	})
}

func Log(x, base *Node) *Node {
	return dyadicOp(x, base, func(a, b float64) DyadicOperation {
		return DyadicOperation{
			Value: stdmath.Log(a) / stdmath.Log(b),
			Grad1: 1 / (a * stdmath.Log(b)),
			Grad2: -stdmath.Log(a) / (b * stdmath.Log(b) * stdmath.Log(b)),
		}
	})
}

func Log10(n *Node) *Node {
	return monadicOp(n, func(v float64) MonadicOperation {
		return MonadicOperation{Value: stdmath.Log10(v), Grad: 1 / (v * stdmath.Ln10)}
	})
}

func Ln(n *Node) *Node {
	return monadicOp(n, func(v float64) MonadicOperation {
		return MonadicOperation{Value: stdmath.Log(v), Grad: 1 / v}
	})
}

func Pow(x, p *Node) *Node {
	return dyadicOp(x, p, func(a, b float64) DyadicOperation {
		if a <= 0 {
			return DyadicOperation{Value: stdmath.Pow(a, b), Grad1: b * stdmath.Pow(a, b-1), Grad2: 0}
		}
		return DyadicOperation{
			Value: stdmath.Pow(a, b),
			Grad1: b * stdmath.Pow(a, b-1),
			Grad2: stdmath.Log(a) * stdmath.Pow(a, b),
		}
	})
}

func Exp(n *Node) *Node {
	return monadicOp(n, func(v float64) MonadicOperation {
		return MonadicOperation{Value: stdmath.Exp(v), Grad: stdmath.Exp(v)}
	})
}

func Sqrt(n *Node) *Node {
	return monadicOp(n, func(v float64) MonadicOperation {
		return MonadicOperation{Value: stdmath.Sqrt(v), Grad: 1 / (2 * stdmath.Sqrt(v))}
	})
}

func Abs(n *Node) *Node {
	return monadicOp(n, func(v float64) MonadicOperation {
		var sign float64
		switch {
		case v == 0:
			sign = 0
		case v < 0:
			sign = -1
		default:
			sign = 1
		}
		return MonadicOperation{Value: stdmath.Abs(v), Grad: sign}
	})
}

func Min(l, r *Node) *Node {
	return dyadicOp(l, r, func(a, b float64) DyadicOperation {
		switch {
		case a < b:
			return DyadicOperation{Value: a, Grad1: 1, Grad2: 0}
		case a > b:
			return DyadicOperation{Value: b, Grad1: 0, Grad2: 1}
		default:
			return DyadicOperation{Value: a, Grad1: 0, Grad2: 0}
		}
	})
}

func Max(l, r *Node) *Node {
	return dyadicOp(l, r, func(a, b float64) DyadicOperation {
		switch {
		case a > b:
			return DyadicOperation{Value: a, Grad1: 1, Grad2: 0}
		case a < b:
			return DyadicOperation{Value: b, Grad1: 0, Grad2: 1}
		default:
			return DyadicOperation{Value: a, Grad1: 0, Grad2: 0}
		}
	})
}
