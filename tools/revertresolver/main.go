package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func extractRevertStrings(line string) []string {
	re := regexp.MustCompile(`\brevert\s+([\w\d_]+)\s*\(\s*\)`)
	matches := re.FindAllStringSubmatch(line, -1)
	var result []string
	for _, match := range matches {
		result = append(result, match[1])
	}
	return result
}

func keccak256Hash(input string) string {
	hash := crypto.Keccak256Hash([]byte(input + "()"))
	return hash.Hex()
}

func processFile(file string, filter string) {
	//nolint: gosec
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", file, err)
		return
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		revertStrings := extractRevertStrings(line)
		for _, revertString := range revertStrings {
			hashedString := keccak256Hash(revertString)
			if filter == "" || strings.HasPrefix(hashedString, filter) {
				fmt.Printf("%s(): %s (File: %s, Line: %d)\n", revertString, hashedString, file, i+1)
			}
		}
	}
}

func findAndPrintReverts(path, filter string) {
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(filePath) == ".sol" {
			processFile(filePath, filter)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
	}
}

func main() {
	var path string
	var filter string
	flag.StringVar(&path, "p", "", "path to search recursively")
	flag.StringVar(&filter, "f", "", "revert hash filter [optional]")
	flag.Parse()
	if len(path) == 0 {
		panic("expected path to be set (use -p flag)")
	}

	findAndPrintReverts(path, filter)
}
