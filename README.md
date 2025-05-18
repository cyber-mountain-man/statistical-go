# 📊 statistical-go

[![Go Report Card](https://goreportcard.com/badge/github.com/cyber-mountain-man/statistical-go?t=1)](https://goreportcard.com/report/github.com/cyber-mountain-man/statistical-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)
![Coverage](https://img.shields.io/badge/Coverage-91.7%25-yellowgreen)

A lightweight, dependency-free Go package for performing descriptive statistics, probability rules, and Monte Carlo simulations — all using the Go standard library.

---

## ✨ Features

- 📈 **Descriptive Statistics**
  - `Mean`, `Median`, `Mode`
  - `Min`, `Max`, `Range`
  - `Quartiles` (Q1, Q2, Q3)
  - `Variance`, `Standard Deviation`, `Z-Score`
  - `Covariance`, `Pearson Correlation`
  - `Skewness`, `Kurtosis`

- 🎲 **Monte Carlo Simulations**
  - Estimate Pi using random sampling
  - Parallel version for performance on multicore systems

- ➕ **Probability Rules**
  - Addition Rule
  - Multiplication Rule (Independent & Dependent)
  - Conditional Probability, Complement, Union, Intersection

- 📦 **Unified API**
  - `statistical.go` wrapper simplifies access to core functions

- 🧪 **Test Coverage**
  - ✅ 100% for core packages (`stat`, `probability`, `montecarlo`, `statistical`)
  - ⚠️ `examples/` directory intentionally excluded from tests

- 💡 **Design Philosophy**
  - No external dependencies
  - Beginner-friendly, readable code
  - Modular structure allows for future expansion (e.g., distributions, hypothesis testing)

---

## 🛠️ Installation

Install the full library:

```bash
go get github.com/cyber-mountain-man/statistical-go
````

Or install a specific module:

```bash
go get github.com/cyber-mountain-man/statistical-go/stat
go get github.com/cyber-mountain-man/statistical-go/probability
go get github.com/cyber-mountain-man/statistical-go/montecarlo
```

---

## 📦 Usage Examples

### 1. Use Core `stat` Package

```go
import (
	"fmt"
	"github.com/cyber-mountain-man/statistical-go/stat"
)

func main() {
	data := []float64{2, 4, 4, 4, 5, 5, 7, 9}
	fmt.Println("Mean:", stat.Mean(data))
	fmt.Println("StdDev:", stat.StdDev(data))
	fmt.Println("Skewness:", stat.Skewness(data))
}
```

### 2. Use Wrapper via `statistical.go`

```go
import (
	"fmt"
	"github.com/cyber-mountain-man/statistical-go"
)

func main() {
	data := []float64{1, 2, 3, 4, 5}
	fmt.Println("Mean:", statistical.Mean(data))
	fmt.Println("Estimated Pi:", statistical.EstimatePi(100000))
}
```

---

## 🧪 Run Tests

```bash
go test ./...
```

### Generate Coverage Report

```bash
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out         # summary in terminal
go tool cover -html=coverage.out         # opens browser in detail
```

> 💡 *Current total test coverage: 91.7% (100% for all core logic)*

---

## 📁 Project Structure

```
statistical-go/
├── stat/               # Descriptive statistics
│   ├── descriptive.go
│   └── descriptive_test.go
├── probability/        # Probability rules
│   ├── basic.go
│   └── conditional.go
├── montecarlo/         # Monte Carlo simulations
│   ├── pi.go
│   └── pi_test.go
├── examples/           # Demonstration programs (excluded from tests)
│   └── main.go
├── statistical.go      # Facade wrapper for common functions
├── statistical_test.go # Wrapper tests
├── go.mod
└── README.md
```

---

## 📌 Roadmap

* [x] Descriptive statistics
* [x] Correlation & variability
* [x] Monte Carlo simulation
* [x] Unified wrapper interface
* [ ] Probability distributions (Normal, Binomial)
* [ ] Hypothesis testing (Z-test, T-test)
* [ ] Linear regression
* [ ] CLI tool or Web API version

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## 🙌 Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss your proposed additions or fixes.

