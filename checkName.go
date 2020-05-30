package main

import (
	"sync"
)

func checkName(name string, done func()) {
	details := newEntry(name)
	defer done()
	defer func (details *nameDetails) {
		details.writeLog()
	}(details)

	var wg sync.WaitGroup

	comAvailable, err := domainAvailable(name + ".com")
	if err != nil {
		details.err = err.Error()
		return;
	}
	details.comAvailable = comAvailable
	if !comAvailable {
		details.err = "Com not available"
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		twitterAvailable, err := url404s("https://twitter.com/"+name)
		if err != nil {
			details.err = err.Error()
			return;
		}
		details.twitterAvailable = twitterAvailable
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		githubAvailable, err := url404s("https://github.com/"+name)
		if err != nil {
			details.err = err.Error()
			return;
		}
		details.githubAvailable = githubAvailable
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		npmAvailable, err := url404s("https://www.npmjs.com/org/"+name)
		if err != nil {
			details.err = err.Error()
			return;
		}
		details.npmAvailable = npmAvailable
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slackAvailable, err := url404s("https://"+name+".slack.com")
		if err != nil {
			details.err = err.Error()
			return;
		}
		details.slackAvailable = slackAvailable
	}()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		redditSubAvailable, err := url404s("https://reddit.com/r/"+name)
		if err != nil {
			details.err = err.Error()
			return;
		}
		details.redditSubAvailable = redditSubAvailable
	}()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		redditUserAvailable, err := url404s("https://reddit.com/u/"+name)
		if err != nil {
			details.err = err.Error()
			return;
		}
		details.redditUserAvailable = redditUserAvailable
	}()

	defer wg.Wait()
}
