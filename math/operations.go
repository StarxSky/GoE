package math

/*
The math package provides fundamental matrix manipulation and numerical computation utilities for GoE. 
It currently supports element-wise arithmetic, scalar operations, matrix transformations, statistical calculations, slicing, and concatenation. 
Built on a row-major data layout, the package establishes a foundation for future scientific computing and machine learning capabilities.

Reference : https://docs.pytorch.org/docs/main/tensors.html

*/
func Add(m1, m2 Matrix) Matrix {
	if m1.Rows != m2.Rows || m1.Cols != m2.Cols {
		return CloneMatrix(m1)
	}
	result := NewMatrix(m1.Rows, m1.Cols)
	for i := range m1.Data {
		result.Data[i] = m1.Data[i] + m2.Data[i]
	}
	return result
}

func AddEle(m Matrix, toAdd float64) Matrix {
	result := NewMatrix(m.Rows, m.Cols)
	for i := range m.Data {
		result.Data[i] = m.Data[i] + toAdd
	}
	return result
}

func Subtract(m1, m2 Matrix) Matrix {
	result := NewMatrix(m1.Rows, m2.Cols)
	for i := range m1.Data {
		result.Data[i] = m1.Data[i] - m2.Data[i]
	}
	return result
}

func SubtractEle(m Matrix, toSubtract float64) Matrix {
	result := NewMatrix(m.Rows, m.Cols)
	for i := range m.Data {
		result.Data[i] = m.Data[i] - toSubtract
	}
	return result
}

func MatrixDivision(m Matrix, divisor float64) Matrix {
	result := NewMatrix(m.Rows, m.Cols)
	for i := range m.Data {
		result.Data[i] = m.Data[i] / divisor
	}
	return result
}

func MatrixPower(m Matrix) Matrix {
	result := NewMatrix(m.Rows, m.Cols)
	for i := range m.Data {
		result.Data[i] = m.Data[i] * m.Data[i]
	}
	return result
}

func MulSimple(m1, m2 Matrix) Matrix {
	if m1.Rows != m2.Rows || m1.Cols != m2.Cols {
		panic("[Error Inputs] Cannot element-wise Multiply matrices: shape A != B")
	}
	result := NewMatrix(m1.Rows, m1.Cols)
	for i := range m1.Data {
		result.Data[i] = m1.Data[i] * m2.Data[i]
	}
	return result
}

func TimesMat(factor float64, m Matrix) Matrix {
	result := NewMatrix(m.Rows, m.Cols)
	for i := range m.Data {
		result.Data[i] = factor * m.Data[i]
	}
	return result
}

func ScalarMat(factor float64, m Matrix) Matrix {
	return TimesMat(factor, m)
}

func Transpose(m Matrix) Matrix {
	result := NewMatrix(m.Cols, m.Rows)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			result.Set(j, i, m.At(i, j))
		}
	}
	return result
}

func GetT(m Matrix) Matrix {
	return Transpose(m)
}

func Flatten(m Matrix) Matrix {
	result := NewMatrix(1, m.Rows*m.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			result.Set(0, i*m.Cols+j, m.At(i, j))
		}
	}
	return result
}

func Reshape(m Matrix, rsRow, rsCol int) Matrix {
	result := NewMatrix(rsRow, rsCol)
	size := m.Rows * m.Cols
	for i := 0; i < size; i++ {
		srcRow := i / m.Cols
		srcCol := i % m.Cols
		dstRow := i / rsCol
		dstCol := i % rsCol
		result.Set(dstRow, dstCol, m.At(srcRow, srcCol))
	}
	return result
}

func MatrixRS(m Matrix, rsRow, rsCol int) Matrix {
	return Reshape(m, rsRow, rsCol)
}

func MatrixSum(m Matrix) float64 {
	var sum float64
	for i := range m.Data {
		sum += m.Data[i]
	}
	return sum
}

func MatrixMean(m Matrix) float64 {
	return MatrixSum(m) / float64(m.Size)
}

func MatrixVar(m Matrix) float64 {
	elements := float64(m.Size - 1)
	mean := MatrixMean(m)
	deviations := SubtractEle(m, mean)
	sumSquares := MatrixSum(MatrixPower(deviations)) + 1e-5
	return sumSquares / elements
}

func MatApply(m1, m2 Matrix, axis int) Matrix {
	if axis == 1 {
		result := NewMatrix(m1.Rows, m1.Cols+m2.Cols)
		for i := 0; i < m1.Rows; i++ {
			for j := 0; j < m1.Cols; j++ {
				result.Set(i, j, m1.At(i, j))
			}
			for j := 0; j < m2.Cols; j++ {
				result.Set(i, m1.Cols+j, m2.At(i, j))
			}
		}
		return result
	}
	result := NewMatrix(m1.Rows+m2.Rows, m2.Cols)
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Cols; j++ {
			result.Set(i, j, m1.At(i, j))
		}
	}
	for i := 0; i < m2.Rows; i++ {
		for j := 0; j < m2.Cols; j++ {
			result.Set(m1.Rows+i, j, m2.At(i, j))
		}
	}
	return result
}

func Head(m Matrix) Matrix {
	result := NewMatrix(6, m.Cols)
	for i := 0; i < 6 && i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			result.Set(i, j, m.At(i, j))
		}
	}
	return result
}

func Iloc(m Matrix, startX, endX, startY, endY int) Matrix {
	if endY == 0 {
		endY = m.Cols
	}
	if endX == 0 {
		endX = m.Rows
	}
	newRow := endX - startX
	newCol := endY - startY
	result := NewMatrix(newRow, newCol)
	for i := startX; i < endX; i++ {
		for j := startY; j < endY; j++ {
			result.Set(i-startX, j-startY, m.At(i, j))
		}
	}
	return result
}

func GetRow(m Matrix, index int) Matrix {
	result := NewMatrix(1, m.Cols)
	for j := 0; j < m.Cols; j++ {
		result.Set(0, j, m.At(index, j))
	}
	return result
}

func CloneMatrix(m Matrix) Matrix {
	result := NewMatrix(m.Rows, m.Cols)
	copy(result.Data, m.Data)
	return result
}
