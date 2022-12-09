package main

import (
	_ "go_blok/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"go_blok/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
