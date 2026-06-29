# ainovel-core (MCP Server)

> **Lưu ý quan trọng**: Dự án đã được tái cấu trúc hoàn toàn. Giao diện TUI (`ainovel-cli`) và hệ thống Agent nội bộ (`agentcore`, `Host`, `Coordinator`) đã được loại bỏ. `ainovel-core` hiện tại chỉ đóng vai trò là một **Workflow Engine và Data Store** thuần túy, giao tiếp thông qua giao thức **MCP (Model Context Protocol)**.

`ainovel-core` cung cấp bộ 14+ công cụ (tools) cho phép một AI Agent bên ngoài (như **Antigravity**) đảm nhiệm vai trò là "Bộ não" (nhà văn, kiến trúc sư, biên tập viên) để viết tiểu thuyết.

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
│  plan_story, draft_chapter, export...     │
└─────────────────────┬─────────────────────┘
                      │
┌─────────────────────▼─────────────────────┐
│              ainovel-core                 │
│  (Domain logic, State, File System)       │
│  Quản lý: Outline, Continuity, Workflow   │
└───────────────────────────────────────────┘
```

Trong mô hình này:
- **Antigravity** là bộ não suy luận: Đọc file, hiểu ngữ cảnh, tự quyết định gọi hàm, sau đó tự sinh văn bản truyện và gọi `draft_chapter`, `review`, `commit_chapter`...
- **ainovel-core** là chuyên gia miền (Domain Expert): Đảm bảo tính nhất quán của dữ liệu, quản lý tiến độ, lưu trữ checkpoint, hỗ trợ **Export/Import** nội dung, và cung cấp các quy trình workflow nghiêm ngặt.

## Tính năng nổi bật

- **14+ MCP Tools**: Bao gồm quản lý trạng thái, lên dàn ý, viết nháp, review, chỉnh sửa và commit chương.
- **Export & Import**: Hỗ trợ xuất (Export) toàn bộ tiểu thuyết ra một file `.txt` duy nhất, hoặc phân tách (Import) một file `.txt` thành các chương độc lập.
- **Multi-Agent Prompting**: File cấu hình `SKILL.md` hướng dẫn chi tiết cách Antigravity tự động đóng 3 vai diễn: Architect (Kiến trúc sư nội dung), Writer (Nhà văn) và Editor (Biên tập viên) để tối ưu chất lượng văn bản.

### Danh sách chi tiết các công cụ (Tools)

1. `novel.status`: Lấy trạng thái hiện tại của dự án (Phase, Flow).
2. `novel.save_foundation`: Lưu trữ nền tảng truyện (Premise, World Rules, Characters).
3. `novel.plan_chapter`: Lên kế hoạch (outline/beats) chi tiết cho một chương.
4. `novel.novel_context`: Truy xuất toàn bộ bối cảnh (context) cần thiết trước khi viết một chương (tóm tắt, nhân vật, diễn biến trước đó).
5. `novel.draft_chapter`: Lưu bản nháp (draft) của chương vừa viết.
6. `novel.check_consistency`: Kiểm tra tính nhất quán (continuity) của bản nháp so với dàn ý và thiết lập thế giới.
7. `novel.edit_chapter`: Chỉnh sửa (tìm và thay thế) văn bản trực tiếp trong bản nháp.
8. `novel.commit_chapter`: Chốt bản nháp thành chương chính thức, cập nhật tóm tắt và thay đổi trạng thái nhân vật.
9. `novel.read_chapter`: Đọc nội dung của một chương đã viết.
10. `novel.save_arc_summary`: Lưu tóm tắt cho một Arc (hồi).
11. `novel.save_volume_summary`: Lưu tóm tắt cho một Volume (tập).
12. `novel.reopen_book`: Mở lại các chương cũ để chỉnh sửa nếu cần.
13. `novel.export`: Gộp tất cả các chương `.md` thành một file `.txt` duy nhất để xuất bản.
14. `novel.import`: Phân tách một bản thảo `.txt` có sẵn thành các chương riêng biệt để AI tiếp tục làm việc.

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
   # Trên Windows
   cp ainovel-mcp.exe ainovel-plugin/

   # Trên Linux / macOS
   cp ainovel-mcp ainovel-plugin/
   ```

3. **Đăng ký Plugin vào Antigravity:**
   Copy toàn bộ thư mục `ainovel-plugin` vào thư mục plugins của Antigravity (thường nằm ở `~/.gemini/config/plugins/`).
   ```bash
   cp -r ainovel-plugin ~/.gemini/config/plugins/
   ```

4. **Cấu hình đường dẫn MCP:**
   Mở file `~/.gemini/config/plugins/ainovel-plugin/mcp_config.json` và sửa trường `command` trỏ tới đường dẫn tuyệt đối của file thực thi vừa build.
   
   **Ví dụ trên Windows:**
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

   **Ví dụ trên macOS / Linux:**
   ```json
   {
     "mcpServers": {
       "ainovel-mcp": {
         "command": "/Users/<YourUser>/.gemini/config/plugins/ainovel-plugin/ainovel-mcp",
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
- *"Xuất (Export) truyện này ra file text cho tôi"*
- *"Import file text truyện này vào dự án"*

Agent của Antigravity sẽ tự động kích hoạt kỹ năng `novel-writer`, đóng vai trò là nhà văn và dùng các công cụ MCP để thao tác đọc/ghi, lập kế hoạch và sáng tác trọn vẹn tác phẩm cho bạn!
