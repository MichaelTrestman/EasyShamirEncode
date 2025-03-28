package main

import (
	"fmt"
	"os"
	"strconv"
	"path/filepath"
	"github.com/gavincarr/go-slip39"
)

func main() {
	// Get env vars
	masterSecret := os.Getenv("secret")
	passphrase := os.Getenv("passphrase")

	threshold_s := os.Getenv("threshold")
	threshold, err := strconv.Atoi(threshold_s)
	if err != nil {
		fmt.Println("Invalid threshold:", err)
		os.Exit(1)
	}

	total_shares_s := os.Getenv("total_shares")
	total_shares, err := strconv.Atoi(total_shares_s)
	if err != nil {
		fmt.Println("Invalid total_shares:", err)
		os.Exit(1)
	}

	masterSecretBytes := []byte(masterSecret)

	// Generate shares
	groupCount := 1
	memberGroupParams := []slip39.MemberGroupParameters{{threshold, total_shares}}

	groups, err := slip39.GenerateMnemonicsWithPassphrase(
		groupCount,
		memberGroupParams,
		masterSecretBytes,
		[]byte(passphrase),
	)
	if err != nil {
		fmt.Println("Error generating mnemonics:", err)
		os.Exit(1)
	}

	// Create shares directory
	shareDir := "shares"
	err = os.MkdirAll(shareDir, 0755)
	if err != nil {
		fmt.Println("Failed to create shares directory:", err)
		os.Exit(1)
	}

	// Write each share to a separate file
	for i, share := range groups[0] {
		filename := filepath.Join(shareDir, fmt.Sprintf("%d.txt", i+1))
		err := os.WriteFile(filename, []byte(share+"\n"), 0644)
		if err != nil {
			fmt.Printf("Failed to write share to %s: %v\n", filename, err)
			os.Exit(1)
		}
		fmt.Printf("Wrote share %d to %s\n", i+1, filename)
	}
}
