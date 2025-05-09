package service

import (
	"fmt"
	"log"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
)

const (
	StdioMode      = "stdio"
	SseMode        = "sse"
	StreamableMode = "streamable"
)

func NewMCP(mode, address string) (*server.Server, error) {
	transport, err := getTransport(mode, address)
	if err != nil {
		log.Fatalf("Failed to get transport: %v", err)
		return nil, err
	}
	srv, err := server.NewServer(
		transport,
		server.WithServerInfo(protocol.Implementation{
			Name:    "Send-Email-MCP",
			Version: "1.0.0",
		}),
	)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
		return nil, err
	}

	tool, toolHandler, err := NewSendEmailTool()
	if err != nil {
		log.Fatalf("Failed to create tool: %v", err)
		return nil, err

	}

	srv.RegisterTool(tool, toolHandler)

	return srv, nil
}

func getTransport(mode, addr string) (t transport.ServerTransport, e error) {
	switch mode {
	case StdioMode:
		log.Println("start mcp server with stdio transport")
		t = transport.NewStdioServerTransport()
	case SseMode:
		log.Printf("start  mcp server with sse transport, listen %s", addr)
		t, _ = transport.NewSSEServerTransport(addr)
	case StreamableMode:
		log.Printf("start mcp server with streamable_http transport, listen %s", addr)
		t = transport.NewStreamableHTTPServerTransport(addr)
	default:
		return nil, fmt.Errorf("unknown mode: %s", mode)
	}

	return t, nil
}
