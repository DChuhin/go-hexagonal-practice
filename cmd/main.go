package main

import (
	"context"
	"go-hexagonal-practice/internal/boot"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	boot.RunApplication(ctx)
}
