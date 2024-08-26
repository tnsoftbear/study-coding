package main

import (
	"crypto/sha512"
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"regexp"
	"strings"
)

func main() {
	token := "ahgpohik2pvkqe80k0cdvh7mk6"
	url := "https://ringzer0ctf.com/challenges/13"

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", "PHPSESSID="+token)
	})
	c.OnHTML("div[class=message]", func(e *colly.HTMLElement) {
		msgBegin := "----- BEGIN MESSAGE -----"
		msgEnd := "----- END MESSAGE -----"
		input := strings.TrimSpace(strings.ReplaceAll(strings.Split(strings.Split(string(e.Text[:]), msgBegin)[1], msgEnd)[0], "<br />", ""))
		// re := regexp.MustCompile(`[01]+`)
		// input := re.Find([]byte(e.Text))
		println(string(input))
		h512 := sha512.New()
		io.WriteString(h512, string(input))
		url2 := fmt.Sprintf("%s/%x", url, h512.Sum(nil))
		println(url2)

		c2 := colly.NewCollector()
		c2.OnRequest(func(r2 *colly.Request) {
			r2.Headers.Set("Cookie", "PHPSESSID="+token)
		})
		c2.OnHTML("html", func(e2 *colly.HTMLElement) {
			re := regexp.MustCompile(`FLAG-[\w\d]+`)
			flag := re.Find([]byte(e2.Text))
			println(fmt.Sprintf("flag: %s", flag))
		})
		c2.Visit(url2)
	})
	c.Visit(url)
}
