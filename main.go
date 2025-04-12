package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
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

	// Create a new Markdown parser
	md := goldmark.New(
		goldmark.WithExtensions(),
	)

	// Parse the Markdown content
	doc := md.Parser().Parse(text.NewReader(content))

	// Walk through the AST and extract code blocks
	var codeBlocks []string
	if err := ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if fencedCodeBlock, ok := n.(*ast.FencedCodeBlock); ok {
			// Get the language specified in the code block
			lang := string(fencedCodeBlock.Language(content))

			// If language flag is set to "all" or matches the code block language
			if *language == "all" || strings.EqualFold(lang, *language) {
				// Extract the code content
				lines := fencedCodeBlock.Lines()
				var codeContent strings.Builder
				for i := 0; i < lines.Len(); i++ {
					line := lines.At(i)
					codeContent.Write(line.Value(content))
				}
				codeBlocks = append(codeBlocks, codeContent.String())
			}
		}
		return ast.WalkContinue, nil
	}); err != nil {
		fmt.Printf("Error parsing markdown: %v\n", err)
		os.Exit(1)
	}

	// Print the extracted code blocks
	for _, block := range codeBlocks {
		fmt.Println(block)
	}
}
