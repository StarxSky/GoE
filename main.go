/*
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


*/


package main

import (
	"fmt"

	"github.com/starxsky/GoE/autodiff"
	gomat "github.com/starxsky/GoE/math"
	goenet "github.com/starxsky/GoE/network"
)

func lossAct(t1, a13 *autodiff.Node) *autodiff.Node {
	diff := autodiff.Sub(t1, a13)
	return autodiff.Mul(autodiff.NewNode(0.5), autodiff.Pow(diff, autodiff.NewNode(2)))
}

func sigmoidAct(z *autodiff.Node) *autodiff.Node {
	return autodiff.Div(
		autodiff.NewNode(1),
		autodiff.Add(
			autodiff.NewNode(1),
			autodiff.Div(autodiff.NewNode(1), autodiff.Exp(z)),
		),
	)
}

func main() {
	data1 := gomat.Ones(1, 3)
	data2 := gomat.Ones(3, 1)
	gomat.CoutMat(gomat.MultiplyMatrix(data1, data2))

	fmt.Println("auto build on gcc")
	fmt.Println("---------autodiff for neraul network-----------")
	dataMine := gomat.Ones(3, 3)
	label := gomat.Ones(3, 1)
	weight1 := gomat.Ones(3, 3)
	bais1 := gomat.Zeros(3, 1)
	weight2 := gomat.Ones(3, 3)
	bais2 := gomat.Zeros(3, 1)
	for epoch := 0; epoch < 1; epoch++ {
		fmt.Println("---------epoch: ", epoch, "------------")
		gomat.CoutMat(weight1)
		inputDim := 3
		outputDim := 3
		sequential := goenet.NewNetwork(inputDim, outputDim)

		output1 := sequential.Forward(dataMine, weight1, bais1)
		output1WithoutAct := sequential.ForwardWithoutAct(dataMine, weight1, bais1)
		fmt.Println("output1_without_act:")
		gomat.CoutMat(output1WithoutAct)

		fmt.Println("output2:")
		output2 := sequential.Forward(output1, weight2, bais2)
		gomat.CoutMat(output2)
		output2WithoutAct := sequential.ForwardWithoutAct(output1, weight2, bais2)
		fmt.Println("output2_without_act:")
		gomat.CoutMat(output2WithoutAct)

		outputEnd := sequential.EndLayerBackward(label, output2WithoutAct, lossAct, sigmoidAct)
		backward3 := sequential.Backward(outputEnd, output1WithoutAct, weight2, sigmoidAct)

		weight2Grad := gomat.MultiplyMatrix(outputEnd, gomat.GetT(output1))
		fmt.Println("weight_2_grad:")
		gomat.CoutMat(weight2Grad)
		weight1Grad := gomat.MultiplyMatrix(backward3, gomat.GetT(dataMine))
		fmt.Println("weight_1_grad:")
		gomat.CoutMat(weight1Grad)

		weight1 = gomat.Subtract(weight1, gomat.TimesMat(0.001, weight1Grad))
		bais1 = gomat.Subtract(bais1, gomat.TimesMat(0.001, backward3))
		weight2 = gomat.Subtract(weight2, gomat.TimesMat(0.001, weight2Grad))
		bais2 = gomat.Subtract(bais2, gomat.TimesMat(0.001, outputEnd))
		fmt.Println("neraul end;")
	}

	fmt.Println()

	A := gomat.Matrix{
        Rows: 2,
        Cols: 3,
        Data: []float64{
            1, 2, 3,
            4, 5, 6,
        },
    }

    B := gomat.Matrix{
        Rows: 3,
        Cols: 2,
        Data: []float64{
            7, 8,
            9, 10,
            11, 12,
        },
    }
	fmt.Println(" ======== ")

    C := gomat.MultiplyMatrix(A, B)
	fmt.Printf("A: %8.2f \nB: %8.2f \n", A.Data, B.Data)

    gomat.Print(C)
	

}
