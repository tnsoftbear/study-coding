package main

/**
 * This program downloads shellcode, converts it to binary format and save in file
 * Run linux tool for disassembling this binary shellcode:
 * $ ndisasm -b64 shellcode.bin
 */

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
)

const PHPSESSID = "sjgbeb33698aj6sgi6ni99t8u6"
const CHALLENGE_URL = "https://ringzer0ctf.com/challenges/125"
const BEGIN_MARKER = "----- BEGIN SHELLCODE -----"
const END_MARKER = "----- END SHELLCODE -----"

func grabShellcode(siteContent string) string {
	re := regexp.MustCompile(BEGIN_MARKER + `\s*(.+)\s*` + END_MARKER)
	matches := re.FindAllStringSubmatch(siteContent, -1)
	rawShellcode := matches[0][1]
	fmt.Printf("input string: %s\n", rawShellcode)
	return rawShellcode
}

func writeToFile(shellcode []byte) {
	file, _ := os.Create("shellcode.bin")
	file.Write(shellcode)
	file.Close()
}

func main() {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", "PHPSESSID="+PHPSESSID)
	})
	c.OnHTML("div[class=message]", func(e *colly.HTMLElement) {
		shellcode := grabShellcode(e.Text)
		bin := toBin(shellcode)
		writeToFile(bin)
	})
	c.Visit(CHALLENGE_URL)
}

func toBin(expression string) []byte {
	re := regexp.MustCompile(`\\x(.{2})`)
	matches := re.FindAllStringSubmatch(expression, -1)
	//fmt.Printf("Matches count: %d => %v\n", len(matches), matches)
	var charBytes []byte
	for _, match := range matches {
		hexStr := match[1]
		dec, err := strconv.ParseUint(hexStr, 16, 8)
		if err != nil {
			fmt.Printf("Error ParseInt(%v) %v\n", hexStr, err)
		}
		charBytes = append(charBytes, byte(dec))
		// fmt.Printf("hexStr: %v, dec: %d, byte(dec): %d\n", hexStr, dec, byte(dec))
	}
	return charBytes
}
