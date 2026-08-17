package main

import "fmt"

type User struct {
	Name string
	Age  int
	ID   int
}
type UserRepository interface {
	Save(User) error
	Get(int) (User, error)
}

// --------------------------------------------------
// MySQL implementation
// --------------------------------------------------

type MySQLRepository struct {
	users map[int]User
}

func NewMySQLRepository() *MySQLRepository {
	return &MySQLRepository{
		users: make(map[int]User),
	}
}

func (r *MySQLRepository) Save(user User) error {
	r.users[user.ID] = user

	fmt.Println("User saved in MySQL")

	return nil
}

func (r *MySQLRepository) Get(id int) (User, error) {
	user, ok := r.users[id]
	if !ok {
		return User{}, fmt.Errorf("user %d not found", id)
	}

	fmt.Println("User fetched from MySQL")

	return user, nil
}

// --------------------------------------------------
// Mock implementation
// --------------------------------------------------

type MockRepository struct {
	users map[int]User
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		users: make(map[int]User),
	}
}

func (r *MockRepository) Save(user User) error {
	r.users[user.ID] = user

	fmt.Println("User saved in Mock Repository")

	return nil
}

func (r *MockRepository) Get(id int) (User, error) {
	user, ok := r.users[id]
	if !ok {
		return User{}, fmt.Errorf("user %d not found", id)
	}

	fmt.Println("User fetched from Mock Repository")

	return user, nil
}

// --------------------------------------------------
// Business logic
// --------------------------------------------------

func CreateUser(repo UserRepository, user User) error {
	return repo.Save(user)
}
func GetUser(repo UserRepository, id int) (User, error) {
	return repo.Get(id)
}

func main() {

	user := User{
		Name: "Pravin",
		Age:  30,
		ID:   123,
	}

	// Production implementation
	mysqlRepo := NewMySQLRepository()
	err := CreateUser(mysqlRepo, user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result, err := GetUser(mysqlRepo, 123)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("User:", result)

	// Mock implementation
	mockRepo := NewMockRepository()
	err = CreateUser(mockRepo, user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result, err = GetUser(mockRepo, 123)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Print("Mock user:", result)

}

/*
Where is LSP here?

Look at this function:

func CreateUser(repo UserRepository, user User) error {
	return repo.Save(user)
}

It doesn't care whether repo is:

MySQLRepository

or:

MockRepository

Both can be substituted:

CreateUser(mysqlRepo, user)

or:

CreateUser(mockRepo, user)

without changing CreateUser().

That's Liskov Substitution Principle.
*/
