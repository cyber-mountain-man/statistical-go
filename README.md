# 📊 statistical-go

[![Go Report Card](https://goreportcard.com/badge/github.com/cyber-mountain-man/statistical-go?t=1)](https://goreportcard.com/report/github.com/cyber-mountain-man/statistical-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)

![Coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen)

![CI](https://github.com/cyber-mountain-man/statistical-go/actions/workflows/ci.yml/badge.svg)


*A lightweight, dependency-free Go package for descriptive statistics, probability, distributions, hypothesis testing, Monte-Carlo simulations & regularised regression — all implemented with the Go standard library.*

---

## ✨ Features

- 📈 **Descriptive Statistics** – `Mean`, `Median`, `Mode`, `Min/Max/Range`, `Quartiles`,  
  `Variance`, `StdDev`, `Z-Score`, `Covariance`, `Pearson r`, `Skewness`, `Kurtosis`
- 🎲 **Monte-Carlo** – Estimate π (serial & parallel)
- 📊 **Probability Rules & Distributions**
  - Rules: Addition, Multiplication (Independent / Dependent), Union, Intersection, Complement
  - Distributions: Normal (PDF, CDF, **Inverse CDF**), Binomial, Uniform, Poisson, Exponential
- 🧪 **Hypothesis Testing** – Z-Test, T-Test (1-sample, Welch, Paired), χ² (GOF & Independence), One-Way ANOVA
- 🛠 **Regularised Regression** – Ridge & Lasso implementations
- 📦 **Unified API** – `statistical.go` provides one-stop wrappers
- ✅ **100 % Test Coverage** – core logic fully tested (examples excluded)
- 💡 **Design Philosophy** – zero external deps, beginner-friendly, modular

---

## 🆚 Go vs Python — Feature Comparison

| Task / Function            | Python (NumPy / SciPy / scikit-learn) | Go (`statistical-go`)                           |
|----------------------------|---------------------------------------|-------------------------------------------------|
| Mean                       | `numpy.mean(data)`                    | `stat.Mean(data)`                               |
| Median                     | `numpy.median(data)`                  | `stat.Median(data)`                             |
| Mode                       | `scipy.stats.mode(data)`              | `stat.Mode(data)`                               |
| Variance                   | `numpy.var(data)`                     | `stat.Variance(data)`                           |
| Standard Deviation         | `numpy.std(data)`                     | `stat.StdDev(data)`                             |
| Quartiles                  | `numpy.percentile(data,[25,75])`      | `stat.Quartiles(data)`                          |
| Z-Score                    | `scipy.stats.zscore(data)`            | `stat.ZScore(data)`                             |
| Covariance                 | `numpy.cov(x,y)`                      | `stat.Covariance(x,y)`                          |
| Pearson Correlation        | `scipy.stats.pearsonr(x,y)`           | `stat.PearsonCorrelation(x,y)`                  |
| Skewness                   | `scipy.stats.skew(data)`              | `stat.Skewness(data)`                           |
| Kurtosis                   | `scipy.stats.kurtosis(data)`          | `stat.Kurtosis(data)`                           |
| Normal Inverse CDF (PPF)   | `scipy.stats.norm.ppf(p, μ, σ)`       | `probability.NormalInverseCDF(p, μ, σ)`         |
| Binomial Quantile (PPF)    | `scipy.stats.binom.ppf()`             | `probability.BinomialQuantile()`                |
| Monte-Carlo π             | custom NumPy                          | `montecarlo.EstimatePi(n)`                      |
| Z-Test                     | `scipy.stats.norm.cdf()`              | `hypothesis.OneSampleZTest()`                   |
| T-Tests (all)              | `scipy.stats.ttest_*()`               | `hypothesis.*TTest()`                           |
| Chi-Square Tests           | `scipy.stats.chisquare()`             | `hypothesis.ChiSquareGoodnessOfFit()`           |
| ANOVA                      | `scipy.stats.f_oneway()`              | `hypothesis.OneWayANOVA()`                      |
| **Ridge Regression**       | `sklearn.linear_model.Ridge()`        | `models.RidgeRegression(X,y,λ)`                 |
| **Lasso Regression**       | `sklearn.linear_model.Lasso()`        | `models.LassoRegression(X,y,λ,iters)`           |

> ✅ **No third-party Go modules** – pure standard library. Perfect when you need statistical tools in Go with zero bloat.

---

## 🛠️ Installation

Install everything:

```bash
go get github.com/cyber-mountain-man/statistical-go
````

Or grab a specific package:

```bash
go get github.com/cyber-mountain-man/statistical-go/stat
go get github.com/cyber-mountain-man/statistical-go/probability
go get github.com/cyber-mountain-man/statistical-go/montecarlo
go get github.com/cyber-mountain-man/statistical-go/hypothesis
```

---

## 📦 Usage Examples

### 1 · Core `stat` package

```go
import (
	"fmt"
	"github.com/cyber-mountain-man/statistical-go/stat"
)

func main() {
	data := []float64{2,4,4,4,5,5,7,9}
	fmt.Println("Mean:", stat.Mean(data))
	fmt.Println("StdDev:", stat.StdDev(data))
	fmt.Println("Skewness:", stat.Skewness(data))
}
```

### 2 · Unified wrapper

```go
import (
	"fmt"
	"github.com/cyber-mountain-man/statistical-go"
)

func main() {
	data := []float64{1,2,3,4,5}
	fmt.Println("Mean:", statistical.Mean(data))
	fmt.Println("Estimated Pi:", statistical.EstimatePi(100_000))
}
```

---

## 🧪 Run Tests & Coverage

```bash
go test ./...                                         # all packages
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out                      # summary
go tool cover -html=coverage.out                      # interactive
```

> 💡 *Current core-logic coverage: **100 %**, examples intentionally excluded.*

---

## 📁 Project Structure

```
statistical-go/
├── stat/            # Descriptive statistics
├── probability/     # Probability rules & distributions
├── hypothesis/      # Z-test, T-test, ANOVA, Chi-Square
├── regression/      # Simple & multiple regression helpers
├── regression/models# Ridge & Lasso
├── montecarlo/      # Monte-Carlo simulations
├── examples/        # Demo programs (not tested)
├── statistical.go   # Unified wrapper API
└── README.md
```

---

## 📌 Roadmap

* [x] Descriptive statistics
* [x] Probability rules & distributions (incl. inverse CDFs)
* [x] Monte-Carlo simulation
* [x] Unified wrapper interface
* [x] Hypothesis testing (Z / T / ANOVA / Chi-Square)
* [x] **Linear, Ridge & Lasso regression**
* [ ] CLI tool or Web API
* [ ] Optional visualisation helpers (histograms, scatter)

---

## 📄 License

MIT – see [LICENSE](LICENSE).

---

## 🙌 Contributing

PRs welcome! For major changes please open an issue first and let’s discuss. 

