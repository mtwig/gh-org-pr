package ghpr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func RepoPulls(org string, repo string) (pulls []Pull) {
	var url = fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls?per_page=100&sort_created&direction=desc&state=open", org, repo)
	repoPullsByUrl(&pulls, getHttpClient(), url)
	return
}

func repoPullsByUrl(pulls *[]Pull, client *http.Client, url string) {
	response, err := client.Get(url)
	if err != nil {
		return
	}
	var pullPage []Pull
	err = json.NewDecoder(response.Body).Decode(&pullPage)
	if err != nil {
		os.Exit(99)
	}
	*pulls = append(*pulls, pullPage...)
	next, err := NextPageUrl(response.Header.Get("Link"))
	if err == nil {
		repoPullsByUrl(pulls, client, next)
	}
}

/*
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
*/
