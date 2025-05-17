# 📊 statistical-go

[![Go Report Card](https://goreportcard.com/badge/github.com/cyber-mountain-man/statistical-go)](https://goreportcard.com/report/github.com/cyber-mountain-man/statistical-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)
![Coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen)

A lightweight, dependency-free Go package for performing descriptive and inferential statistical calculations using the standard library only.

---

## ✨ Features

- 📈 **Descriptive Statistics**
  - `Mean`, `Median`, `Mode`
  - `Min`, `Max`, `Range`
  - `Quartiles` (Q1, Q2, Q3)
  - `Variance` & `Standard Deviation`
  - `Z-Score`
  - `Covariance`, `Pearson Correlation`
  - `Skewness`, `Kurtosis`

- 🧪 **Test Coverage**
  - ✅ 100% coverage via unit tests
  - ⚙️ Thorough testing for edge cases, empty inputs, and ties

- 💡 **Design Philosophy**
  - No external dependencies
  - Beginner-friendly, readable code
  - Modular for future extensions (e.g., regression, probability, hypothesis testing)

---

## 🛠️ Installation

```bash
go get github.com/cyber-mountain-man/statistical-go/stat
````

---

## 📦 Usage

```go
package main

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

---

## 🧪 Run Tests

```bash
go test -v ./stat
```

### Generate Coverage Report

```bash
go test -coverprofile=coverage.out ./stat
go tool cover -func=coverage.out         # summary in terminal
go tool cover -html=coverage.out         # opens browser view
```

---

## 📁 Project Structure

```
statistical-go/
├── stat/
│   ├── descriptive.go      # All stat functions
│   └── descriptive_test.go # Full test coverage
├── go.mod
└── README.md
```

---

## 📌 Roadmap

* [x] Descriptive statistics
* [x] Correlation & variability
* [ ] Probability distributions (Normal, Binomial)
* [ ] Hypothesis testing (z-test, t-test)
* [ ] Linear regression
* [ ] CLI tool or WebAPI version

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## 🙌 Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change or add.
