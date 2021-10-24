#!/usr/bin/env bash

main() {
    local input=${1// /}
    (( ${#input} < 2 )) && echo "false" && return 0;
    (( `expr length "${input//[0-9]/}"` )) && { echo "false"; return 0; }
    local acc=0
    local odd=$((${#input} % 2))
    for (( i=0; i<${#input}; i++)); do
      if (( i % 2 == odd )); then
       local b=$(( 2 * ${input:i:1} ))
       if (( b > 9 )); then b=$(( b - 9 )); fi
       acc=$(( acc + b ))
      else
       acc=$(( acc + ${input:i:1} ))
      fi
    done
    if (( acc % 10 )); then echo "false"; else echo "true"; fi
}

main "$@"
