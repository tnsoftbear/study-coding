#!/bin/bash

function sum1() {
  result=0
  for param in "$@"; do
    result=$(($result+$param))
  done
  echo $result
}

function sum2() {
  result=$(($1+$2))
  return $result
}

function sum3() {
  echo $(($1+$2))
}

sum2 1 2
result=$?
echo "1 + 2 = $result"

res=$(sum3 3 4)
echo "3 + 4 = $res"

result=$(sum1 5 6 7)
echo "5 + 6 + 7 = $result"

