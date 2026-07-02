import re

with open(r'd:\Coding\ainovel-cli\internal\mcp\tools.go', 'r', encoding='utf-8') as f:
    content = f.read()

# Replace RegisterTools signature
content = content.replace('func RegisterTools(s *server.MCPServer, store *store.Store)', 'func RegisterTools(s *server.MCPServer, storeManager *store.StoreManager)')

# Find all s.AddTool occurrences
tool_pattern = re.compile(r'(s\.AddTool\(mcp\.NewTool\("[^"]+",\n(?:\s+mcp\.With[^,]+,\n)+)\s*\), func\(ctx context\.Context, request mcp\.CallToolRequest\) \(\*mcp\.CallToolResult, error\) \{')

def replacer(match):
    prefix = match.group(1)
    new_prefix = prefix + '\t\tmcp.WithString("project_dir", mcp.Description("Optional absolute path to the novel directory. Defaults to CWD.")),\n'
    return new_prefix + '\t), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {'

content = tool_pattern.sub(replacer, content)

# Special case for novel.status which doesn't have any With... parameters initially
status_pattern = re.compile(r'(s\.AddTool\(mcp\.NewTool\("novel\.status",\n\s+mcp\.WithDescription\("[^"]+"\),\n)\s*\), func\(ctx context\.Context, request mcp\.CallToolRequest\) \(\*mcp\.CallToolResult, error\) \{')
def status_replacer(match):
    prefix = match.group(1)
    new_prefix = prefix + '\t\tmcp.WithString("project_dir", mcp.Description("Optional absolute path to the novel directory. Defaults to CWD.")),\n'
    return new_prefix + '\t), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {'
content = status_pattern.sub(status_replacer, content)

# Inject store retrieval at the start of each tool func
func_pattern = re.compile(r'(\), func\(ctx context\.Context, request mcp\.CallToolRequest\) \(\*mcp\.CallToolResult, error\) \{\n)')

def store_inject_replacer(match):
    return match.group(1) + '\t\tdir, _ := request.Params.Arguments["project_dir"].(string)\n\t\tstore := storeManager.GetStore(dir)\n'

content = func_pattern.sub(store_inject_replacer, content)

with open(r'd:\Coding\ainovel-cli\internal\mcp\tools.go', 'w', encoding='utf-8') as f:
    f.write(content)
