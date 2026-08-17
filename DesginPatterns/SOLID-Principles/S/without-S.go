package main

import (
	"fmt"
	"strings"
)

type UserManager struct {
	users  []string
	emails []string
}

func (u *UserManager) Register(name, email string) error {
	// validate the user
	if name == "" || strings.Contains(name, "#") {
		fmt.Println("user name can not be empty or contains #")
		return fmt.Errorf("name validation failed!")
	}

	// validate the email
	if email == "" || !strings.Contains(email, "@") {
		fmt.Println("email can not be empty")
		return fmt.Errorf("Email validation failed!")
	}

	// save to DB
	u.users = append(u.users, name)
	u.emails = append(u.emails, email)

	// send welcome email
	fmt.Printf("Welcome:%s!. Email sent to: %s\n", name, email)
	return nil

}

func main() {
	//users := new(UserManager)
	manager := &UserManager{}
	err := manager.Register("Pravin", "pravinranjan10@gmail.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("All users:", manager.users)
}

/*
Why SRP is voilated?

Register() is responsible for:
	* Validation
	* Data storage
	* Email sending
if name/email validation logic changes required, this function needs to be changed.
if storage changes, this function changes
if email sending logic changes, this function changes

So, multiple reasons to change = SRP violation.
*/
