# Plugin Source Synchronization

When making changes to plugin files (e.g., `SKILL.md`, `mcp_config.json`, or `.exe` binaries) for the `ainovel-plugin`, you MUST ensure that the changes are applied to both:
1. The installed plugin directory: `C:\Users\LEGION\.gemini\config\plugins\ainovel-plugin\`
2. The source code repository: `d:\Coding\ainovel-cli\ainovel-plugin\`

This guarantees that the active agent uses the latest updates immediately, while the changes remain properly tracked in version control for future commits.
