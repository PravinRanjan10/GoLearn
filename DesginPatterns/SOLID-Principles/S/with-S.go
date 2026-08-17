package main

import (
	"fmt"
	"strings"
)

type NameValidator struct{}
type EmailValidator struct{}
type UsersDB struct {
	Users []string
}
type EmailService struct{}

func (n NameValidator) Validate(name string) error {
	// validate the user
	if name == "" || strings.Contains(name, "#") {
		fmt.Println("user name can not be empty or contains #")
		return fmt.Errorf("name validation failed!")
	}
	return nil
}
func (e EmailValidator) Validate(email string) error {
	// validate the email
	if email == "" || !strings.Contains(email, "@") {
		fmt.Println("email can not be empty")
		return fmt.Errorf("Email validation failed!")
	}
	return nil
}
func (u *UsersDB) Save(user string) {
	u.Users = append(u.Users, user)
}

func (e EmailService) SendEmail(name, email string) {
	fmt.Printf("Welcome:%s!. Email sent to: %s\n", name, email)
}

// This is orchestrator for managing all the services
type UserService struct {
	eValidator EmailValidator
	nValidator NameValidator
	users      *UsersDB
	email      EmailService
}

func (u UserService) Register(name, email string) error {
	err := u.nValidator.Validate(name)
	if err != nil {
		return err
	}

	err = u.eValidator.Validate(email)
	if err != nil {
		return err
	}
	u.users.Save(name)
	u.email.SendEmail(name, email)
	return nil

}

func main() {
	service := &UserService{
		nValidator: NameValidator{},
		eValidator: EmailValidator{},
		users:      &UsersDB{},
		email:      EmailService{},
	}

	err := service.Register("Pravin", "PravinRanjan10@gmail.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("All users:", service.users.Users)
}

/*
The first design violates SRP because one method is handling
validation, persistence, and notification.

The second design separates responsibilities:
- Validator → Validation only
- UserRepository → Storage only
- EmailService → Email only
- UserService → Orchestration only

Each component has exactly one reason to change,
making the code easier to test, maintain, and extend.
*/
