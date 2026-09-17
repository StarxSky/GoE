package network

import (
	"fmt"

	"github.com/starxsky/GoE/autodiff"
	gomat "github.com/starxsky/GoE/math" // For Matrix manipulations
)
// The NeuralNetwork Type
type Network struct {
	Input     int 
	NumNeuron int // The number of Neuron
}


func NewNetwork(input, numNeuron int) *Network {
	return &Network{Input: input, NumNeuron: numNeuron}
}

func sigmoidNode(z *autodiff.Node) *autodiff.Node {
	return autodiff.Div(
		autodiff.NewNode(1),
		autodiff.Add(
			autodiff.NewNode(1),
			autodiff.Div(autodiff.NewNode(1), autodiff.Exp(z)),
		),
	)
}

func (n *Network) Forward(data, weights, bais gomat.Matrix) gomat.Matrix {
	gomat.CoutMat(weights)
	gomat.CoutMat(data)
	output := gomat.MultiplyMatrix(weights, data)
	fmt.Println("out:")
	gomat.CoutMat(output)
	output = gomat.Add(output, bais)
	return gomat.ESigmoid(output)
}

func (n *Network) ForwardWithoutAct(data, weights, bais gomat.Matrix) gomat.Matrix {
	output := gomat.MultiplyMatrix(weights, data)
	return gomat.Add(output, bais)
}

func (n *Network) Backward(gradNext, outputBefore gomat.Matrix, weights gomat.Matrix, act func(*autodiff.Node) *autodiff.Node) gomat.Matrix {
	output := gomat.CloneMatrix(outputBefore)
	for index := 0; index < outputBefore.Rows; index++ {
		z := autodiff.NewNode(outputBefore.At(index, 0))
		anyone := act(z)
		output.Set(index, 0, anyone.Gradient(z))
	}
	return gomat.MulSimple(gomat.MultiplyMatrix(weights, gradNext), output)
}

func (n *Network) EndLayerBackward(label, actiVal gomat.Matrix, lossFun func(*autodiff.Node, *autodiff.Node) *autodiff.Node, actFun func(*autodiff.Node) *autodiff.Node) gomat.Matrix {
	lossAct := gomat.NewMatrix(actiVal.Rows, actiVal.Cols)
	actOutput := gomat.NewMatrix(actiVal.Rows, actiVal.Cols)
	for index := 0; index < lossAct.Rows; index++ {
		t1 := autodiff.NewNode(label.At(index, 0))
		z31 := autodiff.NewNode(actiVal.At(index, 0))
		a13 := sigmoidNode(z31)
		loss := lossFun(t1, a13)
		fmt.Printf("loss:%f\n", loss.Value())
		act := actFun(z31)
		actOutput.Set(index, 0, act.Gradient(z31))
		lossAct.Set(index, 0, loss.Gradient(a13))
	}
	midGradEnd := gomat.MulSimple(lossAct, actOutput)
	gomat.CoutMat(midGradEnd)
	return midGradEnd
}
