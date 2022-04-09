package types

type GhRepo struct {
	Name   string   `json:"name"`
	Topics []string `json:"topics"`
}

type GhPull struct {
	HtmlUrl string `json:"html_url""`
	State   string `json:"state"`
	Title   string `json:"title"`
	User    struct {
		Login string `json:"login"`
	} `json:"user"`
	RequestedReviewers []struct {
		Login string
	} `json:"requested_reviewers"`
}
