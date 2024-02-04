package main

import (
	"flag"
	"fmt"

	"github.com/jgfranco17/gitglow/core/pkg/models"
	"github.com/jgfranco17/gitglow/core/pkg/scan"
	"github.com/jgfranco17/gitglow/core/pkg/stats"
)

func main() {
	command := flag.String("command", "help", "Command to run")
	username := flag.String("user", "", "User repository owner")
	folder := flag.String("add", "", "Add a new folder to scan for Git repositories")
	email := flag.String("email", "your@email.com", "The email to scan")
	flag.Parse()

	userProject := models.Project{
		Name:   *username,
		Email:  *email,
		Folder: *folder,
	}
	if flag.NArg() == 0 {
		fmt.Printf("Hello, user!\n")
	} else {
		switch *command {
		case "help":
			fmt.Printf("Welcome to GitGlow!\n")
		case "run":
			if *folder != "" {
				scan.Scan(userProject.Folder)
				return
			}
			stats.GetStats(userProject.Email)
		}
	}
}
