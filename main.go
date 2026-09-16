package main; 


//import "rsc.io/quote";
import (
	"fmt"
	"github.com/starxsky/GoE/math"
)

func main () {
	m := math.NewMatrix(2, 3)

    m.Set(0, 0, 1)
    m.Set(0, 1, 2)
    m.Set(0, 2, 3)

    m.Set(1, 0, 4)
    m.Set(1, 1, 5)
    m.Set(1, 2, 6)

    //fmt.Println(m.At(0, 1))
	//fmt.Println(m.Data)
	fmt.Println()

	//m.Print() //test the Matrix cout

	A := math.Matrix{
        Rows: 2,
        Cols: 3,
        Data: []float64{
            1, 2, 3,
            4, 5, 6,
        },
    }

    B := math.Matrix{
        Rows: 3,
        Cols: 2,
        Data: []float64{
            7, 8,
            9, 10,
            11, 12,
        },
    }

    C := math.MultiplyMatrix(A, B)
	fmt.Printf("A: %8.2f \nB: %8.2f \n", A.Data, B.Data)

    C.Print()
	
	D := math.AdditionMatrices(A, B)
	D.Print()
	//fmt.Println(math.Mul(22, 3));
	//fmt.Println("Hello world!!!");
}