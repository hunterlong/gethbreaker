package main

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
)

var (
	passwords []string
)

func main() {

	wallet := "wallet.json"
	passwordFile := "passwords.txt"

	file, err := os.Open(passwordFile)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}

	data, err := ioutil.ReadFile(wallet)
	if err != nil {
		panic(err)
	}

	for k, pass := range passwords {

		fmt.Printf("%v | trying: %v\n", k, pass)

		key, err := keystore.DecryptKey(data, pass)
		if err != nil {

		} else {
			fmt.Printf("\nFound password for %v!\nAddress: %v\nPassword: %v\n", wallet, key.Address.String(), pass)
			os.Exit(1)
		}

	}


}
