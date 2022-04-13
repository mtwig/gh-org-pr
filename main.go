package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/TwiN/go-color"
	"github.com/cli/go-gh"
	graphql "github.com/cli/shurcooL-graphql"
	ghm "github.com/mtwig/gh-org-pr/ghm"
)

func main() {

	client, err := gh.GQLClient(nil)

	if err != nil {
		log.Fatal(err)
	}
	var orgFilter string
	var repoNameFilter string
	flag.StringVar(&orgFilter, "org", "", "organization to show filters for")
	flag.StringVar(&repoNameFilter, "repo-filter", ".*", "only show repo names that match this pattern")
	flag.Parse()

	repoNameRegex, err := regexp.Compile(repoNameFilter)
	if err != nil {
		log.Fatal(err)
	}

	// log.Printf("The filter is %s\n", orgFilter)
	var graphqlResponse ghm.GQLResponse
	variables := map[string]interface{}{
		"first": graphql.Int(100),
		"query": graphql.String("is:pr is:open review-requested:@me sort:updated-desc"),
	}
	err = client.Query("Search", &graphqlResponse, variables)
	if err != nil {
		log.Fatal(err)
	}

	repos := make(map[string]*ghm.Repository)

	for _, node := range graphqlResponse.Search.Nodes {
		var pr = node.PullRequest

		var repository *ghm.Repository = &ghm.Repository{}
		r, exists := repos[pr.BaseRepository.Id]
		if exists {
			repository = r
		} else {
			repository.Id = pr.BaseRepository.Id
			repository.Name = pr.BaseRepository.Name
			repository.Url = pr.BaseRepository.Url
			repository.Pulls = []ghm.PullRequest{}
			repository.Organization.Id = pr.BaseRepository.Owner.Id
			repository.Organization.Name = pr.BaseRepository.Owner.Login
			if strings.EqualFold(orgFilter, "") ||
				strings.EqualFold(repository.Organization.Name, orgFilter) {
				repos[pr.BaseRepository.Id] = repository
			}
		}

		// var organization Organization
		var pullRequest ghm.PullRequest
		pullRequest.Title = pr.Title
		pullRequest.Url = pr.Url
		pullRequest.CreatedAt = pr.CreatedAt
		pullRequest.UpdatedAt = pr.UpdatedAt
		pullRequest.IsDraft = pr.IsDraft
		pullRequest.Author = pr.Author.User.Name
		pullRequest.AuthorLogin = pr.Author.Login

		if repoNameRegex.MatchString(repository.Name) {
			repository.Pulls = append(repository.Pulls, pullRequest)
		}

	}

	for _, repo := range repos {
		if len(repo.Pulls) == 0 {
			continue
		}
		fmt.Print(color.Ize(color.Green, fmt.Sprintf("[ %s ]\n", repo.Name)))
		for _, pull := range repo.Pulls {
			fmt.Print(color.Ize(color.Gray, fmt.Sprintf("  - %s ", pull.Title)))
			if pull.IsDraft {
				fmt.Print(color.Ize(color.Cyan, "[Draft]"))
			}
			fmt.Print("\n")
			fmt.Print(color.Ize(color.Blue, fmt.Sprintf("    %s\n", pull.Url)))
			fmt.Printf("    added by %s ", color.Ize(color.Yellow, "@"+pull.AuthorLogin))

			if !strings.EqualFold(pull.Author, "") {
				fmt.Printf("(%s)", color.Ize(color.Yellow, pull.Author))
			}
			fmt.Printf("\n")
			color.Ize(color.Yellow, pull.Author)
			var created = getDateString(time.Since(pull.CreatedAt))
			var modified = getDateString(time.Since(pull.UpdatedAt))
			time.Since(pull.CreatedAt)

			fmt.Printf("    created %s, modified %s\n",
				color.Ize(color.Purple, created+" ago"),
				color.Ize(color.Purple, modified+" ago"))
		}

	}
}

// getDateString converts a duration into a string
// representation.
// [0-1) hours: report in minutes
// [1-48] hours: report in hours
// (48,) hours: report in days
func getDateString(duration time.Duration) string {
	var num int
	var unit string

	//minutes
	if duration.Hours() < 1 {
		num = int(duration.Minutes())
		unit = "minute"
	} else if duration.Hours() >= 48 {
		num = int(duration.Hours() / 24)
		unit = "day"
	} else {
		num = int(duration.Hours())
		unit = "hour"
	}

	if num != 1 {
		unit += "s"
	}
	return fmt.Sprintf("%d %s", num, unit)
}
