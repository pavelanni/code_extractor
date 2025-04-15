package extractor

import (
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

// ExtractCodeBlocks extracts code blocks from a Markdown file content.
// It returns a slice of strings containing the code blocks that match the specified language.
// If language is "all", it returns all code blocks regardless of language.
func ExtractCodeBlocks(content []byte, language string) ([]string, error) {
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
			if language == "all" || strings.EqualFold(lang, language) {
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
		return nil, err
	}

	return codeBlocks, nil
}
