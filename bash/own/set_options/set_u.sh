#!/bin/bash

#set -u  # Включаем опцию -u (nounset)

echo "All positional parameters: $@"
echo "All positional parameters: $*"
echo "Value of unset_variable: $unset_variable"
echo "Value of unset_parameter: $1"

