package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/mtwig/gh-org-pr/input"
	"github.com/mtwig/gh-org-pr/theme"
	"os"
)

func main() {

	org, doFilter, topics, err := input.GetOptions()

	if err != nil {
		fmt.Print(color.Ize(theme.ErrorMessage, fmt.Sprintf("unable to parse flags [%s]\n", err)))
		os.Exit(10)
	}
	fmt.Printf("Org [%s]\n", org)
	fmt.Printf("Do topic filtering: [%v]\n", doFilter)
	fmt.Println("Topics ", topics)

	//client, err := gh.RESTClient(nil)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//response := struct{ Login string }{}
	//err = client.Get("user", &response)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
