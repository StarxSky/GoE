# GoE

> **GoE — a Golang Scientific Computing Engine for Machine Learning on Any Platform.**

GoE is an experimental scientific computing engine written in **Golang**.

The goal of GoE is to build a lightweight, portable, and extensible numerical computing foundation for future **machine learning and scientific computing** workloads.

The project is currently in its **early development stage**, starting from fundamental matrix operations and gradually expanding toward a more complete numerical computing system.

---

## 🚧 Project Status
This Engine has supported below Manipulations : 
* [x] Matrix representation
* [x] Matrix construction
* [x] Matrix element access
* [x] Matrix element modification
* [x] Matrix printing
* [x] Matrix addition
* [x] Matrix multiplication
* [x] Scalar multiplication
* [x] Matrix subtraction
* [x] Transpose
* [x] Vector operations
* [x] Linear algebra algorithms
* [x] Tensor operations
* [x] Automatic differentiation
* [x] Machine learning components

The API and internal data structures are expected to change as the project develops.

---

## ✨ Current Features
## Features

- **`autodiff`** — Reverse-mode automatic differentiation built on a computation graph. Supports scalar `Node` operations (add, sub, mul, div, pow, exp, log, trig, hyperbolic, `min`/`max`, `abs`) with vector, matrix, and recursive gradient retrieval.
- **`math`** — Dense matrix/ND-tensor layer: multiplication, element-wise ops, transpose, reshape, flatten, slicing (`Iloc`), row extraction, mean/variance, activations (sigmoid, ReLU, squared loss), padding, and 2D convolution (single & batch, with stride/padding). Includes 3D/4D tensor containers.
- **`network`** — Fully-connected layer-style forward and backward passes (with and without activation), plus an end-layer (loss × activation) backward helper — suitable for training a small MLP with gradient descent.


### Matrix

```go
type Matrix struct {
    Rows int
    Cols int
    Data []float64
}
```

Matrix elements are stored in a contiguous one-dimensional `[]float64`.

For a matrix:

```text
A = | 1  2  3 |
    | 4  5  6 |
```

the internal representation is:

```text
Data = [1, 2, 3, 4, 5, 6]
```

The element at row `i` and column `j` is mapped to:

```text
index = i * Cols + j
```

This provides a simple foundation for future numerical operations and optimization.

---

## 🧮 Matrix Operations

### Matrix Addition

Two matrices with the same shape can be added:

```go
C := math.AdditionMatrices(A, B)
```

Mathematically:

```text
C = A + B
```

---

### Matrix Multiplication

Matrix multiplication is supported when the dimensions are compatible:

```go
C := math.MultiplyMatrix(A, B)
```

For:

```text
A ∈ R^(m × n)

B ∈ R^(n × p)
```

the resulting matrix is:

```text
C ∈ R^(m × p)
```

with:

```text
C[i,j] = Σ A[i,k] B[k,j]
```

---

### Matrix Access

Individual elements can be accessed and modified using:

```go
m.Set(0, 0, 1)
```

and:

```go
value := m.At(0, 0)
```

A matrix can also be constructed directly:

```go
A := math.Matrix{
    Rows: 2,
    Cols: 3,
    Data: []float64{
        1, 2, 3,
        4, 5, 6,
    },
}
```

---

## 🚀 Quick Start

### Requirements

* Go 1.27.1 or compatible Go version
* A supported operating system
Go 1.27+ is required.

```sh
go build ./...
go run .
go test ./...
```


### Build

Build GoE into an executable binary:

```bash
go build -o goe .
```

Then run:

```bash
./goe
```

On Windows:

```powershell
go build -o goe.exe .
```



## Usage

### Compute gradients with the autodiff graph

```go
import "github.com/starxsky/GoE/autodiff"

func main() {
	x := autodiff.NewNode(2.0)
	y := autodiff.Exp(autodiff.Mul(x, autodiff.NewNode(3))) // exp(3x)
	dyDx := y.Gradient(x)                                   // 3 * exp(6)
}
```

Use `GradientVector`/`GradientMatrix` for multiple targets, and `GetGraph().NewRecording()` to reset the graph between passes.

### Matrix math

```go
import gomat "github.com/starxsky/GoE/math"

a := gomat.Ones(1, 3)
b := gomat.Ones(3, 1)
c := gomat.MultiplyMatrix(a, b) // 1x1 matrix
gomat.Print(c)
```

### Train a tiny network

`main.go` demonstrates a single-epoch training loop: two sigmoid layers forward-pass the data, `EndLayerBackward` computes the loss/activation gradient, `Backward` propagates it, and weights/biases are updated with SGD (learning rate 0.001).

## Package Layout

```
├── autodiff/    # graph + scalar differentiable ops, gradient API
│   ├── graph.go # computation graph, UIDs, edges
│   ├── node.go  # Node type and differentiable operations
│   └── node_test.go
├── math/        # matrices, vectors, tensors, convolution
│   ├── matrix.go    # Matrix / Matrix3d / Matrix4d, factories, printing
│   ├── multiply.go  # dense matrix multiplication
│   ├── operations.go# element-wise ops, transpose, reshape, slicing
│   ├── activation.go# sigmoid, ReLU, squared loss
│   └── convolve.go  # padding, 2D conv forward, batch conv
├── network/     # forward/backward passes for a fully-connected layer
│   └── network.go
└── main.go      # demo: matrix mult + one training epoch
```

## 🔬 Why Go?

Go provides several properties that make it interesting for scientific computing infrastructure:

* Simple language design
* Strong static typing
* Built-in concurrency
* Easy cross-compilation
* Fast compilation
* Low deployment overhead
* A relatively small runtime

GoE is also an exploration of whether Go can serve as a practical foundation for numerical computing and machine learning systems.

---

## 📈 Development Direction

The current implementation is intentionally small.

The planned development direction is:

```text
              GoE
               │
               ▼
          Matrix Core
               │
               ▼
        Linear Algebra
               │
               ▼
       Numerical Computing
               │
               ▼
            Tensor
               │
               ▼
    Automatic Differentiation
               │
               ▼
        Machine Learning
```

Each layer will build upon the previous one.

The immediate priority is to make the matrix layer more complete and reliable before moving to higher-level functionality.

---

## 🤝 Contributing

GoE is currently a personal experimental project, but contributions and discussions are welcome.

Possible contribution areas include:

* Matrix algorithms
* Numerical algorithms
* API design
* Performance optimization
* Testing
* Documentation
* Cross-platform development
* Machine learning functionality

For substantial changes, please open an issue first to discuss the proposed design.

---

## ⚠️ Disclaimer

GoE is currently an **experimental project** and is **not production-ready**.

APIs, data structures, package organization, and implementation details may change significantly during development.



---

## Vision

GoE aims to explore a simple question:

> **Can we build a modern scientific computing and machine learning engine from the ground up in Go?**

Starting from a matrix:

```text
A = | 1  2 |
    | 3  4 |
```

and gradually building toward:

```text
Matrix
   ↓
Linear Algebra
   ↓
Tensor
   ↓
Autodiff
   ↓
Optimization
   ↓
Machine Learning
```
---

## 📄 License

GoE is released under the **MIT License**.

Copyright © 2026 StarxSky

