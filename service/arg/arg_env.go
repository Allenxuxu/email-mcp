package arg

import (
	"github.com/alexflint/go-arg"
)

var Config struct {
	Mode    string `arg:"env" default:"stdio"           help:"Optional values: stdio/sse/streamable"`
	Address string `arg:"env" default:"127.0.0.1:8080"  help:"Optional values: use for sse/streamable mode"`

	SmtpPassword string `arg:"env,required"`
	SmtpAddress  string `arg:"env,required" help:"smtp.example.com:465"`
	SenderEmail  string `arg:"env,required" help:"Sender email address"`
}

func MustInit() {
	arg.MustParse(&Config)
}
