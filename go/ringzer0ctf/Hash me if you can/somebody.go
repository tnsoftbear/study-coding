package somebody

import (
	"crypto/sha512"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const token = "ahgpohik2pvkqe80k0cdvh7mk6"

func main() {
	url := "https://ringzer0ctf.com/challenges/13/"
	httpClient := http.Client{
		Timeout: time.Second * 1, // Maximum of 1 secs
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "PHPSESSID",
		Value: token,
	})

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	msgBegin := "----- BEGIN MESSAGE -----"
	msgEnd := "----- END MESSAGE -----"
	unhashed := strings.TrimSpace(strings.ReplaceAll(strings.Split(strings.Split(string(body[:]), msgBegin)[1], msgEnd)[0], "<br />", ""))
	h512 := sha512.New()
	io.WriteString(h512, unhashed)

	url = fmt.Sprintf("%s%x", url, h512.Sum(nil))

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  "PHPSESSID",
		Value: token,
	})

	res, getErr = httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr = ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	println(strings.Split(strings.Split(string(body[:]), `<div class="challenge-wrapper">`)[1], "</div>")[0])
}
