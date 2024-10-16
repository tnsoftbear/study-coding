#!/bin/bash

# This script renames all files in the current directory by removing the last 12 characters from the file name (YouTube video ID).

for file in *; do
    if [ -f "$file" ]; then
        extension="${file##*.}"
        filename="${file%.*}"
        if [ ${#filename} -gt 12 ]; then
            new_filename="${filename::-12}"
            mv "$file" "${new_filename}.${extension}"
            echo "Renamed: $file -> ${new_filename}.${extension}"
        else
            echo "File name $file is too short for cut"
        fi
    fi
done
