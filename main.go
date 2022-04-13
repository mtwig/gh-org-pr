package main

import (
	"fmt"
	"log"
	"time"

	"github.com/TwiN/go-color"
	"github.com/cli/go-gh"
	graphql "github.com/cli/shurcooL-graphql"
)

func main() {

	client, err := gh.GQLClient(nil)

	if err != nil {
		log.Fatal(err)
	}

	var graphqlResponse struct {
		Search struct {
			Nodes []struct {
				PullRequest struct {
					Title          string
					Url            string
					CreatedAt      time.Time
					UpdatedAt      time.Time
					IsDraft        bool
					BaseRepository struct {
						Id   string
						Name string
						Url  string
					}
					Author struct {
						Login string
						User  struct {
							Name string
						} `graphql:"... on User"`
					}
				} `graphql:"... on PullRequest"`
			}
		} `graphql:"search(type: ISSUE, first: $first, query: $query)"`
	}
	variables := map[string]interface{}{
		"first": graphql.Int(100),
		"query": graphql.String("is:pr is:open review-requested:@me sort:updated-desc"),
	}
	err = client.Query("Search", &graphqlResponse, variables)
	if err != nil {
		log.Fatal(err)
	}

	type PullRequest struct {
		Title       string
		Url         string
		CreatedAt   time.Time
		UpdatedAt   time.Time
		IsDraft     bool
		Author      string
		AuthorLogin string
	}
	type Repository struct {
		Id    string
		Name  string
		Url   string
		Pulls []PullRequest
	}

	repos := make(map[string]*Repository)

	for _, node := range graphqlResponse.Search.Nodes {
		var pr = node.PullRequest

		var repository *Repository = &Repository{}
		r, exists := repos[pr.BaseRepository.Id]
		if exists {
			repository = r
		} else {
			repository.Id = pr.BaseRepository.Id
			repository.Name = pr.BaseRepository.Name
			repository.Url = pr.BaseRepository.Url
			repository.Pulls = []PullRequest{}
			repos[pr.BaseRepository.Id] = repository
		}

		// var organization Organization
		var pullRequest PullRequest
		pullRequest.Title = pr.Title
		pullRequest.Url = pr.Url
		pullRequest.CreatedAt = pr.CreatedAt
		pullRequest.UpdatedAt = pr.UpdatedAt
		pullRequest.IsDraft = pr.IsDraft
		pullRequest.Author = pr.Author.User.Name
		pullRequest.AuthorLogin = pr.Author.Login

		repository.Pulls = append(repository.Pulls, pullRequest)
	}

	for _, repo := range repos {
		fmt.Print(color.Ize(color.Green, fmt.Sprintf("[ %s ]\n", repo.Name)))
		for _, pull := range repo.Pulls {
			fmt.Print(color.Ize(color.Gray, fmt.Sprintf("  - %s ", pull.Title)))
			if pull.IsDraft {
				fmt.Print(color.Ize(color.Cyan, "[Draft]"))
			}
			fmt.Print("\n")
			fmt.Print(color.Ize(color.Blue, fmt.Sprintf("    %s\n", pull.Url)))
			fmt.Printf("    added by %s (%s)\n",
				color.Ize(color.Yellow, pull.Author),
				color.Ize(color.Yellow, pull.AuthorLogin))

			var created = getDateString(time.Since(pull.CreatedAt))
			var modified = getDateString(time.Since(pull.UpdatedAt))
			time.Since(pull.CreatedAt)

			fmt.Printf("    Created %s\n", color.Ize(color.Purple, created))
			fmt.Printf("    Modifed %s\n", color.Ize(color.Purple, modified))
		}
	}
}

func getDateString(duration time.Duration) string {
	var num = duration.Hours()
	var unit = "hours"

	if duration.Hours() > 48 {
		num = num / 24
		unit = "days"
	}
	return fmt.Sprintf("%.1f %s", num, unit)
}
