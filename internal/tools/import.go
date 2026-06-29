package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/voocel/ainovel-cli/internal/store"
)

type ImportTool struct {
	store *store.Store
}

func NewImportTool(store *store.Store) *ImportTool {
	return &ImportTool{store: store}
}

func (t *ImportTool) Name() string { return "novel_import" }

func (t *ImportTool) Execute(ctx context.Context, args []byte) ([]byte, error) {
	var a struct {
		InputPath string `json:"input_path"`
	}
	if err := json.Unmarshal(args, &a); err != nil {
		return nil, fmt.Errorf("failed to unmarshal args: %w", err)
	}

	content, err := os.ReadFile(a.InputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read input file: %w", err)
	}

	text := string(content)
	
	re := regexp.MustCompile(`(?m)^(Chương|Chapter)\s+\d+.*$`)
	indices := re.FindAllStringIndex(text, -1)

	if len(indices) == 0 {
		return nil, fmt.Errorf("no chapters found in the input file matching 'Chương X' or 'Chapter X'")
	}

	chaptersDir := filepath.Join(t.store.Dir(), "chapters")
	os.MkdirAll(chaptersDir, 0755)

	count := 0
	for i := 0; i < len(indices); i++ {
		start := indices[i][0]
		end := len(text)
		if i < len(indices)-1 {
			end = indices[i+1][0]
		}
		
		chapterText := strings.TrimSpace(text[start:end])
		chapterNum := i + 1
		
		filePath := filepath.Join(chaptersDir, fmt.Sprintf("%d.md", chapterNum))
		err := os.WriteFile(filePath, []byte(chapterText), 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to write chapter %d: %w", chapterNum, err)
		}
		count++
	}

	return []byte(fmt.Sprintf("Imported %d chapters successfully", count)), nil
}
