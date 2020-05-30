package main

import (
	"fmt"
	"os"
	"strconv"
)

type nameDetails struct {
	name								string
	comAvailable 				bool
	twitterAvailable 		bool
	githubAvailable 		bool
	npmAvailable 				bool
	slackAvailable 			bool
	redditSubAvailable 	bool
	redditUserAvailable bool
	err									string
}

func newEntry(name string) *nameDetails {
	details := nameDetails{
		name: name,
		comAvailable: false,
		twitterAvailable: false,
		githubAvailable: false,
		npmAvailable: false,
		slackAvailable: false,
		redditSubAvailable: false,
		redditUserAvailable: false,
		err: "",
	}

	return &details
}

func (details nameDetails) writeLog() {
		_, err := os.Stat("log.csv")
		if os.IsNotExist(err) {
			f, err := os.Create("log.csv")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			os.Chmod("log.csv", 0777)
			f.WriteString("Name,Com,Twitter,Github,NPM,Slack,Reddit Sub,Reddit User,Error\n")
		}

		file, _ := os.OpenFile("log.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend);
		line := details.name + "," +
						strconv.FormatBool(details.comAvailable) + "," +
						strconv.FormatBool(details.twitterAvailable) + "," +
						strconv.FormatBool(details.githubAvailable) + "," +
						strconv.FormatBool(details.npmAvailable) + "," +
						strconv.FormatBool(details.slackAvailable) + "," +
						strconv.FormatBool(details.redditSubAvailable) + "," +
						strconv.FormatBool(details.redditUserAvailable) + "," +
						details.err + "\n"

		_, err = file.WriteString(line)
		if err != nil {
			fmt.Println(err)
		}
}