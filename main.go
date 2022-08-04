/*
Important note:
- to create a main.exe in go
	- need package to be main
	- need a function called main with no parameters or returns
	- after all of this run
		- go mod init main
		- go build
	- and then use go mod tidy, go clean, and go build
*/

package main

import "fmt"

func main() {
	var THEAPP app // create an app instance
	THEAPP.acclist.ReadAccountListFromFile("accounts.txt")

	for _, acc := range THEAPP.acclist.list {
		fmt.Println(acc.GetAccountStrings())
	}

	client, err := THEAPP.Startup() // run the startup function

	if err != nil {
		return
	}

	client.AccountMainMenu()

	fmt.Println(client.GetAccountStrings())

	THEAPP.acclist.SaveAccountListToFile("accounts.txt")
}
