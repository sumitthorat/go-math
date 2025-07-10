# Math Package

A simple Go package that provides arithmetic operations.

## Installation

```bash
go get github.com/sumitthorat/math
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/sumitthorat/math/pkg/arithmetic"
)

func main() {
    // Create a new Arithmetic instance
    a := &arithmetic.Arithmetic{}
    
    // Use the Add method
    sum := a.Add([]int{1, 2, 3, 4, 5})
    fmt.Println("Sum:", sum)
    
    // Use the Multiply method
    product := a.Multiply([]int{1, 2, 3, 4, 5})
    fmt.Println("Product:", product)
}
```

## License

MIT
