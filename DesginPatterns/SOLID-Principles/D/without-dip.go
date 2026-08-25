package main

import "fmt"

type OrderService struct {
	db *MySQLDB
}

type MySQLDB struct{}

func (db *MySQLDB) SaveOrder(order string) error {
	fmt.Println("Saving order to MySQL:", order)
	return nil
}

func (s *OrderService) CreateOrder(order string) error {
	fmt.Println("Creating order...")
	return s.db.SaveOrder(order)
}

func main() {

	db := &MySQLDB{}

	service := &OrderService{
		db: db,
	}

	service.CreateOrder("Order-101")

}

/*
Dependency Inversion Principle (DIP) in Go

DIP says:

High-level modules should not depend directly on low-level modules. Both should depend on abstractions.

In Go, this is commonly achieved using interfaces + dependency injection.

Let's take a complete example.

1. Problem without DIP

Suppose we are building an order service.

We have:

OrderService
     ↓
   MySQL
Low-level module
type MySQLDB struct{}

func (db *MySQLDB) SaveOrder(order string) error {
	fmt.Println("Saving order to MySQL:", order)
	return nil
}
High-level module
type OrderService struct {
	db *MySQLDB
}

func (s *OrderService) CreateOrder(order string) error {
	fmt.Println("Creating order...")
	return s.db.SaveOrder(order)
}

Usage:

func main() {
	db := &MySQLDB{}

	service := &OrderService{
		db: db,
	}

	service.CreateOrder("Order-101")
}

This works, but there is a problem.

OrderService is tightly coupled to:

MySQLDB

If tomorrow we want PostgreSQL:

OrderService
     ↓
   MySQLDB    ❌

we have to modify OrderService.

*/
