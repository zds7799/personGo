package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Main = &gcmd.Command{
		Name:        "main",
		Brief:       "start http server",
		Description: "this is the command entry for starting your process",
	}
	Http = &gcmd.Command{
		Name:        "http",
		Brief:       "start http server",
		Description: "this is the command entry for starting your http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Println("start http server")
			return
		},
	}
	Grpc = &gcmd.Command{
		Name:        "grpc",
		Brief:       "start grpc server",
		Description: "this is the command entry for starting your grpc server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Println("start grpc server")
			return
		},
	}
)

func main() {
	err := Main.AddCommand(Http, Grpc)
	if err != nil {
		panic(err)
	}
	Main.Run(gctx.New())
}
