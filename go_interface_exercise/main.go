package main

import (
	"errors"
	"fmt"
)

type Account struct {
	firstName, lastName string
}

func (acc *Account) changeName(newName string) {
	acc.firstName = newName
}

func (acc *Account) String() string {
	return fmt.Sprintf("Employee name: %v %v", acc.firstName, acc.lastName)
}

type Employee struct {
	Account
	credit float64
}

func createEmployee(fName string, lName string, creds float64) *Employee {
	return &Employee{Account{firstName: fName, lastName: lName}, creds}
}

func (e *Employee) addCredit(amt float64) (float64, error) {
	if amt < 0.00 {
		return 0, errors.New("cannot credit zero amount")
	}
	return e.credit + amt, nil
}

func (e *Employee) removeCredit(amt float64) (float64, error) {
	if amt <= 0.0 {
		if e.credit < amt {
			return 0, errors.New("insufficient credit to deduct from ")

		}
		return 0, errors.New("invalid amount")

	}

	return e.credit - amt, nil
}

func (e *Employee) checkCredit() float64 {
	return e.credit
}

func main() {
	bruce := createEmployee("Bruce", "Lee", 500)
	fmt.Println(bruce)
	fmt.Println("Current Credit Balance: ", bruce.checkCredit())
	credits, err := bruce.addCredit(250)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}

	fmt.Println("New Credit Balance", bruce.checkCredit())

	credits, err = bruce.removeCredit(0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}

	bruce.changeName("Mark")

	fmt.Println(bruce)

}
