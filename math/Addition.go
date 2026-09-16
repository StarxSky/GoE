package math 

//import "fmt"

func AdditionMatrices(m1, m2 Matrix) Matrix {
	sum := NewMatrix(m1.Rows, m1.Cols)
	for i:=0; i < m1.Rows; i ++ {
		for j := 0; j < m2.Cols; j++ {
			_i := i*m1.Cols + j 
			sum.Data[_i] = m1.Data[_i] + m2.Data[_i]
		}
	}
	return  sum
}