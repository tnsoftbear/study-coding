#!/usr/bin/env bash

main () {
  help="Usage: error_handling.sh <person>"
  if [[ $# -eq 1 ]]
  then echo "Hello, $1"
  else echo $help & return 1
  fi
}

main "$@"

#   case $# in
#     0) printf "%s" "Usage: ./error_handling <greetee>"; return 1 ;;
#     1) printf "%s" "Hello, $1"; return 0 ;;
#     *) return 1 ;;
#   esac


    # (( $# != 1 )) && echo "Usage: ./error_handling <greetee>" && exit 1
    # echo "Hello, ${1}"