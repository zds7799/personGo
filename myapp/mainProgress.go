package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
)

var (
	mainPro = &gcmd.Command{
		Name:  "main",
		Brief: "main process",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Debug(ctx, `this is main process`)
			return gproc.ShellRun(ctx, `go run subProgress.go`)
		},
	}
)

func main() {
	mainPro.Run(gctx.New())
}
