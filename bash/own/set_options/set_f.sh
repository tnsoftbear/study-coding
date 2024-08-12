#!/bin/bash
# Disable pathname expansion with -f option
set -f
# Таким образом * больше не соответствует всем файлам, а как литерал *
for file in *
 do
  echo $file
 done
