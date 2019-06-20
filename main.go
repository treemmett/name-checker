package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/haccer/available"
)

func main() {
	name := os.Args[1]

	fmt.Print(".com\t\t")

	available := available.Domain(name + ".com")

	if available {
		fmt.Println("🔥")
	} else {
		fmt.Println("👎")
	}

	names := []string{"Github", "NPM", "Slack", "Reddit", "Twitter", "Insta", "FB"}

	urls := []string{
		"https://github.com/" + name,
		"https://www.npmjs.com/org/" + name,
		"https://" + name + ".slack.com/",
		"https://www.reddit.com/r/" + name,
		"https://twitter.com/" + name,
		"https://www.instagram.com/" + name,
		"https://www.facebook.com/" + name}

	for i, url := range urls {
		fmt.Print(names[i] + "\t\t")
		resp, _ := http.Get(url)

		if resp.StatusCode == 200 {
			fmt.Println("👎")
		} else {
			fmt.Println("🔥")
		}
	}
}
