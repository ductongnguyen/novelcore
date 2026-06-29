# ainovel-core (MCP Server)

> **Lưu ý quan trọng**: Dự án đã được tái cấu trúc hoàn toàn. Giao diện TUI (`ainovel-cli`) và hệ thống Agent nội bộ (`agentcore`, `Host`, `Coordinator`) đã được loại bỏ. `ainovel-core` hiện tại chỉ đóng vai trò là một **Workflow Engine và Data Store** thuần túy, giao tiếp thông qua giao thức **MCP (Model Context Protocol)**.

`ainovel-core` cung cấp bộ công cụ (tools) cho phép một AI Agent bên ngoài (như **Antigravity**) đảm nhiệm vai trò là "Bộ não" (nhà văn, kiến trúc sư, biên tập viên) để viết tiểu thuyết.

## Kiến trúc mới

```text
┌───────────────────────────────────────────┐
│              Antigravity                  │
│  (Reasoning, Planning, LLM, Subagents)    │
│  Đóng vai trò: Writer, Architect, Editor  │
└─────────────────────┬─────────────────────┘
                      │ MCP Protocol (stdio)
┌─────────────────────▼─────────────────────┐
│           ainovel-mcp (Server)            │
│  Cung cấp tools: create_project,          │
│  plan_chapter, draft_chapter, save...     │
└─────────────────────┬─────────────────────┘
                      │
┌─────────────────────▼─────────────────────┐
│              ainovel-core                 │
│  (Domain logic, State, File System)       │
│  Quản lý: Outline, Continuity, Workflow   │
└───────────────────────────────────────────┘
```

Trong mô hình này:
- **Antigravity** là bộ não suy luận: Đọc file, hiểu ngữ cảnh, tự quyết định gọi hàm `plan_chapter`, sau đó tự sinh văn bản truyện và gọi `draft_chapter`, `check_consistency`, `commit_chapter`...
- **ainovel-core** là chuyên gia miền (Domain Expert): Đảm bảo tính nhất quán của dữ liệu, quản lý tiến độ, lưu trữ checkpoint và các quy trình workflow nghiêm ngặt.

## Cài đặt (Antigravity Plugin)

Để sử dụng `ainovel-core` với Antigravity, bạn cần cài đặt nó dưới dạng một Plugin:

1. **Build file thực thi MCP Server:**
   Đảm bảo bạn đã cài đặt [Go](https://golang.org/).
   ```bash
   # Build cho Windows
   env GOOS=windows go build -o ainovel-mcp.exe ./cmd/ainovel-mcp
   
   # Hoặc Build cho Linux/macOS
   go build -o ainovel-mcp ./cmd/ainovel-mcp
   ```

2. **Copy executable vào thư mục Plugin:**
   Copy file thực thi vừa build vào thư mục `ainovel-plugin`.
   ```bash
   cp ainovel-mcp.exe ainovel-plugin/
   ```

3. **Đăng ký Plugin vào Antigravity:**
   Copy toàn bộ thư mục `ainovel-plugin` vào thư mục plugins của Antigravity (thường nằm ở `~/.gemini/config/plugins/`).
   ```bash
   cp -r ainovel-plugin ~/.gemini/config/plugins/
   ```

4. **Cấu hình đường dẫn MCP:**
   Mở file `~/.gemini/config/plugins/ainovel-plugin/mcp_config.json` và sửa trường `command` trỏ tới đường dẫn tuyệt đối của file thực thi `ainovel-mcp.exe`.
   Ví dụ trên Windows:
   ```json
   {
     "mcpServers": {
       "ainovel-mcp": {
         "command": "C:\\Users\\<YourUser>\\.gemini\\config\\plugins\\ainovel-plugin\\ainovel-mcp.exe",
         "args": [],
         "env": {}
       }
     }
   }
   ```

5. **Khởi động lại Antigravity:**
   Tạo một phiên làm việc (chat) mới hoặc tải lại giao diện Antigravity. Hệ thống sẽ tự động nạp kỹ năng `novel-writer` và kết nối với MCP Server của ainovel.

## Cách sử dụng

Sau khi cài đặt thành công, bạn chỉ cần mở một thư mục truyện (hoặc thư mục trống) trong Antigravity và ra lệnh bằng ngôn ngữ tự nhiên:

- *"Tạo project truyện mới"*
- *"Lên dàn ý cho cuốn tiểu thuyết fantasy"*
- *"Viết chương 1"*
- *"Tiếp tục chương 2"*
- *"Kiểm tra continuity của truyện"*

Agent của Antigravity sẽ tự động kích hoạt kỹ năng `novel-writer`, đóng vai trò là nhà văn và dùng các công cụ MCP để thao tác đọc/ghi, lập kế hoạch và sáng tác trọn vẹn tác phẩm cho bạn!
