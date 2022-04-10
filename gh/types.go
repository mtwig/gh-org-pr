package ghpr

type User struct {
	Login string
}

type Repo struct {
	Name   string
	Topics []string
}

type Pull struct {
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
