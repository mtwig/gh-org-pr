package ghpr

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
	"github.com/mtwig/gh-org-pr/exitcodes"
	"github.com/mtwig/gh-org-pr/theme"
	"net/http"
	"os"
)

func getHttpClient() *http.Client {
	client, err := gh.HTTPClient(nil)
	if err != nil {
		fmt.Printf(color.Ize(theme.ErrorMessage, "unable to initialize http client"))
		os.Exit(exitcodes.UnableToCreateClient)
	}
	return client
}

func getRestClient() api.RESTClient {
	client, err := gh.RESTClient(nil)
	if err != nil {
		fmt.Printf(color.Ize(theme.ErrorMessage, "unable to initialize rest client"))
		os.Exit(exitcodes.UnableToCreateClient)
	}
	return client
}
