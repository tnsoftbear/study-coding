#!/bin/bash

usage() {
  echo "Usage: script.sh [-h|--help] [--ini-name <name>]"
}

# Обработка опций командной строки
options=$(getopt -o 'h?' --long 'help,ini-name:' -- "$@")

# Проверка кода возврата getopt
[ $? -eq 0 ] || { usage >&2 && false; }

#echo $options

# Передача опций в eval для обработки
eval set -- "$options"

# Цикл обработки опций
while true; do
  case "$1" in
    -h|--help)
      usage
      exit 0
      ;;
    --ini-name)
      ini_name="$2"
      shift 2
      ;;
    --)
      shift
      break
      ;;
    *)
      echo "Internal error!"
      exit 1
      ;;
  esac
done

# Ваши действия после обработки опций
echo "INI Name: $ini_name"
