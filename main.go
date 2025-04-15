package main

import (
	"fmt"
	"os"

	"code_extractor/pkg/extractor"

	"github.com/spf13/pflag"
)

func main() {
	// Define command line flags
	language := pflag.StringP("language", "l", "shell", "Language to extract code blocks for (default: shell)")
	pflag.Parse()

	// Check if a file path was provided
	if pflag.NArg() == 0 {
		fmt.Println("Error: Please provide a Markdown file path")
		os.Exit(1)
	}

	filePath := pflag.Arg(0)

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Extract code blocks
	codeBlocks, err := extractor.ExtractCodeBlocks(content, *language)
	if err != nil {
		fmt.Printf("Error extracting code blocks: %v\n", err)
		os.Exit(1)
	}

	// Print the extracted code blocks
	for _, block := range codeBlocks {
		fmt.Println(block)
	}
}
