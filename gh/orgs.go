package ghpr

import (
	"encoding/json"
	"fmt"
	"github.com/mtwig/gh-org-pr/theme"
	"net/http"
	"os"
	"strings"
)

func OrganizationRepos(org string, topics []string) (repos []string) {
	var repositories []string
	var url = fmt.Sprintf("https://api.github.com/orgs/%s/repos?per_page=100", org)
	organizationReposByUrl(&repositories, topics, getHttpClient(), url)
	return repositories
}

func organizationReposByUrl(repositories *[]string, filterTopics []string, client *http.Client, url string) {

	response, err := client.Get(url)
	if err != nil {
		return
	}
	var repos []Repo

	err = json.NewDecoder(response.Body).Decode(&repos)
	if err != nil {
		os.Exit(99)
	}

RepoLoop:
	for _, repo := range repos {
		//theme.PrintDebug(fmt.Sprintf("The repo is [%s]\n", repo.Name))
		//theme.PrintDebug(" Topics:\n")
		for _, repoTopic := range repo.Topics {
			theme.PrintDebug(fmt.Sprintf("    %s\n", repoTopic))
			for _, filterTopic := range filterTopics {
				//theme.PrintDebug(fmt.Sprintf("  comparing [%s] and [%s]\n", repoTopic, filterTopic))
				if strings.EqualFold(repoTopic, filterTopic) {
					*repositories = append(*repositories, repo.Name)
					continue RepoLoop
				}

			}
		}
	}
	next, err := NextPageUrl(response.Header.Get("Link"))
	if err == nil {
		organizationReposByUrl(repositories, filterTopics, client, next)
	}

}
