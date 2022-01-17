#!/bin/bash

function check_space {
        if [[ $1 == *[bdks';''&'' ']* ]]
        then
                echo "true"
                return 0
        fi

        echo "false"

        return 1
}

while :
do
        echo "Your input:"
        read input
        if check_space "$input"
        then
                echo -e '\033[0;31mRestricted characters has been used\033[0m'
        else
                output="echo Your command is: $input"
                eval $output
        fi
done
