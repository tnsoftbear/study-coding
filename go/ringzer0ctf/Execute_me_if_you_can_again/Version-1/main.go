package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/gocolly/colly"
)

const CHALLENGE_URL = "https://ringzer0ctf.com/challenges/125"
const BEGIN_MARKER = "----- BEGIN SHELLCODE -----"
const END_MARKER = "----- END SHELLCODE -----"

var cookieHeader string

func grabShellcode(siteContent string) string {
	re := regexp.MustCompile(BEGIN_MARKER + `\s*(.+)\s*` + END_MARKER)
	matches := re.FindAllStringSubmatch(siteContent, -1)
	rawShellcode := matches[0][1]
	fmt.Printf("Input shellcode: %s\n", rawShellcode)
	return rawShellcode
}

func fixShellcode(rawShellcodeStr string) string {
	rawShellcode := []byte(rawShellcodeStr)
	rawShellcode[6] = byte('5')
	rawShellcode[7] = byte('0')
	rawShellcode[335] = byte('b')
	fmt.Printf("Fixed shellcode: %s\n", rawShellcode)
	return string(rawShellcode)
}

func writeShellcodeToFile(shellcode string) {
	file, _ := os.Create("shellcode.txt")
	file.Write([]byte(shellcode))
	file.Close()
}

func execShellcodeFromFile() string {
	// This works only in windows, and "bash -c" starts wsl to call this command
	out, err := exec.Command("./execshellcode").Output()
	if err != nil {
		fmt.Printf("Error exec command: %v\n", err)
	}
	fmt.Printf("testshell output: %s\n", out)
	return string(out)
}

func sendResultAndReadFlag(result string) {
	c2 := colly.NewCollector()
	c2.OnRequest(func(r2 *colly.Request) {
		r2.Headers.Set("Cookie", cookieHeader)
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
	phpsessid := flag.String("phpsessid", "", "PHP Session ID")
	flag.Parse()
	if len(*phpsessid) == 0 {
		fmt.Println("PHP Session ID value missed\n")
		flag.PrintDefaults()
		return
	}

	cookieHeader = "PHPSESSID=" + *phpsessid
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", cookieHeader)
	})
	c.OnHTML("div[class=message]", func(e *colly.HTMLElement) {
		shellcode := grabShellcode(e.Text)
		shellcode = fixShellcode(shellcode)
		writeShellcodeToFile(shellcode)
		ascii := execShellcodeFromFile()
		sendResultAndReadFlag(ascii)
	})
	c.Visit(CHALLENGE_URL)
}
