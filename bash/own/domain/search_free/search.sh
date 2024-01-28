#!/bin/bash  

set +x

# Путь к файлу со списком слов 
#input_file="words.txt"
input_file="english-adjectives.txt"

# Путь к файлу, в который будут сохранены доступные домены 
output_file=".results/available_adjectives.domain.lv.txt"

# Проверка существования файла со списком слов 
if [ ! -f "$input_file" ]; then   
	echo "Файл '$input_file' не найден."   
	exit 1 
fi  

check_domain_availability() {
  local domain="$1"
  local output_file="$2"

  if whois "$domain" | grep -qiE "Domain not found|No match|No Data Found"; then
    echo "$domain" >> "$output_file"
    echo "(!) ДОСТУПЕН: $domain"
  else
    echo "Домен $domain занят"
  fi
}

# Очистка файла с результатами перед началом 
> "$output_file"  

while IFS= read -r word; do
	word=`echo "$word" | tr -d '[:space:]'`
	domains=("${word}casino.com" "${word}-casino.com")
	for domain in "${domains[@]}"; do
		check_domain_availability "$domain" "$output_file"
	done
done < "$input_file"  

echo "Завершено. Результаты сохранены в '$output_file'."

