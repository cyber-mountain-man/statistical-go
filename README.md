# 📊 statistical-go

[![Go Report Card](https://goreportcard.com/badge/github.com/cyber-mountain-man/statistical-go?t=1)](https://goreportcard.com/report/github.com/cyber-mountain-man/statistical-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)
![Coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen)

A lightweight, dependency-free Go package for performing descriptive statistics, probability rules, distributions, hypothesis testing, and Monte Carlo simulations — all using the Go standard library.

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

- 📊 **Probability Rules & Distributions**
  - Rules: Addition, Multiplication (Ind./Dep.), Union, Intersection, Complement
  - Distributions: Normal, Binomial, Uniform, Poisson, Exponential

- 🧪 **Hypothesis Testing**
  - Z-Test (1-sample & 2-sample)
  - T-Test (1-sample, 2-sample Welch, Paired)
  - One-Way ANOVA
  - Chi-Square Goodness of Fit & Independence

- 📦 **Unified API**
  - `statistical.go` wrapper simplifies access to core functions

- ✅ **Test Coverage**
  - 100% coverage for all core logic
  - `examples/` directory intentionally excluded from test suite

- 💡 **Design Philosophy**
  - No external dependencies
  - Beginner-friendly, readable code
  - Modular structure for easy expansion and real-world use

---

## 🆚 Go vs Python: Feature Comparison

| Statistical Task            | Python (NumPy/SciPy)              | Go (`statistical-go`)                          |
|-----------------------------|-----------------------------------|------------------------------------------------|
| Mean                        | `numpy.mean(data)`                | `stat.Mean(data)`                              |
| Median                      | `numpy.median(data)`              | `stat.Median(data)`                            |
| Mode                        | `scipy.stats.mode(data)`          | `stat.Mode(data)`                              |
| Variance                    | `numpy.var(data)`                 | `stat.Variance(data)`                          |
| Standard Deviation          | `numpy.std(data)`                 | `stat.StdDev(data)`                            |
| Quartiles                   | `numpy.percentile(data, [25,75])` | `stat.Quartiles(data)`                         |
| Z-Score                     | `scipy.stats.zscore(data)`        | `stat.ZScore(data)`                            |
| Covariance                  | `numpy.cov(x, y)`                 | `stat.Covariance(x, y)`                        |
| Pearson Correlation         | `scipy.stats.pearsonr(x, y)`      | `stat.PearsonCorrelation(x, y)`                |
| Skewness                    | `scipy.stats.skew(data)`          | `stat.Skewness(data)`                          |
| Kurtosis                    | `scipy.stats.kurtosis(data)`      | `stat.Kurtosis(data)`                          |
| Estimate Pi (Monte Carlo)   | custom NumPy code                 | `montecarlo.EstimatePi(n)`                     |
| Z-Test                      | `scipy.stats.norm.cdf()`          | `hypothesis.OneSampleZTest(...)`               |
| T-Test (1/2 sample/paired)  | `scipy.stats.ttest_*()`           | `hypothesis.*TTest(...)`                       |
| Chi-Square Test             | `scipy.stats.chisquare()`         | `hypothesis.ChiSquareGoodnessOfFit(...)`       |
| ANOVA                       | `scipy.stats.f_oneway()`          | `hypothesis.OneWayANOVA(...)`                  |

> ✅ `statistical-go` uses **no third-party libraries** — just pure Go.
> Ideal for developers who want to build statistical tooling in Go with full control and zero bloat.

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
go get github.com/cyber-mountain-man/statistical-go/hypothesis
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

### 2. Use Unified Wrapper

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
go tool cover -html=coverage.out         # interactive HTML
```

> 💡 *Current total test coverage: **100%** (core logic fully tested; examples excluded by design)*

---

## 📁 Project Structure

```
statistical-go/
├── stat/               # Descriptive statistics
├── probability/        # Probability rules & distributions
├── hypothesis/         # Z-test, T-test, ANOVA, Chi-Square
├── montecarlo/         # Monte Carlo simulations
├── examples/           # Demonstration programs (excluded from tests)
├── statistical.go      # Unified wrapper interface
├── go.mod              # Module definition
└── README.md           # Project documentation
```

---

## 📌 Roadmap

* [x] Descriptive statistics
* [x] Probability rules
* [x] Probability distributions
* [x] Monte Carlo simulation
* [x] Unified wrapper interface
* [x] Hypothesis testing (Z-test, T-test, ANOVA, Chi-Square)
* [ ] Linear regression
* [ ] CLI tool or Web API version
* [ ] Visualization support (e.g., histogram output)

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## 🙌 Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss your proposed additions or fixes.


