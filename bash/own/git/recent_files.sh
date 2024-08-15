#!/bin/bash

git diff --name-only HEAD HEAD^ > tmp.txt
while read -r file; do
  if [[ $file == bash/* ]]; then
    echo "Found file in bash/ $file"
  else
    echo $file
  fi
done < tmp.txt
rm -f tmp.txt
