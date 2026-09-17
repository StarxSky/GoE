package math
/* Reference : https://docs.pytorch.org/docs/main/generated/torch.nn.modules.conv.Conv2d.html
This file implements all the functions required for the Convolution Operations, that include the Padding,  BatchPadding, Rot etc.

*/

func EdgePadding(m Matrix, shape1, shape2 int) Matrix {
	result := NewMatrix(shape1, shape2)
	rows := m.Rows
	cols := m.Cols
	topPad := (shape1 - rows) / 2
	leftPad := (shape2 - cols) / 2
	for i := 0; i < shape1; i++ {
		for j := 0; j < shape2; j++ {
			if i < topPad || i >= rows+topPad || j < leftPad || j >= cols+leftPad {
				result.Set(i, j, 0.0)
			} else {
				result.Set(i, j, m.At(i-topPad, j-leftPad))
			}
		}
	}
	return result
}


func CreatePadding4D(X Matrix4d, pad int) Matrix4d {
	n := X.Batch
	c := X.Dep
	h := X.Wid
	w := X.High
	hPad := h + 2*pad
	wPad := w + 2*pad
	output := CreateMatrix4dValue(n, c, hPad, wPad, func(b int) Matrix3d {
		mats := make([]Matrix, c)
		for ch := 0; ch < c; ch++ {
			mats[ch] = EdgePadding(X.Mats[b].Mats[ch], hPad, wPad)
		}
		return Matrix3d{Dep: c, Wid: hPad, High: wPad, Mats: mats}
	})
	return output
}



func ConvElement(m, kernel Matrix, kernelSize, stride int) Matrix {
	rows := (m.Rows-kernelSize)/stride + 1
	cols := (m.Cols-kernelSize)/stride + 1
	result := NewMatrix(rows, cols)
	for x := 0; x <= m.Rows-kernelSize; x += stride {
		for y := 0; y <= m.Cols-kernelSize; y += stride {
			crop := Iloc(m, x, x+kernelSize, y, y+kernelSize)
			result.Set(x/stride, y/stride, MatrixSum(MulSimple(crop, kernel)))
		}
	}
	return result
}

func convolveChannels(channels []Matrix, filters [][]Matrix, kernelSize, stride, rowLen, colLen int) []Matrix {
	maps := make([]Matrix, len(filters))
	for f := 0; f < len(filters); f++ {
		sum := NewMatrix(rowLen, colLen)
		for ch := 0; ch < len(channels); ch++ {
			element := ConvElement(channels[ch], filters[f][ch], kernelSize, stride)
			sum = Add(sum, element)
		}
		maps[f] = sum
	}
	return maps
}

func ConvTest(m Matrix3d, inputDim, outputChannels, stride, kernelSize, mode, padding int) Matrix3d {
	if mode != 0 {
		return Matrix3d{Dep: 1, Wid: 1, High: 1, Mats: []Matrix{Ones(1, 1)}}
	}
	rowLen := (m.Wid-kernelSize)/stride + 1
	colLen := (m.High-kernelSize)/stride + 1
	filters := make([][]Matrix, outputChannels)
	for f := 0; f < outputChannels; f++ {
		filters[f] = make([]Matrix, inputDim)
		for ch := 0; ch < inputDim; ch++ {
			filters[f][ch] = Ones(kernelSize, kernelSize)
		}
	}
	maps := convolveChannels(m.Mats[:inputDim], filters, kernelSize, stride, rowLen, colLen)
	return Matrix3d{Dep: outputChannels, Wid: rowLen, High: colLen, Mats: maps}
}

func ConvTestWithOutput(m Matrix3d, inputDim, outputChannels, stride, kernelSize, mode int, verbose bool) Matrix3d {
	if verbose {
		CoutMat3d(m)
	}
	paddingWid := stride - (m.Wid-kernelSize)%stride
	if paddingWid == stride {
		paddingWid = 0
	}
	paddingHigh := stride - (m.High-kernelSize)%stride
	if paddingHigh == stride {
		paddingHigh = 0
	}
	padded := make([]Matrix, inputDim)
	for rgb := 0; rgb < inputDim; rgb++ {
		padded[rgb] = EdgePadding(m.Mats[rgb], m.Mats[rgb].Rows+paddingHigh, m.Mats[rgb].Cols+paddingWid)
		if verbose {
			CoutMat(padded[rgb])
		}
	}
	filters := make([][]Matrix, outputChannels)
	for f := 0; f < outputChannels; f++ {
		filters[f] = make([]Matrix, inputDim)
		for ch := 0; ch < inputDim; ch++ {
			filters[f][ch] = Ones(kernelSize, kernelSize)
		}
	}
	rowLen := (m.Wid-kernelSize+2*paddingWid)/stride + 1
	colLen := (m.High-kernelSize+2*paddingHigh)/stride + 1
	maps := convolveChannels(padded, filters, kernelSize, stride, rowLen, colLen)
	if verbose {
		for f := 0; f < outputChannels; f++ {
			CoutMat(maps[f])
		}
	}
	return Matrix3d{Dep: outputChannels, Wid: rowLen, High: colLen, Mats: maps}
}

func Rot180(input Matrix) Matrix {
	height := input.Rows
	width := input.Cols
	output := NewMatrix(height, width)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			output.Set(i, j, input.At(height-1-i, width-1-j))
		}
	}
	return output
}

func BatchConvTest(m4 Matrix4d, inputDim, outputChannels, stride, kernelSize, mode int, verbose bool) Matrix4d {
	outputs := make([]Matrix3d, m4.Batch)
	for b := 0; b < m4.Batch; b++ {
		outputs[b] = ConvTestWithOutput(m4.Mats[b], inputDim, outputChannels, stride, kernelSize, mode, verbose)
	}
	return Matrix4d{Batch: m4.Batch, Dep: outputChannels, Wid: outputs[0].Wid, High: outputs[0].High, Mats: outputs}
}

func FetchSkipOne(m Matrix3d) Matrix4d {
	newDepth := m.Dep
	newHeight := m.Wid
	newWidth := m.High
	samples := make([][]Matrix, 4)
	for s := 0; s < 4; s++ {
		samples[s] = make([]Matrix, newDepth)
	}
	for i := 0; i < newDepth; i++ {
		slice := m.Mats[i]
		halfRows := newHeight / 2
		halfCols := newWidth / 2
		for s := 0; s < 4; s++ {
			samples[s][i] = NewMatrix(halfRows, halfCols)
		}
		for j := 0; j < newHeight; j++ {
			for k := 0; k < newWidth; k++ {
				idxX := j / 2
				idxY := k / 2
				if j%2 == 0 && k%2 == 0 {
					samples[0][i].Set(idxX, idxY, slice.At(j, k))
				} else if j%2 == 0 && k%2 == 1 {
					samples[1][i].Set(idxX, idxY, slice.At(j, k))
				} else if j%2 == 1 && k%2 == 0 {
					samples[2][i].Set(idxX, idxY, slice.At(j, k))
				} else {
					samples[3][i].Set(idxX, idxY, slice.At(j, k))
				}
			}
		}
	}
	outputs := make([]Matrix3d, 4)
	for s := 0; s < 4; s++ {
		outputs[s] = Matrix3d{Dep: newDepth, Wid: newHeight / 2, High: newWidth / 2, Mats: samples[s]}
	}
	return Matrix4d{Batch: 4, Dep: newDepth, Wid: newHeight / 2, High: newWidth / 2, Mats: outputs}
}
