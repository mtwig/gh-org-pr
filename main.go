package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/mtwig/gh-org-pr/exitcodes"
	"github.com/mtwig/gh-org-pr/input"
	"github.com/mtwig/gh-org-pr/theme"
	"os"
)

func main() {

	org, doFilter, topics, err := input.GetOptions()

	if err != nil {
		fmt.Print(color.Ize(theme.ErrorMessage, fmt.Sprintf("unable to parse flags [%s]\n", err)))
		os.Exit(exitcodes.UnableToParseInput)
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
