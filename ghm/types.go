package ghm

import "time"

type GQLResponse struct {
	Search struct {
		Nodes []struct {
			PullRequest struct {
				Title          string
				Url            string
				CreatedAt      time.Time
				UpdatedAt      time.Time
				IsDraft        bool
				BaseRepository struct {
					Id    string
					Name  string
					Url   string
					Owner struct {
						Id    string
						Login string
					}
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

type Organization struct {
	Id   string
	Name string
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
	Id           string
	Name         string
	Url          string
	Organization Organization
	Pulls        []PullRequest
}
