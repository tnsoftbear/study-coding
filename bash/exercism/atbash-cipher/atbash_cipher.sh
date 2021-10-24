#!/usr/bin/env bash

forward="abcdefghijklmnopqrstuvwxyz0123456789"
backward="zyxwvutsrqponmlkjihgfedcba0123456789"
direct=""
revert=""

findIndex() {
  local i
  for ((i = 0; i < ${#direct}; i++)); do
    [ "${direct:$i:1}" == "$1" ] && return $i
  done
}

main() {
  local is_encode=0
  if [[ "$1" == "encode" ]]; then
    direct=$forward
    revert=$backward
    is_encode=1
  else
    direct=$backward
    revert=$forward
  fi
  local input=${2,,}
  input=${input//[^a-zA-Z0-9]/}
  local output=""
  local i
  for ((i = 0; i < ${#input}; i++)); do
    if ((is_encode && i != 0 && i % 5 == 0)); then output="$output "; fi
    char=${input:$i:1}
    findIndex "$char"
    index=$?
    output="$output${revert:$index:1}"
  done
  echo "$output"
}

main "$@"
