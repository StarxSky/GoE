/*package math
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


*/

package math

import (
	"fmt"
	"math/rand"
)

type Matrix struct {
	Rows int
	Cols int
	Data []float64
	Size int
}

func (m Matrix) At(row, col int) float64 {
	return m.Data[row*m.Cols+col]
}

func (m *Matrix) Set(row, col int, value float64) {
	m.Data[row*m.Cols+col] = value
}

func NewMatrix(rows, cols int) Matrix {
	return Matrix{
		Rows: rows,
		Cols: cols,
		Data: make([]float64, rows*cols),
		Size: rows * cols,
	}
}

func CreateMatrix(ro, co int) Matrix {
	return NewMatrix(ro, co)
}

func Ones(ro, co int) Matrix {
	m := NewMatrix(ro, co)
	for i := range m.Data {
		m.Data[i] = 1
	}
	return m
}

func Zeros(ro, co int) Matrix {
	return NewMatrix(ro, co)
}

func CreateRandMat(xDim, yDim int) Matrix {
	result := NewMatrix(xDim, yDim)
	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			result.Set(x, y, float64(rand.Intn(30000))*0.0001-1)
		}
	}
	return result
}

func ChangeValue(m Matrix, row, col int, value float64) {
	m.Data[row*m.Cols+col] = value
}
//
// Sent the Matrix's information to the console
func Print(m Matrix) {
	fmt.Println("\nMatrix: ")
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			fmt.Printf("[%8.2f ]", m.Data[i*m.Cols+j])
		}
		fmt.Println()
	}
	fmt.Printf("Size: %d \nShape: %d x %d \n", m.Size, m.Rows, m.Cols) // %decimal
}

// The Member function of the Matrix Structure. 
// Bind the `Print()` method for the `Matrix` data type. (Based on the pointer)
func (m *Matrix) Print() {
	fmt.Println("\nMatrix: ")
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			fmt.Printf("[%8.2f ]", m.Data[i*m.Cols+j])
		}
		fmt.Println()
	}
	fmt.Printf("Size: %d \nShape: %d x %d \n", m.Size, m.Rows, m.Cols) // %decimal
}

// C++ Format
func CoutMat(m Matrix) {
	fmt.Println("[")
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			fmt.Printf("%12.6f ", m.Data[i*m.Cols+j])
		}
		fmt.Println()
	}
	fmt.Println("]")
}


// Matrix with 3 dimensions
type Matrix3d struct {
	Dep  int
	Wid  int
	High int
	Mats []Matrix
}


func CreateMatrix3d(depth, height, width int) Matrix3d {
	mats := make([]Matrix, depth) // creat the arry with 'depth'
	for i := 0; i < depth; i++ {
		mats[i] = CreateRandMat(width, height)
	}
	return Matrix3d{Dep: depth, Wid: height, High: width, Mats: mats}
}

func CreateMatrix3dValue(depth, height, width int, value func(d int) Matrix) Matrix3d {
	mats := make([]Matrix, depth)
	for i := 0; i < depth; i++ {
		mats[i] = value(i)
	}
	return Matrix3d{Dep: depth, Wid: height, High: width, Mats: mats}
}

func (m Matrix3d) Dim(depth int) Matrix {
	return m.Mats[depth]
}

func GetShape3D(m Matrix3d) {
	fmt.Printf("EDGE Matrix Shape:(%d,%d,%d)\n", m.Dep, m.Wid, m.High)
}

func Matrix3dMean(m Matrix3d) float64 {
	var sum float64
	for i := 0; i < m.Dep; i++ {
		sum += MatrixMean(m.Mats[i])
	}
	return sum / float64(m.Dep)
}

func CoutMat3d(mat3d Matrix3d) {
	separatorLen := 20 + len(fmt.Sprintf("%d", mat3d.Dep))
	fmt.Println(stringsRepeat("=", separatorLen))
	for i := 0; i < mat3d.Dep; i++ {
		fmt.Printf("Matrix %d:\n", i)
		CoutMat(mat3d.Mats[i])
		fmt.Println()
	}
	fmt.Println(stringsRepeat("=", separatorLen))
}

type Matrix4d struct {
	Batch int
	Dep   int
	Wid   int
	High  int
	Mats  []Matrix3d
}

func CreateMatrix4d(bat, dep, wid, high int) Matrix4d {
	mats := make([]Matrix3d, bat)
	for i := 0; i < bat; i++ {
		mats[i] = CreateMatrix3d(dep, wid, high)
	}
	return Matrix4d{Batch: bat, Dep: dep, Wid: wid, High: high, Mats: mats}
}

func CreateMatrix4dValue(bat, dep, wid, high int, value func(b int) Matrix3d) Matrix4d {
	mats := make([]Matrix3d, bat)
	for i := 0; i < bat; i++ {
		mats[i] = value(i)
	}
	return Matrix4d{Batch: bat, Dep: dep, Wid: wid, High: high, Mats: mats}
}

func (m Matrix4d) Dim(batch int) Matrix3d {
	return m.Mats[batch]
}

type Matrix4dDimensions struct {
	Batch int
	Dep   int
	Wid   int
	High  int
}

func GetShape4D(m Matrix4d) Matrix4dDimensions {
	dims := Matrix4dDimensions{Batch: m.Batch, Dep: m.Dep, Wid: m.Wid, High: m.High}
	fmt.Printf("EDGE Matrix Shape: (%d,%d,%d,%d)\n", dims.Batch, dims.Dep, dims.Wid, dims.High)
	return dims
}

func PrintShape4D(m Matrix4d) {
	fmt.Printf("EDGE Matrix Shape: (%d,%d,%d,%d)\n", m.Batch, m.Dep, m.Wid, m.High)
}

func CoutMat4d(mat4d Matrix4d) {
	separatorLen := 20 + len(fmt.Sprintf("%d", mat4d.Batch))
	fmt.Println(stringsRepeat("=", separatorLen))
	for i := 0; i < mat4d.Batch; i++ {
		fmt.Printf("Batch %d:\n", i)
		CoutMat3d(mat4d.Mats[i])
		fmt.Println()
	}
	fmt.Println(stringsRepeat("=", separatorLen))
}

func stringsRepeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

type StringMatrix struct {
	Rows int
	Cols int
	Data [][]string

}

func CreateStrMa(ro, co int) StringMatrix {
	
	data := make([][]string, ro)
	for i := 0; i < ro; i++ {
		data[i] = make([]string, co)
		for j := 0; j < co; j++ {
			data[i][j] = "edge"
		}
	}
	return StringMatrix{Rows: ro, Cols: co, Data: data}
}

func CoutStrMat(m StringMatrix) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			fmt.Printf("%s,", m.Data[i][j])
		}
		fmt.Println()
	}
}
