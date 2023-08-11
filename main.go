package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/devoc09/caps/internal/cli"
)

func main() {
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
    defer stop()
    c := cli.NewCLI()
    if err := c.Run(ctx, os.Args); err != nil {
        log.Fatal(err)
    }
}
