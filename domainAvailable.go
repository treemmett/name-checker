package main

import (
	"strings"

	"github.com/domainr/whois"
)

func domainAvailable(domain string) (bool, error) {
	req, err := whois.NewRequest(domain)
	if err != nil {
		return false, err
	}
	res, err := whois.DefaultClient.Fetch(req)
		if err != nil {
			return false, err
		}
	str := res.String()
	str = strings.ToLower(str)

	if strings.HasPrefix(str, "no") {
		return true, nil
	}

	return false, nil
}