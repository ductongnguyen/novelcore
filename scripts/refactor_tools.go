package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filepath.WalkDir("internal/tools", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(path, ".go") {
			content, _ := os.ReadFile(path)
			text := string(content)

			// Remove imports
			text = strings.ReplaceAll(text, "\"github.com/voocel/agentcore/schema\"", "")
			text = strings.ReplaceAll(text, "\"github.com/voocel/agentcore/tools\"", "")

			// Comment out Schema() methods (hacky but works for now, we'll fix compile errors later)
			lines := strings.Split(text, "\n")
			inSchema := false
			for i, line := range lines {
				if strings.Contains(line, ") Schema() map[string]any") {
					inSchema = true
				}
				if inSchema {
					lines[i] = "// " + line
					if strings.HasPrefix(line, "}") {
						inSchema = false
					}
				}
			}

			os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
		}
		return nil
	})
}
