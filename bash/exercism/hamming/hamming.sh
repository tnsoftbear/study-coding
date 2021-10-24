#!/usr/bin/env bash

main () {
    local count=0;
    (( ${#@} != 2 )) && echo "Usage: hamming.sh <string1> <string2>" && return 1;
    (( ${#1} != ${#2} )) && echo "strands must be of equal length" && return 1;
    for ((i=0; i<${#1}; i++))
    do
      # if [ ${1:$i:1} != ${2:$i:1} ]; then count=$((count+1)); fi
      # [[ "${1:$i:1}" != "${2:$i:1}" ]] && ((count++));
      [ ${1:$i:1} != ${2:$i:1} ] && ((count++));
    done
    echo $count;
}

main "$@"