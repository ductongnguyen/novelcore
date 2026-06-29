package mcp

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/voocel/ainovel-cli/internal/store"
	"github.com/voocel/ainovel-cli/internal/tools"
)

func RegisterTools(s *server.MCPServer, store *store.Store) {
	s.AddTool(mcp.NewTool("novel.status",
		mcp.WithDescription("Get the current status of the novel project"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

	s.AddTool(mcp.NewTool("novel.check_consistency",
		mcp.WithDescription("Executes novel.check_consistency"),
		mcp.WithNumber("chapter", mcp.Description("chapter")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewCheckConsistencyTool(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})

	s.AddTool(mcp.NewTool("novel.commit_chapter",
		mcp.WithDescription("Executes novel.commit_chapter"),
		mcp.WithNumber("chapter", mcp.Description("chapter")),
		mcp.WithString("summary", mcp.Description("summary")),
		mcp.WithString("characters", mcp.Description("characters")),
		mcp.WithString("key_events", mcp.Description("key_events")),
		mcp.WithString("timeline_events", mcp.Description("timeline_events")),
		mcp.WithString("foreshadow_updates", mcp.Description("foreshadow_updates")),
		mcp.WithString("relationship_changes", mcp.Description("relationship_changes")),
		mcp.WithString("state_changes", mcp.Description("state_changes")),
		mcp.WithString("cast_intros", mcp.Description("cast_intros")),
		mcp.WithString("hook_type", mcp.Description("hook_type")),
		mcp.WithString("dominant_strand", mcp.Description("dominant_strand")),
		mcp.WithString("feedback", mcp.Description("feedback")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewCommitChapterTool(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})

	s.AddTool(mcp.NewTool("novel.draft_chapter",
		mcp.WithDescription("Executes novel.draft_chapter"),
		mcp.WithNumber("chapter", mcp.Description("chapter")),
		mcp.WithString("content", mcp.Description("content")),
		mcp.WithString("mode", mcp.Description("mode")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewDraftChapterTool(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})

	s.AddTool(mcp.NewTool("novel.edit_chapter",
		mcp.WithDescription("Executes novel.edit_chapter"),
		mcp.WithNumber("chapter", mcp.Description("chapter")),
		mcp.WithString("old_string", mcp.Description("old_string")),
		mcp.WithString("new_string", mcp.Description("new_string")),
		mcp.WithBoolean("replace_all", mcp.Description("replace_all")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewEditChapterTool(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})

	s.AddTool(mcp.NewTool("novel.novel_context",
		mcp.WithDescription("Executes novel.novel_context"),
		mcp.WithNumber("chapter", mcp.Description("chapter")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewContextTool(store, tools.References{}, "")
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})

	s.AddTool(mcp.NewTool("novel.plan_chapter",
		mcp.WithDescription("Executes novel.plan_chapter"),
		mcp.WithNumber("chapter", mcp.Description("chapter")),
		mcp.WithString("title", mcp.Description("title")),
		mcp.WithString("goal", mcp.Description("goal")),
		mcp.WithString("conflict", mcp.Description("conflict")),
		mcp.WithString("hook", mcp.Description("hook")),
		mcp.WithString("emotion_arc", mcp.Description("emotion_arc")),
		mcp.WithString("notes", mcp.Description("notes")),
		mcp.WithString("required_beats", mcp.Description("required_beats")),
		mcp.WithString("forbidden_moves", mcp.Description("forbidden_moves")),
		mcp.WithString("continuity_checks", mcp.Description("continuity_checks")),
		mcp.WithString("evaluation_focus", mcp.Description("evaluation_focus")),
		mcp.WithString("emotion_target", mcp.Description("emotion_target")),
		mcp.WithString("payoff_points", mcp.Description("payoff_points")),
		mcp.WithString("hook_goal", mcp.Description("hook_goal")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewPlanChapterTool(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})

	s.AddTool(mcp.NewTool("novel.read_chapter",
		mcp.WithDescription("Executes novel.read_chapter"),
		mcp.WithNumber("chapter", mcp.Description("chapter")),
		mcp.WithNumber("from", mcp.Description("from")),
		mcp.WithNumber("to", mcp.Description("to")),
		mcp.WithString("source", mcp.Description("source")),
		mcp.WithString("character", mcp.Description("character")),
		mcp.WithNumber("max_runes", mcp.Description("max_runes")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewReadChapterTool(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})

	s.AddTool(mcp.NewTool("novel.reopen_book",
		mcp.WithDescription("Executes novel.reopen_book"),
		mcp.WithNumber("chapters", mcp.Description("chapters")),
		mcp.WithString("reason", mcp.Description("reason")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewReopenBookTool(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})

	s.AddTool(mcp.NewTool("novel.save_arc_summary",
		mcp.WithDescription("Executes novel.save_arc_summary"),
		mcp.WithNumber("volume", mcp.Description("volume")),
		mcp.WithNumber("arc", mcp.Description("arc")),
		mcp.WithString("title", mcp.Description("title")),
		mcp.WithString("summary", mcp.Description("summary")),
		mcp.WithString("key_events", mcp.Description("key_events")),
		mcp.WithString("character_snapshots", mcp.Description("character_snapshots")),
		mcp.WithString("style_rules", mcp.Description("style_rules")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewSaveArcSummaryTool(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})

	s.AddTool(mcp.NewTool("novel.save_foundation",
		mcp.WithDescription("Executes novel.save_foundation"),
		mcp.WithString("type", mcp.Description("type")),
		mcp.WithString("content", mcp.Description("content")),
		mcp.WithString("scale", mcp.Description("scale")),
		mcp.WithNumber("volume", mcp.Description("volume")),
		mcp.WithNumber("arc", mcp.Description("arc")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewSaveFoundationTool(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})

	s.AddTool(mcp.NewTool("novel.save_volume_summary",
		mcp.WithDescription("Executes novel.save_volume_summary"),
		mcp.WithNumber("volume", mcp.Description("volume")),
		mcp.WithString("title", mcp.Description("title")),
		mcp.WithString("summary", mcp.Description("summary")),
		mcp.WithString("key_events", mcp.Description("key_events")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewSaveVolumeSummaryTool(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})

	s.AddTool(mcp.NewTool("novel.export",
		mcp.WithDescription("Executes novel export"),
		mcp.WithString("output_path", mcp.Description("output_path")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewExportTool(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})

	s.AddTool(mcp.NewTool("novel.import",
		mcp.WithDescription("Executes novel import"),
		mcp.WithString("input_path", mcp.Description("input_path")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := tools.NewImportTool(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(res)), nil
	})
}
