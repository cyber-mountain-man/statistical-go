package main

import (
	"fmt"
	"github.com/cyber-mountain-man/statistical-go"
)

func main() {
	fmt.Println("=== Probability Rules Example ===")

	probA := 0.4
	probB := 0.5
	probAandB := 0.2

	add := statistical.AdditionRule(probA, probB, probAandB)
	fmt.Printf("Addition Rule: P(A ∪ B) = %.4f\n", add)

	indMult := statistical.MultiplicationRuleIndependent(probA, probB)
	fmt.Printf("Multiplication (Independent): P(A ∩ B) = %.4f\n", indMult)

	depMult := statistical.MultiplicationRuleDependent(probA, 0.5) // P(B|A)
	fmt.Printf("Multiplication (Dependent): P(A ∩ B) = %.4f\n", depMult)

	cond := statistical.ConditionalProbability(probAandB, probB)
	fmt.Printf("Conditional Probability: P(A|B) = %.4f\n", cond)

	compl := statistical.Complement(probA)
	fmt.Printf("Complement: P(¬A) = %.4f\n", compl)
}
