#!/bin/bash

echo -ne "bits 64\nglobal main\nmain:\ndb " > sc.asm;
cat input.txt | \
 sed -r 's/\\x(..)/0x\1, /g;s/, $//' >> sc.asm;
nasm -f elf64 sc.asm;
ld -Ttext 0x400000 -o sc sc.o;
file sc
objdump -M intel -D -w sc