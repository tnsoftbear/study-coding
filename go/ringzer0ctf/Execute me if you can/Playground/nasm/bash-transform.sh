#!/bin/bash

echo -ne "bits 64\nglobal main\nmain:\ndb " > sc.asm;
echo "\xeb\x4d\x5e\x66\x83\xec\x0c\x48\x89\xe0\x48\x31\xc9\x68\xe6\x78\x94\x4a\x48\x89\xcf\x80\xc1\x0c\x40\x8a\x3e\x40\xf6\xd7\x40\x88\x38\x48\xff\xc6\x68\x4c\x05\x0b\xaa\x48\xff\xc0\xe2\xea\x2c\x0c\x48\x89\xc6\x68\xe2\x65\x32\x78\x48\x31\xc0\x48\x89\xc7\x04\x01\x48\x89\xc2\x80\xc2\x0b\x0f\x05\x48\x31\xc0\x04\x3c\x0f\x05\xe8\xae\xff\xff\xff\xa8\x8a\xcb\xab\xab\xc7\x8e\x8a\x96\xbb\xb9\xb4\xd7\xf7\x2b\xc5\x6d\xef\x68\x37\x33\x9a\x2f\x5b\x52\x41\x4e\x44\x53\x54\x52\x32\x5d" | \
 sed -r 's/\\x(..)/0x\1, /g;s/, $//' >> sc.asm;
nasm -f elf64 sc.asm;
ld -Ttext 0x400000 -o sc sc.o;
file sc
objdump -M intel -D -w sc