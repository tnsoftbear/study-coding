package main

import (
	"fmt"
	"regexp"
	"github.com/gocolly/colly"
	"strconv"
)

const PHPSESSID = "62u05k7rubiigheghn8kvpq394"
const CHALLENGE_URL = "https://ringzer0ctf.com/challenges/32"

func extractRawInput(siteContent string) string {
	re := regexp.MustCompile(`----- BEGIN MESSAGE -----\s*(.+)\s*----- END MESSAGE -----`)
	matches := re.FindAllStringSubmatch(siteContent, -1)
	inputBinString := matches[0][1]
	fmt.Printf("inputBinString: %s\n", inputBinString)
	return inputBinString
}

func calculateExpression(expression string) int64 {
	re := regexp.MustCompile(`\s*(\d+)\s*\+\s*0x([abcdef\d]+)\s\-\s*(\d+).*`)
	matches := re.FindAllStringSubmatch(expression, -1)
	// fmt.Printf("Matches: %v", matches)
	decStr := matches[0][1]
	hexStr := matches[0][2]
	binStr := matches[0][3]
	dec, _ := strconv.ParseInt(decStr, 10, 64)
	hex, _ := strconv.ParseInt(hexStr, 16, 64)
	bin, _ := strconv.ParseInt(binStr, 2, 64)
	result := dec + hex - bin
	fmt.Printf("Matches: %v + %v - %v = %v \n", dec, hex, bin, result)
	return result
}

func sendResultAndReadFlag(result int64) {
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
		result := calculateExpression(inputExpression)
		sendResultAndReadFlag(result)
	})
	c.Visit(CHALLENGE_URL)
}
