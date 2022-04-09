package input

import (
	"errors"
	"flag"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/mtwig/gh-org-pr/theme"
	"regexp"
	"strings"
)

const regexTopics = "^[\\w][\\w,]*[\\w]$"

func getBool(input string, result *bool, err *error) {
	if strings.EqualFold(input, "true") || strings.EqualFold(input, "yes") {
		*result = true
		*err = nil
		return
	}
	if strings.EqualFold(input, "false") || strings.EqualFold(input, "no") {
		*result = false
		*err = nil
		return
	}

	*err = errors.New("illegal value")
}

func GetOptions() (org string, doFilter bool, topics []string, err error) {
	var topicFlag string
	flag.StringVar(&org, "org", "", "organization to list pulls from")
	flag.StringVar(&topicFlag, "topics", "", "comma separated list of required topicFlag")
	flag.Parse()

	if strings.EqualFold(org, "") {
		fmt.Print(color.Ize(theme.InputRequest, "Which organization? "))
		fmt.Scanln(&org)
	}
	if !strings.EqualFold(topicFlag, "") {
		doFilter = true
		match, matchErr := regexp.MatchString(regexTopics, topicFlag)
		if err != nil {
			err = matchErr
			return
		}
		if !match {
			err = errors.New("invalid topics input")
			return
		}
		topics = strings.Split(topicFlag, ",")
	} else {
		for true {
			var userInput = ""
			var parseErr error
			fmt.Print(color.Ize(theme.InputRequest, "Filter repositories by topic? "))
			fmt.Scanln(&userInput)
			getBool(userInput, &doFilter, &parseErr)
			if parseErr == nil {
				break
			}
			fmt.Printf(color.Ize(theme.ErrorMessage, "accepted values: true/false yes/no.\n"))
		}
		if doFilter {
			for true {
				fmt.Printf(color.Ize(theme.InputRequest, "Which topics would you like to include?  "))
				fmt.Scanln(&topicFlag)
				match, matchErr := regexp.MatchString(regexTopics, topicFlag)

				if err != nil {
					err = matchErr
					return
				}
				if !match || strings.Contains(topicFlag, ",,") {
					fmt.Println(color.Ize(theme.ErrorMessage, "input is not comma seperated string"))
					continue
				}
				topics = strings.Split(topicFlag, ",")
				break
			}
		}
	}
	return
}
