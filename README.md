# Helper Library

A Go library providing helper functions for calculations and logging.

## Features

### Calculator
- Add: Add two integers
- Subtract: Subtract two integers
- Multiply: Multiply two integers
- Divide: Divide two integers (with zero check)

### Logger
- Info: Log information messages
- Warn: Log warning messages
- Error: Log error messages

## Installation

```bash
go get github.com/yanmoai6666-star/helper
```

## Usage

### Calculator

```go
import "github.com/yanmoai6666-star/helper/pkg/calculator"

func main() {
	result := calculator.Add(10, 5) // 15
	result = calculator.Subtract(10, 5) // 5
	result = calculator.Multiply(10, 5) // 50
	result = calculator.Divide(10, 5) // 2
}
```

### Logger

```go
import "github.com/yanmoai6666-star/helper/internal/logger"

func main() {
	log := logger.NewLogger("INFO")
	log.Info("This is an info message")
	log.Warn("This is a warning message")
	log.Error("This is an error message")
}
```

## License

MIT
