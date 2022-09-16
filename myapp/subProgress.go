package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	SubPro = &gcmd.Command{
		Name:  "sub",
		Brief: "sub process",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Debug(ctx, `this is sub process`)
			return nil
		},
	}
)

func main() {
	SubPro.Run(gctx.New())
}
