---
name: wsl-go
description: How to use Go compiler located in WSL from Windows
---

## WSL Go Compiler Access

When you need to compile a Go project for this user using Windows PowerShell, DO NOT rely on the standard Windows `go` command if it is unavailable or if compiling for Windows from WSL is required.

The `go` executable is located in the user's WSL environment at `/usr/local/go/bin/go`.

### Usage Instructions

To run Go commands, wrap them in a `wsl` invocation:

```bash
wsl -e bash -l -c 'cd /mnt/d/Coding/ainovel-cli && env GOOS=windows /usr/local/go/bin/go build -o output.exe ./cmd/path'
```

- Always use `wsl -e bash -l -c '...'` to ensure the correct environment variables are loaded.
- Translate Windows paths to WSL paths (e.g., `d:\Coding\ainovel-cli` becomes `/mnt/d/Coding/ainovel-cli`).
- If you need to copy the resulting executable to a Windows directory that is currently locking the file (like an active plugin), copy it to a new file name (e.g., `file-v2.exe`) and update the corresponding configuration JSON instead of trying to overwrite the locked file.
