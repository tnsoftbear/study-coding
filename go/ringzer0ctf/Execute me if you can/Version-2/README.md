# Ringzer0cfg "Execute me if you can" challenge - Version 2

Solution for ["Execute me if you can"](https://ringzer0ctf.com/challenges/121) challenge

Run in linux (wsl)

```sh
# Compile shellcode execution programm
gcc -fno-stack-protector -z execstack execshellcode.c -o execshellcode
# Compile and run go application, that doewnloads shellcode, executes it, sends output and parse flag from response
go run main.go
# Output should look like this:
input string: \xeb\x4d\x5e\x66\x83\xec\x0c\x48\x89\xe0\x48\x31\xc9\x68\xe6\x78\x94\x4a\x48\x89\xcf\x80\xc1\x0c\x40\x8a\x3e\x40\xf6\xd7\x40\x88\x38\x48\xff\xc6\x68\x4c\x05\x0b\xaa\x48\xff\xc0\xe2\xea\x2c\x0c\x48\x89\xc6\x68\xe2\x65\x32\x78\x48\x31\xc0\x48\x89\xc7\x04\x01\x48\x89\xc2\x80\xc2\x0b\x0f\x05\x48\x31\xc0\x04\x3c\x0f\x05\xe8\xae\xff\xff\xff\xa8\x8a\xcb\xab\xab\xc7\x8e\x8a\x96\xbb\xb9\xb4\xd7\xf7\x2b\xc5\x6d\xef\x68\x37\x33\x9a\x2f\x5b\x52\x41\x4e\x44\x53\x54\x52\x32\x5d
testshell output: Wu4TT8quiDFK
Solution url: https://ringzer0ctf.com/challenges/121/Wu4TT8quiDFK
flag: FLAG-W2gudjVCAlhexK1c3IfPun0CGs
```

`execshellcode.c` executes shellcode.

If you look into `strace ./execshellcode` it writes to stdin ("0") file descriptor

```c
read(3, "\\xeb\\x4d\\x5e\\x66\\x83\\xec\\x0c\\x48"..., 512) = 508
read(3, "", 512)                        = 0
mprotect(0x55b9c5234000, 127, PROT_EXEC) = 0
write(0, "Ce07Kd50Zv1z", 12Ce07Kd50Zv1z)            = 12
```

Thus we need to reallocate stdin file descriptor to stdout with help of `dup(1, 0)`

## Links

* [Problem discussion](https://stackoverflow.com/questions/29593556/directing-shellcode-output-to-a-file-c)
* [Execute shellcode from txt file](https://stackoverflow.com/questions/17842499/read-and-execute-shellcode-from-a-txt-file)
* [Online Assembler and Disassembler](http://shell-storm.org/online/Online-Assembler-and-Disassembler/)
* [Go: os/File](https://pkg.go.dev/os#File.Write)
* [Go exec/Cmd](https://pkg.go.dev/os/exec#example-Cmd.Output)

## Task

You have 1 second to execute this code and get the output.
Send the answer back using https://ringzer0ctf.com/challenges/121/[string]

```txt
----- BEGIN SHELLCODE -----
\xeb\x4d\x5e\x66\x83\xec\x0c\x48\x89\xe0\x48\x31\xc9\x68\x45\x89\xbf\x2d\x48\x89\xcf\x80\xc1\x0c\x40\x8a\x3e\x40\xf6\xd7\x40\x88\x38\x48\xff\xc6\x68\xcf\xf0\xb4\x4d\x48\xff\xc0\xe2\xea\x2c\x0c\x48\x89\xc6\x68\xaa\xb6\xe4\x7f\x48\x31\xc0\x48\x89\xc7\x04\x01\x48\x89\xc2\x80\xc2\x0b\x0f\x05\x48\x31\xc0\x04\x3c\x0f\x05\xe8\xae\xff\xff\xff\xcb\xa7\xb2\x88\xbb\xb2\x8c\xa6\x87\xb1\xb1\xbd\x55\x43\x71\xe3\x32\x4f\x68\xda\xe4\x8c\xcc\x5b\x52\x41\x4e\x44\x53\x54\x52\x32\x5d
----- END SHELLCODE -----
```

## Solution description

"github.com/gocolly/colly" go library is used for fetching site data and sending calculation result as well.

## Other solutions (not mine)
