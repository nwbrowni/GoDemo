package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// app structure and methods
type app struct {
	acclist accountlist
}

/*
BEGIN STARTUP FUNCTIONS
*/

func (a *app) Startup() (account, error) {
	var logincounter int = 0

	fmt.Println("Welcome to THE APP!")

	acc, err := AccountLogIn(&a.acclist)

	if err != nil && logincounter < 3 {
		logincounter++
		acc, err = AccountLogIn(&a.acclist)
	}

	if logincounter == 3 {
		err = errors.New("too many unsuccessful login attempts")
	}

	return acc, err
}

func AccountLogIn(al *accountlist) (account, error) {
	var acc account
	var err error = nil

	haveaccount := DoYouHaveAnExistingAccount()

	if !haveaccount {
		createaccount := WouldYouLikeToCreateAnAccount()

		if !createaccount {
			err = errors.New("no account desired")
			return acc, err
		}

		err = AddAccount(al)

		if err != nil {
			return acc, err
		}
	}

	user, pass := GetUsernameAndPassword()
	acc, err = AccessAccount(al, user, pass)

	return acc, err
}

func DoYouHaveAnExistingAccount() bool {
	scanner := bufio.NewScanner(os.Stdin)
	var response string

	for response == "" {
		fmt.Print("Do you have an existing account on THE APP (y/n)? ")
		scanner.Scan()
		response = scanner.Text()

		// if BOTH are not true, then response is invalid
		if response != "y" && response != "n" {
			response = ""
			fmt.Println("Sorry. That was an invalid input.")
		}
	}

	if response == "y" {
		return true
	} else {
		return false
	}
}

func WouldYouLikeToCreateAnAccount() bool {
	scanner := bufio.NewScanner(os.Stdin)
	var response string

	for response == "" {
		fmt.Print("Would you like to create an account on THE APP (y/n)? ")
		scanner.Scan()
		response = scanner.Text()

		// if BOTH are not true, then response is invalid
		if response != "y" && response != "n" {
			response = ""
			fmt.Println("Sorry. That was an invalid input.")
		}
	}

	if response == "y" {
		return true
	} else {
		return false
	}
}

func AddAccount(al *accountlist) error {
	// create a new account
	var newaccount account
	var i int
	var err error = nil
	var found bool = false

	newaccount.SetUpNewAccount()

	for i = 0; i < len(al.list); i++ {
		if newaccount.username == al.list[i].username {
			fmt.Println("This account cannot be created...")
			fmt.Printf("An account with the username %s already exists\n", newaccount.username)
			found = true
			break
		}
	}

	// only append if the new account doesn't share a username with an existing account
	if !found {
		fmt.Println("Adding new account...")
		al.list = append(al.list, newaccount)
		fmt.Println("New account has been added.")
	} else {
		err = errors.New("account already exists")
	}

	return err
}

func GetUsernameAndPassword() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	var user string
	var pass string

	fmt.Println("Please enter the following information")

	// get the username
	fmt.Print("Username: ")
	scanner.Scan()
	user = scanner.Text()

	// get the password
	fmt.Print("Password: ")
	scanner.Scan()
	pass = scanner.Text()

	return user, pass
}

func AccessAccount(al *accountlist, username string, password string) (account, error) {
	var found bool
	var acc account
	var err error = nil

	// see if username and password combination is valid, loop through elements using range
	for _, element := range al.list {
		if element.username == username && element.password == password {
			acc = element
			found = true
			break
		}
	}

	if found {
		fmt.Println("Found Account...")
	} else {
		fmt.Println("Account was not found...")
		err = errors.New("account not found")
	}

	return acc, err
}

/*
END STARTUP FUNCTIONS
*/
