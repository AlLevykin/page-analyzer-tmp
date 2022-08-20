package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"page-analyzer/internal/analyzer"
	"page-analyzer/internal/urls"
	"syscall"
)

func main() {
	if len(os.Args) == 2 {

		us, err := urls.Parse(os.Args[1], ",")
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
		defer cancel()

		a := analyzer.NewAnalyzer(us)
		a.Process(ctx)

	}
}
