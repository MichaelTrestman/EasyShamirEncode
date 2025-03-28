package main

import (
	// "encoding/hex"
    "strconv"
	"fmt"
	"os"
	"github.com/gavincarr/go-slip39"
)

func main() {
	// masterSecret="bleep blooop"
	// passphrase="asdf"
	// threshold="2"
	// total_shares="3"
	// ./create-slip39-shares

	masterSecret := os.Getenv("secret")
	passphrase := os.Getenv("passphrase")
fmt.Println("os.Getenv(\"secret\")")
	fmt.Println(os.Getenv("secret"))
	fmt.Println("os.Getenv(\"threshold\")")
	fmt.Println(os.Getenv("threshold"))
	threshold_s := os.Getenv("threshold")
	threshold, err := strconv.Atoi(threshold_s)
	if err != nil {
		fmt.Println("Invalid threshold:", err)
				fmt.Println("Usage: shit create-slip39-shares <secret> <passphrase> <minimum_shares> <total_shares>")
		os.Exit(1)
	}

	total_shares_s := os.Getenv("total_shares")
	total_shares, err := strconv.Atoi(total_shares_s)
	if err != nil {
		fmt.Println("Invalid number of signatories:", err)
				fmt.Println("Usage: ploo create-slip39-shares <secret> <passphrase> <minimum_shares> <total_shares>")
		os.Exit(1)
	}
	// Generate a single group of 3 of 5 shares for masterSecret
	// masterSecretBytes, err := hex.DecodeString(masterSecret)
	// masterSecretBytes, err := hex.DecodeString(masterSecret)
	masterSecretBytes := []byte(masterSecret)
	fmt.Println("len(masterSecretBytes)")
	fmt.Println(len(masterSecretBytes))
	groupCount := 1
	memberGroupParams := []slip39.MemberGroupParameters{{threshold, total_shares}}

	groups, err := slip39.GenerateMnemonicsWithPassphrase(
		groupCount,
		memberGroupParams,
		masterSecretBytes,
		[]byte(passphrase),
	)
		if err != nil {
		fmt.Println("EEEEERRRRROR:", err)
				fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(groups)
	// Output: 5

	// // Combine 3 of the 5 shares to recover the master secret
	// shares := []string{groups[0][0], groups[0][2], groups[0][4]}
	// recoveredSecret, _ := slip39.CombineMnemonicsWithPassphrase(
	// 	shares,
	// 	[]byte(passphrase),
	// )
	// fmt.Println(hex.EncodeToString(recoveredSecret))
	// // Output: bb54aac4b89dc868ba37d9cc21b2cece
}