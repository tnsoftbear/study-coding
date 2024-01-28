#!/bin/bash

set -u

readonly GLOBAL_VAR='global'
my_function() {
    local local_variable=1
    OTHER_VARIABLE=2
    echo "GLOBAL_VAR in function is $GLOBAL_VAR"
}

my_function

echo "GLOBAL_VAR is $GLOBAL_VAR"
echo "OTHER_VARIABLE is $OTHER_VARIABLE"
echo "local_variable is $local_variable"
echo "missing_variable is $missing_variable"