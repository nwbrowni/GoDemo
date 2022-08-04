package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// account struct(s) and methods
type account struct {
	name     nombre
	address  addr
	birthday bday
	username string
	password string
}

func (a *account) SetUpNewAccount() error {
	scanner := bufio.NewScanner(os.Stdin)

	// all required data to set up an account
	var fname string
	var lname string
	var street string
	var apt string
	var city string
	var state string
	var zip string
	var day int
	var month int
	var year int

	var err error

	fmt.Println("Please enter the following name information")
	fmt.Print("first name: ")
	scanner.Scan()
	fname = scanner.Text()

	fmt.Print("last name: ")
	scanner.Scan()
	lname = scanner.Text()

	fmt.Println("Please enter the following address information")
	fmt.Print("street: ")
	scanner.Scan()
	street = scanner.Text()

	fmt.Print("apt: ")
	scanner.Scan()
	apt = scanner.Text()

	fmt.Print("city: ")
	scanner.Scan()
	city = scanner.Text()

	fmt.Print("state: ")
	scanner.Scan()
	state = scanner.Text()

	fmt.Print("zip: ")
	scanner.Scan()
	zip = scanner.Text()

	fmt.Println("Please enter the following birthday information")
	fmt.Print("day (integer): ")
	scanner.Scan()
	day, err = strconv.Atoi(scanner.Text()) // second value is an error

	if err != nil {
		return err
	}

	fmt.Print("month (integer): ")
	scanner.Scan()
	month, err = strconv.Atoi(scanner.Text())

	if err != nil {
		return err
	}

	fmt.Print("year (integer): ")
	scanner.Scan()
	year, err = strconv.Atoi(scanner.Text())

	if err != nil {
		return err
	}

	a.name.SetNombre(fname, lname)
	a.address.SetAddress(street, apt, city, state, zip)
	a.birthday.SetBirthday(day, month, year)

	fmt.Println("Please enter a username and password")
	fmt.Print("username: ")
	scanner.Scan()
	a.username = scanner.Text()

	fmt.Print("password: ")
	scanner.Scan()
	a.password = scanner.Text()

	return nil
}

func (a *account) GetAccountStrings() []string {
	slc := make([]string, 0)

	slc = append(slc, a.name.GetNombreStrings()...)
	slc = append(slc, a.address.GetAddressStrings()...)
	slc = append(slc, a.birthday.GetBirthdayStrings()...)
	slc = append(slc, a.username)
	slc = append(slc, a.password)

	return slc
}

// account list struct and methods
type accountlist struct {
	list []account
}

/*
To read a file line by line:
    file, err := os.Open("text.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
*/

func (al *accountlist) ReadAccountListFromFile(filepath string) {
	file, err := os.Open(filepath)

	var readready bool = false
	var temp account

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var str string = scanner.Text()

		if str == "START ACCOUNT" {
			readready = true
		}

		if !readready {
			continue
		}

		// read in and store the account name information
		scanner.Scan()
		temp.name.fname = scanner.Text()
		scanner.Scan()
		temp.name.lname = scanner.Text()

		// read in and store the account address information
		scanner.Scan()
		temp.address.street = scanner.Text()
		scanner.Scan()
		temp.address.apt = scanner.Text()
		scanner.Scan()
		temp.address.city = scanner.Text()
		scanner.Scan()
		temp.address.state = scanner.Text()
		scanner.Scan()
		temp.address.zip = scanner.Text()

		// read in and store the account birthday information
		scanner.Scan()
		temp.birthday.day, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		temp.birthday.month, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		temp.birthday.year, _ = strconv.Atoi(scanner.Text())

		// read in and store the account username and password
		scanner.Scan()
		temp.username = scanner.Text()
		scanner.Scan()
		temp.password = scanner.Text()

		al.list = append(al.list, temp)
		// reset flag for the start of new account
		readready = false

		scanner.Scan()
		str = scanner.Text()

		// if account not ended properly, complain
		if str != "END ACCOUNT" {
			fmt.Println("ERROR: should have been 'END ACCOUNT' line here")
			log.Fatal()
		}

		// need to read information in to create accounts
		// once the accounts are created, then add them to the account list
	}
}

