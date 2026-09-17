package math
// This file stores functions used for NeuralNet workloads or other non-linear transformations.
import (
	stdmath "math"  // 'import "math" as stdmath '
)


//Define the Sigmoid Posiliblity Function
func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + stdmath.Exp(-x))
}

func GetSigmoid(temp float64) float64 {
	return Sigmoid(temp)
}

func ESigmoid(m Matrix) Matrix {
	result := NewMatrix(m.Rows, m.Cols)
	for i := range m.Data {
		result.Data[i] = Sigmoid(m.Data[i])
	}
	return result
}
//Define the ReLU activation 
func EdgeRelu(val float64) float64 {
	if val > 0 {
		return val
	}
	return 0
}

// acts on the Matrix
func Relu(m Matrix) Matrix {
	result := NewMatrix(m.Rows, m.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			result.Set(i, j, EdgeRelu(m.At(i, j)))
		}
	}
	return result
}

func MatRelu(m Matrix) Matrix {
	return Relu(m)
}

/*
Mean Squared Loss (MSE):
\mathcal{L}(A,B)
=
\sum_{i=1}^{m}\sum_{j=1}^{n}
(A_{ij}-B_{ij})^2
*/
func MatSqLoss(m1, m2 Matrix) Matrix {
	result := NewMatrix(m1.Rows, 1)
	for i := 0; i < m1.Rows; i++ {
		diff := m1.At(i, 0) - m2.At(i, 0)
		result.Set(i, 0, diff*diff)
	}
	return result
}
