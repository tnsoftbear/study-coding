#!/bin/bash

# dollarStar.sh
echo "With *:"
for arg in "$*"; do
  echo "<$arg>"
done

echo

echo "With @:"
for arg in "$@"; do
  echo "<$arg>"
done
