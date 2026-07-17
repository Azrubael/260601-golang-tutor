package main

/* https://en.wikipedia.org/wiki/List_of_quantum_logic_gates
 * List of quantum logic gates
 */

import (
	"fmt"
	"math"
	"math/rand"
)

// Complex matrix type
type Matrix [][]complex128

// Vector type
type Vector []complex128

// ── single‑qubit primitives ────────────────────────────────────────────────

// I returns the 1-qubit identity gate
func I() Matrix {
	return Matrix{
		{1, 0},
		{0, 1},
	}
}

// X returns the flip gate (Pauli-X or NOT function)
func X() Matrix {
	return Matrix{
		{0, 1},
		{1, 0},
	}
}

// H returns the Hadamard gate
func H() Matrix {
	s := 1.0 / math.Sqrt(2)
	return Matrix{
		{complex(s, 0), complex(s, 0)},
		{complex(s, 0), complex(-s, 0)},
	}
}

// ── two‑qubit primitives ───────────────────────────────────────────────────

// SWAP returns the swap gate (swaps two qubits)
func SWAP() Matrix {
	m := identity(4)
	// Swap rows 1 and 2
	m[1], m[2] = m[2], m[1]
	return m
}

// CX returns the controlled NOT gate (CNOT function)
func CX() Matrix {
	m := identity(4)
	// Swap rows 2 and 3
	m[2], m[3] = m[3], m[2]
	return m
}

// ── helper functions ───────────────────────────────────────────────────────

// identity returns an n×n identity matrix
func identity(n int) Matrix {
	m := make(Matrix, n)
	for i := 0; i < n; i++ {
		m[i] = make(Vector, n)
		m[i][i] = 1
	}
	return m
}

// kronecker computes the Kronecker product of two matrices
func kronecker(a, b Matrix) Matrix {
	rows := len(a) * len(b)
	cols := len(a[0]) * len(b[0])
	result := make(Matrix, rows)
	for i := range rows {
		result[i] = make(Vector, cols)
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			for pi := 0; pi < len(b); pi++ {
				for pj := 0; pj < len(b[0]); pj++ {
					result[i*len(b)+pi][j*len(b[0])+pj] = a[i][j] * b[pi][pj]
				}
			}
		}
	}
	return result
}

// matrixMult multiplies two matrices
func matrixMult(a, b Matrix) Matrix {
	rows := len(a)
	cols := len(b[0])
	result := make(Matrix, rows)
	for i := range rows {
		result[i] = make(Vector, cols)
		for j := 0; j < cols; j++ {
			for k := 0; k < len(b); k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

// matrixVectorMult multiplies a matrix by a vector
func matrixVectorMult(m Matrix, v Vector) Vector {
	result := make(Vector, len(m))
	for i := range m {
		for j := 0; j < len(v); j++ {
			result[i] += m[i][j] * v[j]
		}
	}
	return result
}

// log2 returns the base-2 logarithm of n
func log2(n int) int {
	if n <= 0 {
		panic("log2: input must be positive")
	}
	return int(math.Log2(float64(n)))
}

// ── core routine ───────────────────────────────────────────────────────────

// Apply applies a sequence of gates to a quantum state vector
func Apply(v Vector, gates ...Matrix) Vector {
	numQubits := log2(len(v))
	var m Matrix
	// var err error

	for _, gate := range gates {
		gateSize := log2(len(gate))

		var expandedGate Matrix
		if gateSize == 1 {
			// Single-qubit gate: expand to full system (I ⊗ gate)
			eye := identity(1 << uint(numQubits-1))
			expandedGate = kronecker(eye, gate)
		} else {
			// Multi-qubit gate: use as-is
			expandedGate = gate
		}

		// Compose gates
		if m == nil {
			m = expandedGate
		} else {
			m = matrixMult(expandedGate, m)
		}
	}

	if m == nil {
		m = identity(len(v))
	}

	return matrixVectorMult(m, v)
}

// conj returns the complex conjugate
func conj(c complex128) complex128 {
	return complex(real(c), -imag(c))
}

// Observe collapses the wave function and returns a random basis index
func Observe(v Vector) int {
	// Calculate probabilities (|amplitude|^2)
	probs := make([]float64, len(v))
	var total float64
	for i, amp := range v {
		prob := real(amp * conj(amp))
		probs[i] = prob
		total += prob
	}

	// Normalize
	for i := range probs {
		probs[i] /= total
	}

	// Random selection weighted by probabilities
	r := rand.Float64()
	cumulative := 0.0
	for i, p := range probs {
		cumulative += p
		if r <= cumulative {
			return i
		}
	}
	return len(v) - 1
}

// ── main ───────────────────────────────────────────────────────────────────

func main() {
	a := Vector{1, 0, 0, 0}
	fmt.Println("measure:", Observe(a))

	a = Apply(a, X(), H())
	fmt.Println("after X,H:", Observe(a))

	a = Vector{0, 0, 0, 1}
	fmt.Println("measure:", Observe(a))

	a = Apply(a, I())
	fmt.Println("after I:", Observe(a))

	underscore := []byte{}
	for range 50 {
		underscore = append(underscore, 45) // ASCII code for '-'
	}

	fmt.Println(string(underscore))

	a = Apply(a, I(), SWAP())
	fmt.Println("after I,SWAP:", Observe(a))
}
