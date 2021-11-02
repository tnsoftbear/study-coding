package main

import (
	"fmt"
	"regexp"
	"crypto/sha512"
	"io"
	"github.com/gocolly/colly"
	"strconv"
)

func main() {
	token := "62u05k7rubiigheghn8kvpq394"
	url := "https://ringzer0ctf.com/challenges/14"

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", "PHPSESSID=" + token)
	})
	c.OnHTML("div[class=message]", func(e *colly.HTMLElement) {
		re := regexp.MustCompile(`----- BEGIN MESSAGE -----\s*([01]+)\s*----- END MESSAGE -----`)
		matches := re.FindAllStringSubmatch(e.Text, -1)
		inputBinString := matches[0][1]
		charTotal := len(inputBinString) / 8
		var chars []byte
		var i int
		for i = 0; i < charTotal; i++ {
			charBinString := inputBinString[i*8:(i+1)*8]
			if charInt, err := strconv.ParseInt(charBinString, 2, 8); err != nil {
				fmt.Print(err)
			} else {
				chars = append(chars, byte(charInt))
			}
		}
		var input = string(chars)
		h512 := sha512.New()
		io.WriteString(h512, input)
		url2 := fmt.Sprintf("%s/%x", url, h512.Sum(nil))
		fmt.Printf("Challenge url: %s\n", url2)
		
		c2 := colly.NewCollector()
		c2.OnRequest(func(r2 *colly.Request) {
			r2.Headers.Set("Cookie", "PHPSESSID=" + token)
		})
		c2.OnHTML("html", func (e2 *colly.HTMLElement) {
			re := regexp.MustCompile(`FLAG-[\w\d]+`)
			flag := re.Find([]byte(e2.Text))
			fmt.Printf("flag: %s\n", flag)
		})
		c2.Visit(url2)
	})
	c.Visit(url)
}
