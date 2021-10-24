#!/usr/bin/env bash

main () {
    result=`echo "$1" | sed -E "s/'//ig" | sed -E "s/[^a-z]+/ /ig" | sed -E "s/(\w)\w* */\1/ig"`
    echo ${result^^}
}

main "$@"


# words="${1//[^a-zA-Z\']/ }"
# acronym=""
# for word in $words; do
#   acronym="$acronym${word:0:1}"
# done
# echo "${acronym^^}"


# declare -ra WORDS=( ${1//[![:alpha:]]/ } )
# for word in "${WORDS[@]}"; do
#   output+=${word:0:1}
# done
# printf "%s\n" ${output^^}