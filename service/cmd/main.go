package main

import (
	"flag"

	"github.com/jgfranco17/gitglow/core/pkg/models"
	"github.com/jgfranco17/gitglow/core/pkg/scan"
	"github.com/jgfranco17/gitglow/core/pkg/stats"
)

func main() {
	username := flag.String("add", "", "add a new folder to scan for Git repositories")
	folder := flag.String("add", "", "add a new folder to scan for Git repositories")
	email := flag.String("email", "your@email.com", "the email to scan")
	flag.Parse()

	userProject := models.Project{
		Name:   *username,
		Email:  *email,
		Folder: *folder,
	}

	if *folder != "" {
		scan.Scan(userProject.Folder)
		return
	}

	stats.GetStats(userProject.Email)
}
