package extractor

import (
	"testing"
)

func TestExtractCodeBlocks(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		language string
		want     []string
		wantErr  bool
	}{
		{
			name:     "extract shell code blocks",
			content:  "# Test\n```shell\necho \"hello\"\n```\n",
			language: "shell",
			want:     []string{"echo \"hello\"\n"},
			wantErr:  false,
		},
		{
			name:     "extract all code blocks",
			content:  "# Test\n```shell\necho \"hello\"\n```\n```python\nprint(\"world\")\n```\n",
			language: "all",
			want:     []string{"echo \"hello\"\n", "print(\"world\")\n"},
			wantErr:  false,
		},
		{
			name:     "no matching language",
			content:  "# Test\n```shell\necho \"hello\"\n```\n",
			language: "python",
			want:     []string{},
			wantErr:  false,
		},
		{
			name:     "empty content",
			content:  "",
			language: "shell",
			want:     []string{},
			wantErr:  false,
		},
		{
			name:     "code block without language",
			content:  "# Test\n```\necho \"hello\"\n```\n",
			language: "shell",
			want:     []string{},
			wantErr:  false,
		},
		{
			name:     "code block with comments",
			content:  "# Test\n```shell\n# This is a comment\necho \"hello\"\n# Another comment\n```\n",
			language: "shell",
			want:     []string{"# This is a comment\necho \"hello\"\n# Another comment\n"},
			wantErr:  false,
		},
		{
			name:     "multiple code blocks with same language",
			content:  "# Test\n```shell\necho \"first\"\n```\n```shell\necho \"second\"\n```\n",
			language: "shell",
			want:     []string{"echo \"first\"\n", "echo \"second\"\n"},
			wantErr:  false,
		},
		{
			name:     "code block with empty content",
			content:  "# Test\n```shell\n```\n",
			language: "shell",
			want:     []string{""},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractCodeBlocks([]byte(tt.content), tt.language)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractCodeBlocks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("ExtractCodeBlocks() got %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("ExtractCodeBlocks() got[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
