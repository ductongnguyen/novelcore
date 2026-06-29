package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/voocel/ainovel-cli/internal/store"
)

type ExportTool struct {
	store *store.Store
}

func NewExportTool(store *store.Store) *ExportTool {
	return &ExportTool{store: store}
}

func (t *ExportTool) Name() string { return "novel_export" }

func (t *ExportTool) Execute(ctx context.Context, args []byte) ([]byte, error) {
	var a struct {
		OutputPath string `json:"output_path"`
	}
	if err := json.Unmarshal(args, &a); err != nil {
		return nil, fmt.Errorf("failed to unmarshal args: %w", err)
	}

	if a.OutputPath == "" {
		a.OutputPath = "exported_novel.txt"
	}

	chaptersDir := filepath.Join(t.store.Dir(), "chapters")
	files, err := os.ReadDir(chaptersDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read chapters dir: %w", err)
	}

	var output strings.Builder
	
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".md") {
			content, err := os.ReadFile(filepath.Join(chaptersDir, f.Name()))
			if err != nil {
				continue
			}
			output.Write(content)
			output.WriteString("\n\n")
		}
	}

	err = os.WriteFile(a.OutputPath, []byte(output.String()), 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to write export file: %w", err)
	}

	return []byte(fmt.Sprintf("Novel exported successfully to %s", a.OutputPath)), nil
}
