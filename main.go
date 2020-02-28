package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/domainr/whois"
)

func toLetters(num int) string {
	power := num / 26
	mod := num % 26

	var output string

	if mod > 0 {
		output = string('A' - 1 + mod)
	} else {
		power -= 1
		output = string('Z')
	}

	var returnValue string

	if power > 0 {
		returnValue = toLetters(power) + output
	} else {
		returnValue = output
	}

	return strings.ToLower(returnValue)
}

func appendFile(value string) {
	file, err := os.OpenFile("names.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("\n-------------------------------")
		fmt.Println("Write to file failed.")
		fmt.Println(value)
		fmt.Println("-------------------------------\n")
	}
	defer file.Close()
	if _, err := file.WriteString(value + "\n"); err != nil {
		fmt.Println("\n-------------------------------")
		fmt.Println("Write to file failed.")
		fmt.Println(value)
		fmt.Println("-------------------------------\n")
	}
}

func nameAvailable(name string) bool {
	request, err := whois.NewRequest(name)

	if err != nil {
		fmt.Println("\n-------------------------------")
		fmt.Println("Request failed.")
		fmt.Println(name)
		fmt.Println("-------------------------------\n")
		return false
	}

	response, err := whois.DefaultClient.Fetch(request)

	if err != nil {
		fmt.Println("\n-------------------------------")
		fmt.Println("Response failed.")
		fmt.Println(name)
		fmt.Println("-------------------------------\n")
		return false
	}

	str := response.String()
	return strings.HasPrefix(str, "")
}

func main() {
	max := 100
	min := 27
	var wg sync.WaitGroup

	wg.Add(max - min)

	for i := min; i < max; i++ {
		go func(i int) {
			defer wg.Done()
			host := toLetters(i)
			fmt.Println(i, host)
			domain := host + ".io"
			available := nameAvailable(domain)
			if available {
				appendFile(domain)
			}
		}(i)
	}

	wg.Wait()
}
