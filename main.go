package main

import (
	"bufio"
	//"context"
	"fmt"
	//"math/big"
	//"net/http"
	"os"
	"strings"
	//"sync"
	//"time"
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/ethclient"
)

type Network struct {
	Name string
	URL  string
}

type Address struct {
	Name    string
	Address string
	Network string
}

var (
	networks  = make(map[string]*Network)
	addresses []*Address
)

// loadNetworks loads network configurations from the specified file into the global `networks` map.
// Each line in the file must follow the format `name=url`.
// Returns an error if the file cannot be opened or if there's an issue reading its content.
func loadNetworks(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Error opening networks file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "=")
		if len(parts) != 2 {
			continue
		}

		name, url := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		networks[name] = &Network{
			Name: name,
			URL:  url,
		}
	}

	return scanner.Err()
}

func loadAddresses(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Error opening addresses file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) != 3 {
			continue
		}

		name := strings.TrimSpace(parts[0])
		address := strings.TrimSpace(parts[1])
		network := strings.TrimSpace(parts[2])

		// TODO: Check address format

		if _, exists := networks[network]; !exists {
			return fmt.Errorf("Network %s not found for address %s", network, address)
		}

		addresses = append(addresses, &Address{
			Name:    name,
			Address: address,
			Network: network,
		})
	}

	return scanner.Err()
}

func main() {
	// Load networks
	if err := loadNetworks("networks.txt"); err != nil {
		panic(fmt.Errorf("Failed to load networks: %v", err))
	}

	// Load addresses
	if err := loadAddresses("addresses.txt"); err != nil {
		panic(fmt.Errorf("Failed to load addresses: %v", err))
	}

	// Out
	fmt.Println("Loaded networks:")
	for _, n := range networks {
		fmt.Printf("%s %s\n", n.Name, n.URL)
	}

	//
	fmt.Println("Loaded addresses:")
	for _, a := range addresses {
		fmt.Printf("%s %s %s\n", a.Name, a.Address, a.Network)
	}
}
