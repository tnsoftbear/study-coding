#!/usr/bin/env bash

# name=${1}
# printf "One for %s, one for me." "${name:=you}"

main () {
  if [[ "$1" = "" ]] 
  then
    NICK="you"
  else
    NICK=$1
  fi
  echo "One for $NICK, one for me."
}

main "$@"