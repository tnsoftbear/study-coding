#!/bin/bash
# https://leancrew.com/all-this/2011/12/more-shell-less-egg/
# Run: ./more-shell-less-eggs.sh < README.md

# 1. Make one-word lines by transliterating the complement (-c) of the alphabet into newlines (note the quoted newline), and squeezing out (-s) multiple newlines.
# 2. Transliterate upper case to lower case.
# 3. Sort to bring identical words together.
# 4. Replace each run of duplicate words with a single representative and include a count (-c).
# 5. Sort in reverse (-r) numeric (-n) order.
# 6. Filter 10 first lines

tr -cs A-Za-z '\n' |
tr A-Z a-z |
sort |
uniq -c |
sort -rn |
head -${1}
