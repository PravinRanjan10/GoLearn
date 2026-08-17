package main

import "fmt"

type Payment interface {
	Pay(float64) error
}

type creditcard struct{}
type upi struct{}
type debit struct{}

func (c creditcard) Pay(amount float64) error {
	fmt.Printf("Creditcard Payment of %.2f!!\n", amount)
	return nil
}

func (u upi) Pay(amount float64) error {
	fmt.Printf("UPI Payment of %.2f!!\n", amount)
	return nil
}

func (d debit) Pay(amount float64) error {
	fmt.Printf("Debit card Payment of %.2f!!\n", amount)
	return nil
}

// ProcessPayment depends on the interface,
// not on CreditCard, UPI, or PayPal.
func ProcessPayment(p Payment, amount float64) error {
	return p.Pay(amount)
}

func main() {
	credit := &creditcard{}
	upi := &upi{}
	debit := &debit{}

	ProcessPayment(credit, 20.0)
	ProcessPayment(upi, 50)
	ProcessPayment(debit, 60)

	/* Suppose the requirement changes:
	Add Apple Pay.
	We don't modify ProcessPayment().
	applePay := ApplePay{}
	ProcessPayment(applePay, 400)
	*/

}

/*
If the interviewer asks "Explain Open/Closed Principle with a Go example", you can say:

"OCP means a component should be open for extension but closed for modification. For example, in a payment system, instead of having a switch statement for Credit Card, UPI, and PayPal, I define a Payment interface. Each payment method implements the interface, and ProcessPayment depends only on that interface. If I need to add Apple Pay, I create a new implementation without modifying the existing payment-processing logic."

The important code to remember
type Payment interface {
	Pay(amount float64) error
}


func ProcessPayment(p Payment, amount float64) error {
	return p.Pay(amount)
}

New payment method → new implementation → no change to existing processing logic.

That's the core idea of Open/Closed Principle in Go.
*/
