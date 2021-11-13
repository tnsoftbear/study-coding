package main

import (
	"fmt"
	"regexp"
	"github.com/gocolly/colly"
	"strconv"
)

const PHPSESSID = "62u05k7rubiigheghn8kvpq394"
const CHALLENGE_URL = "https://ringzer0ctf.com/challenges/121"
const BEGIN_MARKER = "----- BEGIN SHELLCODE -----"
const END_MARKER = "----- END SHELLCODE -----"

func extractRawInput(siteContent string) string {
	re := regexp.MustCompile(BEGIN_MARKER + `\s*(.+)\s*` + END_MARKER)
	matches := re.FindAllStringSubmatch(siteContent, -1)
	rawInput := matches[0][1]
	fmt.Printf("input string: %s\n", rawInput)
	return rawInput
}

func toAscii(expression string) string {
	re := regexp.MustCompile(`\\x(.{2})`)
	matches := re.FindAllStringSubmatch(expression, -1)
	// fmt.Printf("Matches: %v", matches)
	var charBytes []byte
	for _, match := range matches {
		hexStr := match[1]
		dec, _ := strconv.ParseInt(hexStr, 16, 8)
		charBytes = append(charBytes, byte(dec))
	}
	ascii := string(charBytes)
	// fmt.Printf("%s", ascii)
	return ascii
}

func sendResultAndReadFlag(result string) {
	c2 := colly.NewCollector()
	c2.OnRequest(func(r2 *colly.Request) {
		r2.Headers.Set("Cookie", "PHPSESSID="+PHPSESSID)
	})
	c2.OnHTML("html", func(e2 *colly.HTMLElement) {
		re := regexp.MustCompile(`FLAG-[\w\d]+`)
		flag := re.Find([]byte(e2.Text))
		fmt.Printf("flag: %s\n", flag)
	})

	solutionUrl := fmt.Sprintf("%s/%v", CHALLENGE_URL, result)
	fmt.Printf("Solution url: %s\n", solutionUrl)
	c2.Visit(solutionUrl)
}

func main() {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", "PHPSESSID="+PHPSESSID)
	})
	c.OnHTML("div[class=message]", func(e *colly.HTMLElement) {
		inputExpression := extractRawInput(e.Text)
		ascii := toAscii(inputExpression)
		sendResultAndReadFlag(ascii)
	})
	c.Visit(CHALLENGE_URL)
}
