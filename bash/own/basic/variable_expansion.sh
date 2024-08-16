#!/bin/bash

# set -x

V1="c:/path/to/file.txt"
echo ${V1#*/}
echo ${V1##*/}
echo ${V1%/*}
echo ${V1%%/*}

V2="c:/path/to/path.txt"
echo ${V2/path/xxx}
echo ${V2//path/xxx}
echo ${V2/#c/xxx}
echo ${V2/%txt/xxx}

