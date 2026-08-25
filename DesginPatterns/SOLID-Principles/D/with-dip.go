package main

import "fmt"

// Abstraction
type OrderRepository interface {
	SaveOrder(order string) error
}

// High-level module
type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) CreateOrder(order string) error {
	fmt.Println("Creating order...")
	return s.repo.SaveOrder(order)
}

// Low-level module
type MySQLDB struct{}

func (db *MySQLDB) SaveOrder(order string) error {
	fmt.Println("Saving order to MySQL:", order)
	return nil
}

// Another low-level implementation
type PostgreSQLDB struct{}

func (db *PostgreSQLDB) SaveOrder(order string) error {
	fmt.Println("Saving order to PostgreSQL:", order)
	return nil
}

func main() {

	// MySQL
	mysql := &MySQLDB{}

	service := NewOrderService(mysql)

	service.CreateOrder("Order-101")

	// PostgreSQL
	postgres := &PostgreSQLDB{}

	service = NewOrderService(postgres)

	service.CreateOrder("Order-102")
}
