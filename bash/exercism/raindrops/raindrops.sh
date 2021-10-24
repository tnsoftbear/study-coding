#!/usr/bin/env bash

main () {
  let factor3=($1%3==0)
  let factor5=($1%5==0)
  let factor7=($1%7==0)
  output="";
  if (( $factor3 )); then output="${output}Pling"; fi
  if [[ $factor5 == 1 ]]; then output="${output}Plang"; fi
  if [ $factor7 == 1 ]; then output="${output}Plong"; fi
  if [[ -z $output ]]; then output=$1; fi
  echo $output
}

main "$@"


# (( $1 % 3 )) || result+=Pling
# (( $1 % 5 )) || result+=Plang
# (( $1 % 7 )) || result+=Plong
# echo ${result:-$1}


# (($1%3==0)) && output+=Pling
# (($1%5==0)) && output+=Plang
# (($1%7==0)) && output+=Plong
# echo ${output:-$1}


# n=$1
# declare -a divisors=(
#     [3]=Pling
#     [5]=Plang
#     [7]=Plong
# )
# out=""
# for i in "${!divisors[@]}"; do
#     if ((n%i == 0)); then
#         out="${out}${divisors[$i]}"
#     fi
# done
# if [[ -z "$out" ]]; then
#     echo "$n"
# else
#     echo "$out"
# fi


# function addIfDivisable {
#   if [ $(expr $1 % $2) == 0 ]
#   then 
#     result+=$3
#   fi
# } 
# addIfDivisable $1 3 'Pling'
# addIfDivisable $1 5 'Plang'
# addIfDivisable $1 7 'Plong' 
# echo ${result:-$1}


