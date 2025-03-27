package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/SSSaaS/sssa-golang"

)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run create_shares.go <secret> <minimum_shares> <total_shares>")
		os.Exit(1)
	}

	secret := os.Args[1]

	minShares, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid minimum_shares:", err)
		os.Exit(1)
	}

	totalShares, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid total_shares:", err)
		os.Exit(1)
	}

	shares, err := sssa.Create(minShares, totalShares, secret)
	for _, share := range shares {
		fmt.Println(share)
	}
}
