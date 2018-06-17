package main

import (
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"os"
	"sync"
)

var (
	passwords    []string
	walletData   []byte
	walletFile   string
	passwordFile string
	guard        chan struct{}
	passWait     *sync.WaitGroup
)

func main() {
	fmt.Println("=============================")
	fmt.Println("       GETH BREAKER")
	fmt.Println("=============================")
	fmt.Println("Usage:")
	fmt.Println("  gethbreaker wallet.json passwords.txt")

	if len(os.Args) < 2 {
		fmt.Println("Missing data! Usage command:")
		fmt.Println("  gethbreaker wallet.json passwords.txt")
		os.Exit(1)
	}

	walletFile = os.Args[1]
	passwordFile = os.Args[2]

	guard = make(chan struct{}, 16)
	passWait = new(sync.WaitGroup)

	fmt.Printf("Loading wallet file:    %v\n", walletFile)
	fmt.Printf("Loading passwords list: %v\n", passwordFile)
	fmt.Println("=============================")

	file, err := os.Open(passwordFile)
	if err != nil {
		fmt.Printf("Could not open file: %v\n", passwordFile)
		os.Exit(0)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passWait.Add(1)
		passwords = append(passwords, scanner.Text())
	}

	walletData, err = ioutil.ReadFile(walletFile)
	if err != nil {
		panic(err)
	}

	for _, v := range passwords {
		guard <- struct{}{}
		pass := v
		go TryPassword(pass)
	}

	passWait.Wait()
	fmt.Println("Process is complete")

}

func TryPassword(password string) {
	fmt.Printf("Trying: %v\n", password)
	key, err := keystore.DecryptKey(walletData, password)
	if err == nil {
		fmt.Printf("\nFound password!!!\n")
		fmt.Printf("File:     %v\n", walletFile)
		fmt.Printf("Address:  %v\n", key.Address.String())
		fmt.Printf("Password: %v\n", password)
		os.Exit(0)
	}
	<-guard
}
