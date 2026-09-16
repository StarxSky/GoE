package math
import "fmt"

type Matrix struct {
    Rows int
    Cols int
    Data []float64
	Size int 
}

func NewMatrix(rows, cols int) Matrix {
	
    return Matrix{
        Rows: rows,
        Cols: cols,
        Data: make([]float64, rows*cols),
		Size: rows*cols,
    }
}

func (m Matrix) At(i, j int) float64 {
    return m.Data[i*m.Cols+j]
}

func (m *Matrix) Set(i, j int, value float64) {
    m.Data[i*m.Cols+j] = value
}

func (m Matrix) Print() {
	fmt.Println("\nMatrix: ")
    for i := 0; i < m.Rows; i++ {
        for j := 0; j < m.Cols; j++ {
            fmt.Printf("[%8.2f ]", m.Data[i*m.Cols+j]) // skip i * Cols elements to get the target's row. 
        }
        fmt.Println()
    }
	fmt.Printf("Size: %d \nShape: %d x %d \n", m.Size, m.Rows, m.Cols)
	//fmt.Println(" ")
}