package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// GetUserInput Function to get user input
func GetUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// Remove trailing newline characters
	return strings.TrimSpace(input)
}

func SanitizeFileName(filename string) string {
	filename = strings.ReplaceAll(filename, " ", "_")
	filename = strings.ToLower(filename)
	return filename
}
