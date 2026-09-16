package math 



func Mul[T int | float64 | int64 ](a T, b T) T { 
	return a * b 
};

func MultiplyMatrix(m1 ,m2 Matrix) Matrix{
	if m1.Cols != m2.Rows {
		panic("[Error Inputs] Cannot Multiply matrices :dimensions do not match, please check the shape of input.")

	}

	result := NewMatrix(m1.Rows, m2.Cols)
	for i := 0;  i < m1.Rows; i ++ {
		for j := 0; j < m2.Cols; j++ {
			sum := 0.0; 
			for k := 0; k < m1.Cols; k++ {
				sum += m1.Data[i*m1.Cols+k] * m2.Data[k*m2.Cols+j]
			}
			result.Data[i*result.Cols+j] = sum;
		}
	}
	return  result;
}