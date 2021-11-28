# Playground

## savebin.go

This program downloads shellcode, converts it to binary format and save in file.  
Run linux tool for disassembling this binary shellcode: `ndisasm -b64 shellcode.txt > disassemeblered.txt`

`disassemeblered.txt`

```sh
00000000  EB4D              jmp short 0x4f
00000002  5E                pop rsi
00000003  6683EC0C          sub sp,byte +0xc
00000007  4889E0            mov rax,rsp
0000000A  4831C9            xor rcx,rcx
0000000D  684A1698A7        push qword 0xffffffffa798164a
00000012  4889CF            mov rdi,rcx
00000015  80C10C            add cl,0xc
00000018  408A3E            mov dil,[rsi]
0000001B  40F6D7            not dil
0000001E  408838            mov [rax],dil
00000021  48FFC6            inc rsi
00000024  6823A1214C        push qword 0x4c21a123
00000029  48FFC0            inc rax
0000002C  E2EA              loop 0x18
0000002E  2C0C              sub al,0xc
00000030  4889C6            mov rsi,rax
00000033  68C0BCD29C        push qword 0xffffffff9cd2bcc0
00000038  4831C0            xor rax,rax   // rax := 0
0000003B  4889C7            mov rdi,rax   // rdi := 0
0000003E  0401              add al,0x1    // rax += 1 (=> 1)
00000040  4889C2            mov rdx,rax   // rdx := 1
00000043  80C20B            add dl,0xb    // rdx += 0xb (=> 0xC)
00000046  0F05              syscall       // %rax=1 - это sys_write с параметрами в: rdi (fd), rsi (buf), rdx (len)
00000048  4831C0            xor rax,rax
0000004B  043C              add al,0x3c
0000004D  0F05              syscall       // %rax=0x3c (60) - это sys_exit
0000004F  E8AEFFFFFF        call 0x2
00000054  A98ACFCDBA        test eax,0xbacdcf8a
00000059  C7                db 0xc7
0000005A  9A                db 0x9a
0000005B  C7                db 0xc7
0000005C  91                xchg eax,ecx
0000005D  A88D              test al,0x8d
0000005F  9B7ADA            wait jpe 0x3c
00000062  F5                cmc
00000063  E36D              jrcxz 0xd2
00000065  AF                scasd
00000066  B6F9              mov dh,0xf9
00000068  4D                rex.wrb
00000069  492D6238F039      sub rax,0x39f03862
0000006F  68D34F9074        push qword 0x74904fd3
00000074  5B                pop rbx
00000075  52                push rdx
00000076  41                rex.b
00000077  4E                rex.wrx
00000078  4453              push rbx
0000007A  54                push rsp
0000007B  52                push rdx
0000007C  32                db 0x32
0000007D  5D                pop rbp
```

* [Linux syscall table](https://filippo.io/linux-syscall-table/)