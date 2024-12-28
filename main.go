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

func main() {
	// Load networks
	if err := loadNetworks("networks.txt"); err != nil {
		panic(fmt.Errorf("Failed to load networks: %v", err))
	}

	// Out
	fmt.Println("Loaded networks:")
	for _, n := range networks {
		fmt.Printf("%s %s\n", n.Name, n.URL)
	}
}
