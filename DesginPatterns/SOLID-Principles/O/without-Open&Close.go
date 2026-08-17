package main

import "fmt"

func Pay(method string, amount float64) {
	switch method {
	case "creditcard":
		fmt.Printf("Creditcard Payment of %.2f!!\n", amount)
	case "upi":
		fmt.Printf("UPI Payment of %.2f!!\n", amount)
	case "debit":
		fmt.Printf("Debit card payment of %.2f!!\n", amount)
	default:
		fmt.Println("No method used !!")
	}

}

func main() {

	Pay("creditcard", 20)
	Pay("upi", 30.9)
	Pay("debit", 50)
	Pay("apple", 40)

}
