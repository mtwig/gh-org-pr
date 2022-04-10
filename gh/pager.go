package ghpr

import (
	"errors"
	"regexp"
	"strings"
)

// NextPageUrl accepts the Link header
// and returns the URL of the next page
func NextPageUrl(link string) (url string, funcErr error) {
	pattern := regexp.MustCompile("\\<(.*)\\>")
	var links = strings.Split(link, ",")
	for _, link := range links {
		if strings.HasSuffix(strings.TrimSpace(link), "rel=\"next\"") {
			var res = pattern.FindAllStringSubmatch(link, 1)

			if len(res) < 1 {
				funcErr = errors.New("no match")
				return
			} else if len(res[0]) < 2 {
				funcErr = errors.New("no sub-match")

				return
			}
			url = res[0][1]
			return
		}

	}
	funcErr = errors.New("no next page")
	return
}