func (al *accountlist) SaveAccountListToFile(filepath string) {
	file, err := os.Create(filepath)

	if err != nil {
		log.Fatal(err) // need to look into log some more
	}
	defer file.Close()

	for _, acc := range al.list {
		file.WriteString("START ACCOUNT\n") // the "\n" moves to the next line
		for _, str := range acc.GetAccountStrings() {
			_, err := file.WriteString(str + "\n")

			if err != nil {
				log.Fatal(err)
			}
		}
		file.WriteString("END ACCOUNT\n")
	}

}

/*
BEGIN ACCOUNT MENU FUNCTIONS
*/

func (a *account) AccountMainMenu() *account {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for input != "0" {
		if input < "0" || input > "9" {
			fmt.Println("INVALID INPUT: Input should be a value between 0 and 9")
			fmt.Println()
		} else {
			a.ProcessMainMenuInput(input)
		}

		fmt.Println("ACCOUNT MENU:")
		fmt.Println("0) Exit")
		fmt.Println("1) View Account Information")
		fmt.Println("2) Edit Account Information")
		fmt.Println()
		fmt.Print(">> ")

		scanner.Scan()
		input = scanner.Text()
	}

	fmt.Println("LEAVING ACCOUNT MENU...")
	fmt.Println()

	return a
}

func (a *account) ProcessMainMenuInput(s string) {
	// NOTE: don't need to precess 0 because that is already processed by the infinite loop
	if s == "1" {
		a.DisplayAccountInformation()
	} else if s == "2" {
		a.EditAccountInformationMenu()
	}
}

func (a *account) DisplayAccountInformation() {
	fmt.Printf("Name: %s %s\n", a.name.fname, a.name.lname)

	if a.address.apt == "" {
		fmt.Printf("Address: %s, %s, %s %s\n", a.address.street, a.address.city, a.address.state, a.address.zip)
	} else {
		fmt.Printf("Address: %s, %s, %s, %s %s\n", a.address.street, a.address.apt, a.address.city, a.address.state, a.address.zip)
	}

	fmt.Printf("Birthday: %d-%d-%d\n", a.birthday.day, a.birthday.month, a.birthday.year)

	// go through range and display information

	fmt.Println()
}

func (a *account) EditAccountInformationMenu() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for input != "0" {
		if input < "0" || input > "9" {
			fmt.Println("INVALID INPUT: Input should be a value between 0 and 9")
			fmt.Println()
		} else {
			a.ProcessEditAccountInformationMenuInput(input)
		}

		fmt.Println("EDIT ACCOUNT INFORMATION MENU:")
		fmt.Println("0) Exit")
		fmt.Println("1) Edit Name Menu")
		fmt.Println("2) Edit Address Menu")
		fmt.Println("3) Edit Birthday Menu")
		fmt.Println()
		fmt.Print(">> ")

		scanner.Scan()
		input = scanner.Text()
	}

	fmt.Println("LEAVING EDIT ACCOUNT INFORMATION WINDOW...")
	fmt.Println()
}

func (a *account) ProcessEditAccountInformationMenuInput(s string) {
	// NOTE: don't need to handle 0 because the infinite loop already does
	if s == "1" {
		a.EditAccountNameMenu()
	} else if s == "2" {
		a.EditAccountAddressMenu()
	} else if s == "3" {
		a.EditAccountBirthdayMenu()
	}
}

func (a *account) EditAccountNameMenu() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for input != "0" {
		if input < "0" || input > "9" {
			fmt.Println("INVALID INPUT: Input should be a value between 0 and 9")
			fmt.Println()
		} else {
			a.ProcessEditAccountNameMenuInput(input)
		}

		fmt.Println("EDIT ACCOUNT NAME MENU:")
		fmt.Println("0) Exit")
		fmt.Println("1) Edit First Name")
		fmt.Println("2) Edit Last Name")
		fmt.Println()
		fmt.Print(">> ")

		scanner.Scan()
		input = scanner.Text()
	}

	fmt.Println("LEAVING EDIT ACCOUNT NAME MENU...")
	fmt.Println()
}

