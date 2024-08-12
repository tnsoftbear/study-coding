str="hello world"

# В [ ] нужно обязательно заключать строки с пробелами в кавычки:
if [ "$str" = "hello world" ]; then
  echo "Строки совпадают"
fi

# В [[ ]] кавычки не обязательны:
if [[ $str = hello\ world ]]; then
  echo "Строки совпадают"
fi