package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/mtwig/gh-org-pr/exitcodes"
	ghpr "github.com/mtwig/gh-org-pr/gh"
	"github.com/mtwig/gh-org-pr/input"
	"github.com/mtwig/gh-org-pr/theme"
	"os"
	"strings"
)

func main() {

	org, doFilter, topics, err := input.GetOptions()
	var me = ghpr.Me()
	if err != nil {
		fmt.Print(color.Ize(theme.ErrorMessage, fmt.Sprintf("unable to parse flags [%s]\n", err)))
		os.Exit(exitcodes.UnableToParseInput)
	}
	theme.PrintDebug(fmt.Sprintf("Org [%s]\n", org))
	theme.PrintDebug(fmt.Sprintf("Do topic filtering: [%v]\n", doFilter))
	theme.PrintDebug(fmt.Sprintln("Topics ", topics))

	var openPullsReport strings.Builder
	var orgRepos = ghpr.OrganizationRepos(org, topics)

	for _, repo := range orgRepos {
		var repoReport strings.Builder
		repoReport.WriteString(color.Ize(theme.RepositoryTitle, fmt.Sprintf("[ %s ]\n", repo)))
		var pulls = ghpr.RepoPulls(org, repo)
		var printRepo = false

		for _, pull := range pulls {
			for _, reviewer := range pull.RequestedReviewers {
				//fmt.Printf("am I %s\n?", reviewer.Login)
				if strings.EqualFold(me, reviewer.Login) {
					printRepo = true
					repoReport.WriteString(color.Ize(theme.PullUrl, fmt.Sprintf("  • %s\n", pull.HtmlUrl)))
					repoReport.WriteString(fmt.Sprintf("    %s [%s]\n", color.Ize(theme.PullTitle, pull.Title), color.Ize(theme.PullState, pull.State)))
					repoReport.WriteString(fmt.Sprintf("    by %s\n", color.Ize(theme.UserLogin, pull.User.Login)))
				}
			}
		}
		if printRepo {
			openPullsReport.WriteString(repoReport.String())
		}

	}

	fmt.Print(openPullsReport.String())
}
