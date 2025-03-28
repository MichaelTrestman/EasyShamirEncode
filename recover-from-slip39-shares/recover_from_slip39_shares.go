package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gavincarr/go-slip39"
)

func main() {
	passphrase := os.Getenv("passphrase")
	if passphrase == "" {
		fmt.Println("Please set the 'passphrase' environment variable.")
		os.Exit(1)
	}

	shareDir := "shares/"
	files, err := ioutil.ReadDir(shareDir)
	if err != nil {
		fmt.Println("Failed to read shares directory:", err)
		os.Exit(1)
	}

	var shares []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			data, err := os.ReadFile(filepath.Join(shareDir, file.Name()))
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", file.Name(), err)
				os.Exit(1)
			}
			share := strings.TrimSpace(string(data))
			if share != "" {
				shares = append(shares, share)
			}
		}
	}

	if len(shares) == 0 {
		fmt.Println("No valid shares found in 'shares/' directory.")
		os.Exit(1)
	}

	recoveredSecret, err := slip39.CombineMnemonicsWithPassphrase(
		shares,
		[]byte(passphrase),
	)
	if err != nil {
		fmt.Println("Failed to recover secret:", err)
		os.Exit(1)
	}

	fmt.Println("Recovered secret bytes:")
	fmt.Println(recoveredSecret)

	fmt.Println("Recovered secret (hex):")
	fmt.Println(hex.EncodeToString(recoveredSecret))
}
