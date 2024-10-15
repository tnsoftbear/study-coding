#!/bin/bash

# Проходим по всем файлам в текущей директории
for file in *; do
    # Проверяем, является ли это файлом
    if [ -f "$file" ]; then
        # Получаем расширение файла
        extension="${file##*.}"

        # Получаем имя файла без расширения
        filename="${file%.*}"

        # Проверяем, длина имени больше 12 символов, чтобы избежать ошибок
        if [ ${#filename} -gt 12 ]; then
            # Обрезаем последние 12 символов
            new_filename="${filename::${#filename}-12}"

            # Переименовываем файл с сохранением расширения
            mv "$file" "${new_filename}.${extension}"

            echo "Переименован: $file -> ${new_filename}.${extension}"
        else
            echo "Файл $file слишком короткий для обрезки"
        fi
    fi
done
