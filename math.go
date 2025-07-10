// Package math provides mathematical operations
package math

import "github.com/sumitthorat/go-math/pkg/arithmetic"

// Arithmetic provides access to basic arithmetic operations
type Arithmetic struct {
	arith *arithmetic.Arithmetic
}

// New creates a new Arithmetic instance
func New() *Arithmetic {
	return &Arithmetic{
		arith: &arithmetic.Arithmetic{},
	}
}

// Add calculates the sum of a slice of integers
func (a *Arithmetic) Add(nums []int) int {
	return a.arith.Add(nums)
}

// Multiply calculates the product of a slice of integers
func (a *Arithmetic) Multiply(nums []int) int {
	return a.arith.Multiply(nums)
}
