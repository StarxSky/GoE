# GoE

> **GoE — a Golang Scientific Computing Engine for Machine Learning on Any Platform.**

GoE is an experimental scientific computing engine written in **Golang**.

The goal of GoE is to build a lightweight, portable, and extensible numerical computing foundation for future **machine learning and scientific computing** workloads.

The project is currently in its **early development stage**, starting from fundamental matrix operations and gradually expanding toward a more complete numerical computing system.

---

## 🚧 Project Status

**Early Development**

Currently, GoE provides a basic `Matrix` implementation with support for:

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
* [ ] Automatic differentiation
* [ ] Machine learning components

The API and internal data structures are expected to change as the project develops.

---

## ✨ Current Features

### Matrix

The current implementation provides a basic matrix abstraction:

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

### Clone

```bash
git clone https://github.com/StarxSky/GoE.git
cd GoE
```

### Run

```bash
go run .
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

---

## 📦 Project Structure

The current project structure is intentionally small:

```text
GoE/
├── go.mod
├── main.go
│
├── math/ # The math operations 
│   └── ...
│
└── README.md
```

The `math` package currently contains the matrix implementation and related numerical operations.

The project will gradually be reorganized into more specialized packages as functionality grows.

A possible future architecture is:

```text
GoE/
├── math/
│   ├── matrix/
│   ├── vector/
│   ├── linalg/
│   └── ...
│
├── tensor/
│
├── autodiff/
│
├── nn/
│
└── ...
```

---

## 🛣️ Roadmap

GoE will be developed incrementally, starting from basic numerical primitives.

### Phase 1 — Matrix Foundation

* [x] Matrix data structure
* [x] Matrix construction
* [x] Element access
* [x] Element modification
* [x] Matrix printing
* [x] Matrix addition
* [x] Matrix multiplication
* [ ] Scalar multiplication
* [ ] Matrix subtraction
* [ ] Transpose
* [ ] Element-wise operations
* [ ] Identity matrix
* [ ] Zero matrix
* [ ] Random matrix generation
* [ ] Matrix slicing

### Phase 2 — Linear Algebra

* [ ] Vector
* [ ] Dot product
* [ ] Vector norms
* [ ] Gaussian elimination
* [ ] Matrix inverse
* [ ] Determinant
* [ ] LU decomposition
* [ ] QR decomposition
* [ ] Eigenvalues and eigenvectors
* [ ] Singular Value Decomposition (SVD)

### Phase 3 — Numerical Computing

* [ ] Numerical differentiation
* [ ] Numerical integration
* [ ] Interpolation
* [ ] Numerical optimization
* [ ] Statistics
* [ ] Probability distributions
* [ ] Random number generation

### Phase 4 — Tensor Computing

* [ ] N-dimensional tensors
* [ ] Tensor indexing
* [ ] Reshape
* [ ] Broadcasting
* [ ] Reduction operations
* [ ] Tensor arithmetic
* [ ] Efficient memory management

### Phase 5 — Automatic Differentiation

* [ ] Computational graph
* [ ] Forward-mode autodiff
* [ ] Reverse-mode autodiff
* [ ] Gradient calculation
* [ ] Backpropagation

### Phase 6 — Machine Learning

* [ ] Linear regression
* [ ] Logistic regression
* [ ] Loss functions
* [ ] Optimizers
* [ ] Activation functions
* [ ] Neural network layers
* [ ] Training utilities
* [ ] Model serialization

### Phase 7 — Hardware Acceleration

Potential future backends include:

```text
                 GoE
                  │
        ┌─────────┴─────────┐
        │                   │
       CPU                 GPU
        │                   │
   ┌────┴────┐         ┌────┴────┐
   │         │         │         │
 x86-64    ARM64     CUDA      Metal
```

The long-term goal is to provide a unified numerical API while allowing different hardware backends to be explored independently.


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

See the [LICENSE](LICENSE) file for the full license text.

Copyright © 2026 StarxSky

