package main

import (
	"flag"

	"github.com/jgfranco17/gitglow/core/pkg/scan"
	"github.com/jgfranco17/gitglow/core/pkg/stats"
)

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan.Scan(folder)
		return
	}

	stats.GetStats(email)
}
