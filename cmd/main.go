package main

import (
	"context"
	"github.com/danyukod/observability-optl-go/pkg"
	"os"
	"os/signal"
)

const serviceName = "obsercability-optl-go"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	pkg.InitObservability(ctx)
}
