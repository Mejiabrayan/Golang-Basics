package main

import "fmt"

func main() {
	// prerequisite for profit calculation
	// 1. ask for revenue, expenses, and  tax rate
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses) // populate expenses by using a pointer (&)
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)
	// 2. calculate earnings before tax(EBT) and earnings after tax (profits)
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(ratio)
	fmt.Println(profit)

}
