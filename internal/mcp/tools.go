package mcp

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/voocel/ainovel-cli/internal/store"
	"github.com/voocel/ainovel-cli/internal/tools"
)

// RegisterTools registers all ainovel tools to the MCP server
func RegisterTools(s *server.MCPServer, store *store.Store) {
	s.AddTool(mcp.NewTool("novel.status",
		mcp.WithDescription("Get the current status of the novel project"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Basic status reporting
		state, _ := store.Progress.Load()
		var phase, flow string
		if state != nil {
			phase = string(state.Phase)
			flow = string(state.Flow)
		}
		status := map[string]any{
			"phase": phase,
			"flow":  flow,
		}
		data, _ := json.Marshal(status)
		return mcp.NewToolResultText(string(data)), nil
	})

	s.AddTool(mcp.NewTool("novel.plan_story",
		mcp.WithDescription("Save the foundational elements of the story (premise, world rules, characters)"),
		mcp.WithString("premise", mcp.Required(), mcp.Description("The story premise")),
		mcp.WithString("world_rules", mcp.Description("World rules")),
		mcp.WithString("characters", mcp.Description("Character profiles")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Note: we can map this to tools.SaveFoundation
		// Since we don't have the exact exact JSON args for tools.SaveFoundation mapped out here, we simulate it.
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewSaveFoundationTool(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})

	s.AddTool(mcp.NewTool("novel.write_chapter",
		mcp.WithDescription("Write or commit a chapter"),
		mcp.WithNumber("chapter", mcp.Required(), mcp.Description("Chapter number")),
		mcp.WithString("content", mcp.Required(), mcp.Description("Chapter content")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewDraftChapterTool(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})

	// Other tools can be added similarly: load_project, review, export, resume
}
