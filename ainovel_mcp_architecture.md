# Kiến trúc hoạt động tầng dưới của AINovel-MCP

Biểu đồ tuần tự dưới đây thể hiện cách mà hệ thống vận hành đằng sau hậu trường. Điểm mấu chốt là **MCP Server không tự có trí tuệ**, nó chỉ đóng vai trò là "cầu nối" (cung cấp các API quản lý state) và "bộ nhớ" (đọc/ghi file). Mọi logic suy nghĩ, đóng vai, tự đánh giá đều được thực thi bởi **LLM** (mô hình AI) dựa trên bộ quy tắc khắt khe của `SKILL.md`.

```mermaid
sequenceDiagram
    autonumber
    actor User as Người dùng
    participant LLM as LLM (Đóng các vai AI)
    participant MCP as MCP Server (ainovel-mcp)
    participant FS as File System (Ổ cứng)

    User->>LLM: Yêu cầu viết truyện / Góp ý ý tưởng
    
    note over LLM, MCP: 1. Khởi tạo & Kiểm tra (Vai trò: Coordinator)
    LLM->>MCP: Gọi `novel.status`
    MCP->>FS: Đọc file trạng thái (state.json)
    FS-->>MCP: Trả về tiến độ hiện tại
    MCP-->>LLM: Trả về trạng thái (Đang ở tập nào, chương nào, chờ gì)

    note over LLM, MCP: 2. Lập kế hoạch (Vai trò: Architect)
    LLM->>MCP: Gọi `novel.novel_context` (Yêu cầu ngữ cảnh)
    MCP->>FS: Trích xuất các bản tóm tắt Arc/Volume
    MCP-->>LLM: Trả về bối cảnh (Không gửi toàn bộ chữ để tiết kiệm token)
    
    alt Có góp ý/Sửa đổi nền tảng
        LLM->>MCP: Gọi `novel.save_foundation`
        MCP->>FS: Cập nhật quy tắc thế giới, nhân vật
    end
    
    LLM->>MCP: Gọi `novel.plan_chapter`
    MCP->>FS: Lưu trữ dàn ý chương mới
    
    note over LLM, MCP: 3. Viết nội dung (Vai trò: Writer)
    LLM->>MCP: Gọi `novel.draft_chapter` (Có thể gọi lặp lại nhiều vòng)
    MCP->>FS: Ghi nối văn bản vào file draft.md
    
    note over LLM, MCP: 4. Kiểm duyệt chéo (Vai trò: Editor)
    LLM->>MCP: Gọi `novel.check_consistency`
    MCP->>FS: Đọc draft hiện tại & đối chiếu foundation
    MCP-->>LLM: Trả về kết quả đánh giá (Lỗi logic, OOC, nhịp độ...)
    
    alt Phát hiện có lỗi logic/OOC
        LLM->>MCP: Gọi `novel.edit_chapter`
        MCP->>FS: Ghi đè file draft với bản đã sửa
    end
    
    note over LLM, MCP: 5. Chốt chương & Nén ngữ cảnh (Vai trò: Coordinator)
    LLM->>MCP: Gọi `novel.commit_chapter`
    MCP->>FS: Lưu bản chính thức, tăng biến đếm chương
    
    opt Nếu kết thúc Mạch truyện (Arc) hoặc Tập (Volume)
        LLM->>MCP: Gọi `novel.save_arc_summary` / `save_volume_summary`
        MCP->>FS: Tạo file tóm tắt nén chặt lại để dùng cho Tương lai
    end
    
    LLM-->>User: Báo cáo kết quả hoặc xin phép đi tiếp
```

### Giải thích các tầng (Layers)

1. **Tầng Tương tác (User <-> LLM):** 
   - Bạn chỉ giao tiếp bằng ngôn ngữ tự nhiên.
   - LLM tự hiểu ý định của bạn và quyết định nên đóng vai trò nào (Architect để sửa bối cảnh, Writer để viết tiếp, hoặc Editor để kiểm tra).

2. **Tầng Logic & Vòng lặp (LLM <-> MCP Server):**
   - Sự kết hợp giữa bộ kỹ năng (SKILL) và các công cụ (Tool). 
   - Các tool như `check_consistency` hoạt động như những chốt chặn để bắt LLM phải dừng lại tự nhìn nhận lỗi của mình trước khi nhảy sang bước tiếp theo.

3. **Tầng Lưu trữ (MCP Server <-> File System):**
   - Đảm bảo tính "bền vững" (persistent) của dữ liệu. 
   - Ngay cả khi bạn tắt cửa sổ chat và mở lại vào hôm sau, `novel.status` và các hàm đọc/ghi sẽ móc nối chính xác vào điểm dừng cuối cùng. Đây là lõi sức mạnh giúp mô hình ngôn ngữ vốn "não cá vàng" có thể nhớ được bối cảnh của 500 chương trước đó.
