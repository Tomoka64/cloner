package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	token      = flag.String("token", "", "the token of github")
	org        = flag.String("org", "", "the name of organization")
	typeOfRepo = flag.String("type", "private", "private or public. default: private")
	page       = flag.Int("page", 10, "the page num. default: 1")
	perPage    = flag.Int("per", 100, "the number of results to include per page. default: 100")
	language   = flag.String("language", "", "the language you want to get")
)

func main() {
	flag.Parse()

	if *token == "" || *org == "" {
		fmt.Println("-token and -org are requered")
		os.Exit(1)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *token},
	)

	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	opt := &github.RepositoryListByOrgOptions{
		Type:        *typeOfRepo,
		ListOptions: github.ListOptions{Page: *page, PerPage: *perPage},
	}

	repos, _, err := client.Repositories.ListByOrg(context.Background(), *org, opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan bool)
	for _, repo := range repos {
		go func(repo github.Repository) {
			if *language != "" && repo.GetLanguage() != *language {
				ch <- false
				return
			}

			_, err := os.Stat(*repo.Name)
			if err == nil {
				ch <- false
				return
			}

			fmt.Println(fmt.Sprintf("git clone %s", *repo.SSHURL))

			cmd := exec.Command("git", "clone", *repo.SSHURL)
			cmd.Start()
			ch <- true
		}(*repo)
	}

	for i := 0; i < len(repos); i++ {
		<-ch
	}
}