func (a *account) ProcessEditAccountNameMenuInput(s string) {
	if s == "1" {
		a.EditAccountNameFName()
	} else if s == "2" {
		a.EditAccountNameLName()
	}
}

func (a *account) EditAccountNameFName() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter a new first name: ")
	scanner.Scan()
	a.name.fname = scanner.Text()

	fmt.Println()
}

func (a *account) EditAccountNameLName() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter a new last name: ")
	scanner.Scan()
	a.name.lname = scanner.Text()

	fmt.Println()
}

func (a *account) EditAccountAddressMenu() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for input != "0" {
		if input < "0" || input > "9" {
			fmt.Println("INVALID INPUT: Input should be a value between 0 and 9")
			fmt.Println()
		} else {
			a.ProcessEditAccountAddressMenuInput(input)
		}

		fmt.Println("EDIT ACCOUNT ADDRESS MENU:")
		fmt.Println("0) Exit")
		fmt.Println("1) Edit Street")
		fmt.Println("2) Edit Apartment")
		fmt.Println("3) Edit City")
		fmt.Println("4) Edit State")
		fmt.Println("5) Edit Zip")
		fmt.Println()
		fmt.Print(">> ")

		scanner.Scan()
		input = scanner.Text()
	}

	fmt.Println("LEAVING EDIT ACCOUNT ADDRESS MENU...")
	fmt.Println()
}

func (a *account) ProcessEditAccountAddressMenuInput(s string) {
	if s == "1" {
		a.EditAccountAddressStreet()
	} else if s == "2" {
		a.EditAccountAddressApt()
	} else if s == "3" {
		a.EditAccountAddressCity()
	} else if s == "4" {
		a.EditAccountAddressState()
	} else if s == "5" {
		a.EditAccountAddressZip()
	}
}

func (a *account) EditAccountAddressStreet() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter a new street: ")
	scanner.Scan()
	a.address.street = scanner.Text()

	fmt.Println()
}

func (a *account) EditAccountAddressApt() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter a new apt: ")
	scanner.Scan()
	a.address.apt = scanner.Text()

	fmt.Println()
}

func (a *account) EditAccountAddressCity() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter a new city: ")
	scanner.Scan()
	a.address.city = scanner.Text()

	fmt.Println()
}

func (a *account) EditAccountAddressState() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter a new state: ")
	scanner.Scan()
	a.address.state = scanner.Text()

	fmt.Println()
}

func (a *account) EditAccountAddressZip() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter a new zip: ")
	scanner.Scan()
	a.address.zip = scanner.Text()

	fmt.Println()
}

func (a *account) EditAccountBirthdayMenu() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for input != "0" {
		if input < "0" || input > "9" {
			fmt.Println("INVALID INPUT: Input should be a value between 0 and 9")
			fmt.Println()
		} else {
			a.ProcessEditAccountBirthdayMenuInput(input)
		}

		fmt.Println("EDIT ACCOUNT BIRTHDAY MENU:")
		fmt.Println("0) Exit")
		fmt.Println("1) Edit Day")
		fmt.Println("2) Edit Month")
		fmt.Println("3) Edit Year")
		fmt.Println()
		fmt.Print(">> ")

		scanner.Scan()
		input = scanner.Text()
	}

	fmt.Println("LEAVING EDIT ACCOUNT BIRTHDAY MENU...")
	fmt.Println()
}

func (a *account) ProcessEditAccountBirthdayMenuInput(s string) {
	if s == "1" {
		fmt.Printf("You selected: %s\n", s)
		// edit day
	} else if s == "2" {
		fmt.Printf("You selected: %s\n", s)
		// edit month
	} else if s == "3" {
		fmt.Printf("You selected: %s\n", s)
		// edit year
	}
}
