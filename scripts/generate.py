import os, re

def go_type_to_mcp(gotype, name):
    if 'int' in gotype:
        return f'mcp.WithNumber("{name}", mcp.Description("{name}"))'
    if 'bool' in gotype:
        return f'mcp.WithBoolean("{name}", mcp.Description("{name}"))'
    return f'mcp.WithString("{name}", mcp.Description("{name}"))'

tools_dir = 'internal/tools'
blocks = []

for file in os.listdir(tools_dir):
    if file.endswith('.go') and not file.endswith('_test.go'):
        content = open(os.path.join(tools_dir, file), encoding='utf-8').read()
        matches = re.findall(r'var a struct \{(.*?)\}', content, re.DOTALL)
        
        if matches:
            struct_body = matches[0]
            fields = []
            for line in struct_body.strip().split('\n'):
                line = line.strip()
                if not line: continue
                parts = re.split(r'\s+', line)
                if len(parts) >= 3:
                    field_name = parts[0]
                    field_type = parts[1]
                    json_tag = re.search(r'`json:"(.*?)"`', line)
                    if json_tag:
                        jname = json_tag.group(1)
                        fields.append((jname, field_type))
            
            tool_name = "novel." + file.replace('.go', '')
            
            # find NewXXXTool
            new_func = re.search(r'func New([A-Za-z0-9_]+Tool)', content)
            if not new_func:
                continue
            constructor = "tools.New" + new_func.group(1)
            
            mcp_args = []
            for jname, ftype in fields:
                mcp_args.append(go_type_to_mcp(ftype, jname))
            
            args_str = ",\n\t\t".join(mcp_args)
            
            block = f"""
	s.AddTool(mcp.NewTool("{tool_name}",
		mcp.WithDescription("Executes {tool_name}"),
		{args_str},
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {{
		argsRaw, _ := json.Marshal(request.Params.Arguments)
		f := {constructor}(store)
		res, err := f.Execute(ctx, argsRaw)
		if err != nil {{
			return mcp.NewToolResultError(err.Error()), nil
		}}
		return mcp.NewToolResultText(string(res)), nil
	}})
"""
            blocks.append(block)

# Open tools.go and inject
code = """package mcp

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

"""
code += "\n".join(blocks)
code += "\n}\n"

open("internal/mcp/tools.go", "w", encoding="utf-8").write(code)
print("Generated internal/mcp/tools.go")
