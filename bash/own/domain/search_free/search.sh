#!/bin/bash  

# Путь к файлу со списком слов 
input_file="words.txt"

# Путь к файлу, в который будут сохранены доступные домены 
output_file="available.txt"

# Проверка существования файла со списком слов 
if [ ! -f "$input_file" ]; then   
	echo "Файл '$input_file' не найден."   
	exit 1 
fi  

# Очистка файла с результатами перед началом 
> "$output_file"  
# Цикл по каждому слову в файле 
while IFS= read -r word; do   
	# Формирование доменного имени   
	domain="$word.software"    
	# Проверка доступности домена   
	if whois "$domain" | grep -qi "No match"; then     
		# Если "No match" найдено в выводе WHOIS, домен доступен     
		echo "$domain" >> "$output_file"
		echo "(!) ДОСТУПЕН: $domain"
	else     
		echo "Домен $domain занят"   
	fi 
done < "$input_file"  

echo "Завершено. Результаты сохранены в '$output_file'."

