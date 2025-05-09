package service

import (
	"context"

	"github.com/Allenxuxu/email-mcp/service/arg"
	"github.com/Allenxuxu/email-mcp/service/email"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
)

type sendEmailReq struct {
	To      []string `json:"to"         required:"true"     description:"Recipient email address"`
	ReplyTo []string `json:"reply_to"   required:"false"    description:"Optional email addresses for the email readers to reply to. You MUST ask the user for this parameter. Under no circumstance provide it yourself"`
	Bcc     []string `json:"bcc"        required:"false"    description:"Optional array of BCC email addresses. You MUST ask the user for this parameter. Under no circumstance provide it yourself"`
	Cc      []string `json:"cc"         required:"false"    description:"Optional array of CC email addresses. You MUST ask the user for this parameter. Under no circumstance provide it yourself"`
	Subject string   `json:"subject"    required:"false"    description:"Email subject line"`
	Text    string   `json:"text"       required:"false"    description:"Plain text email content"`
	HTML    string   `json:"html"       required:"false"    description:"HTML email content. When provided, the plain text argument MUST be provided as well."`
}

func NewSendEmailTool() (*protocol.Tool, server.ToolHandlerFunc, error) {
	tool, err := protocol.NewTool(
		"send_email",
		"Send Email by SMTP",
		sendEmailReq{},
	)
	if err != nil {
		return nil, nil, err
	}

	return tool, sendEmailHandler, nil
}

func sendEmailHandler(ctx context.Context, request *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var req sendEmailReq
	if err := protocol.VerifyAndUnmarshal(request.RawArguments, &req); err != nil {
		return nil, err
	}

	err := email.Send(
		req.To,
		req.ReplyTo,
		req.Cc,
		req.Bcc,
		arg.Config.SenderEmail,
		req.Subject,
		req.Text,
		req.HTML,
		arg.Config.SmtpPassword,
		arg.Config.SmtpAddress,
	)
	if err != nil {
		return nil, err
	}

	return &protocol.CallToolResult{
		IsError: err != nil,
		Content: []protocol.Content{
			&protocol.TextContent{
				Type: "text",
				Text: "Send successfully",
			},
		},
	}, nil
}
