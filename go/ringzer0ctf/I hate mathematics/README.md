# Ringzer0cfg "I hate mathematics" challenge

Solution for ["I hate mathematics"](https://ringzer0ctf.com/challenges/32) challenge

## Task

You have 2 seconds to send the answer
Send the answer back using https://ringzer0ctf.com/challenges/32/[answer]

```txt
----- BEGIN MESSAGE -----
6472 + 0x1e06 - 1010111101010 = ?
----- END MESSAGE -----
```

## Solution description

"github.com/gocolly/colly" go library is used for fetching site data and sending calculation result as well.

This challenge displays ariphmetic expression that consists of decimal, hexadecimal and binary number. To solve it, we need to parse this expression input to numeric strings, and convert them to integers in decimal numeric system on base 10. Then calculate expression and post it to challenge url.

## Other solutions

[Bash solution article](http://linusson.org/2018/11/09/i-hate-mathematics-the-challenge-not-the-science/)

### Bash-1

```sh
#!/bin/bash

wget --load-cookie cookies.txt http://ringzer0team.com/challenges/32 -O tmp.txt
NB1=`grep -A1 "BEGIN MESSAGE" tmp.txt | grep -v "BEGIN MESSAGE" | awk '{print $1}'`
NB2=16#`grep -A1 "BEGIN MESSAGE" tmp.txt | grep -v "BEGIN MESSAGE" | awk '{print $3}'  | cut -d"x" -f2`
NB3=2#`grep -A1 "BEGIN MESSAGE" tmp.txt | grep -v "BEGIN MESSAGE" | awk '{print $5}'`
RES=$(($NB1 + $NB2 - $NB3))
wget --load-cookie cookies.txt http://ringzer0team.com/challenges/32/$RES
rm -f tmp.txt
```

### Bash-2

```sh
#!/bin/bash
# dec + hex - bin = ?

if [ $# -ne 1 ]; then
   echo "Usage: $0 [PHPSESSID]"
   exit 1
fi

LEVEL=32
COOKIE="PHPSESSID=$1"
URL="https://ringzer0team.com/challenges/$LEVEL"
XPATH_MESG="//div[@class='message']"
XPATH_FLAG="//div[@class='alert alert-info']/text()"

# readchall(cookie, url)
# Read the challenge from ringzer0team.com
readchall() {
  curl -skb $1 $2
}

# readmsg(chall, n)
# Read message N off the challange
readmsg() {
  echo $1 | \
  xmllint \
    --nowarning \
    --recover \
    --xpath \
    "$XPATH_MESG[$2]/text()" - 2>/dev/null | \
  sed -n 3p | tr -d '\t' | sed 's/ //g'
}

# readflag(cookie, url)
# Read flag off the challange
readflag() {
  curl -skb $1 $2 | \
  xmllint \
    --nowarning \
    --recover \
    --xpath \
    "$XPATH_FLAG" - 2>/dev/null
  echo
}

# main
CHALLENGE=$(readchall "$COOKIE" "$URL")
EQUATION=$(readmsg "$CHALLENGE" 1)

DATA=$(echo $EQUATION | sed -r 's/^([0-9]+)\+(0x[0-9a-f]+)\-([0-1]*)=\?$/\1 \2 \3/')
DATA=($DATA)
DEC=${DATA[0]}
HEX=${DATA[1]}; HEX=$(($HEX))
BIN=${DATA[2]}; BIN=$((2#$BIN))

ANSWER=$((DEC + HEX - BIN))

readflag "$COOKIE" "$URL/$ANSWER"

exit 0
```

### zsh

```sh
#!/bin/zsh
#zsh because sh bash give "value too great for base" errors and bc doesn't like mixed base calculations
url="https://ringzer0ctf.com/challenges/32/"
cookie="VALUE"

message=`curl -s $url --cookie "PHPSESSID=$cookie" | grep -A 1 "BEGIN MESSAGE" | sed "s/<br \/>//" | tail -n 1 | head -c-5 | sed "s/\(.* \- \)\(.*\)/\10b\2/"`
result=$(($message))
curl -s $url"$result" --cookie "PHPSESSID=$cookie" | grep FLAG
```
