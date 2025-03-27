package main

import (
	"bufio"
	"fmt"
	"os"

	sssa "github.com/SSSaaS/sssa-golang"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: combine_shares <shares_file>")
		os.Exit(1)
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open file %s: %v\n", fileName, err)
		os.Exit(1)
	}
	defer file.Close()

	var shares []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		share := scanner.Text()
		if sssa.IsValidShare(share) {
			shares = append(shares, share)
		} else {
			fmt.Printf("Invalid share skipped: %s\n", share)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(shares) == 0 {
		fmt.Println("No valid shares found.")
		os.Exit(1)
	}
	
	secret, err := sssa.Combine(shares)
	if err != nil {
		fmt.Println("Error combining shares:", err)
		os.Exit(1)
	}

	fmt.Println("Recovered secret:", secret)
}
