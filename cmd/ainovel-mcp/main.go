package main

import (
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"github.com/voocel/ainovel-cli/internal/mcp"
	"github.com/voocel/ainovel-cli/internal/store"
)

func main() {
	// Initialize the core store (domain expert)
	// Note: typically we load project from CWD or config.
	// Here we mock the initialization for the entrypoint.
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get working directory: %v\n", err)
		os.Exit(1)
	}

	// We initialize a new novel store manager with CWD as default
	storeManager := store.NewStoreManager(cwd)

	// Create MCP server
	mcpServer := server.NewMCPServer(
		"ainovel-mcp",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
	)

	// Register tools
	mcp.RegisterTools(mcpServer, storeManager)

	// Start the stdio server
	if err := server.ServeStdio(mcpServer); err != nil {
		fmt.Fprintf(os.Stderr, "MCP server error: %v\n", err)
		os.Exit(1)
	}
}
