#!/bin/bash

file="list.txt"

while IFS= read -r url; do
  sudo youtube-dl -vvv -f 137+140 --user-agent "Mozilla/5.0 (Android 14; Mobile; rv:128.0) Gecko/128.0 Firefox/128.0" "$url"
done < "$file"

