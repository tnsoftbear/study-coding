	.file	"overload.cpp"
	.intel_syntax noprefix
	.text
#APP
	.globl _ZSt21ios_base_library_initv
	.section	.rodata.str1.1,"aMS",@progbits,1
.LC0:
	.string	"int f(double a)"
#NO_APP
	.text
	.globl	_Z1fd
	.type	_Z1fd, @function
_Z1fd:
.LFB2003:
	.cfi_startproc
	sub	rsp, 24
	.cfi_def_cfa_offset 32
	movsd	QWORD PTR 8[rsp], xmm0
	lea	rdi, .LC0[rip]
	call	puts@PLT
	cvttsd2si	eax, QWORD PTR 8[rsp]
	add	rsp, 24
	.cfi_def_cfa_offset 8
	ret
	.cfi_endproc
.LFE2003:
	.size	_Z1fd, .-_Z1fd
	.section	.rodata.str1.1
.LC1:
	.string	"double f(int a)"
	.text
	.globl	_Z1fi
	.type	_Z1fi, @function
_Z1fi:
.LFB2004:
	.cfi_startproc
	push	rbx
	.cfi_def_cfa_offset 16
	.cfi_offset 3, -16
	mov	ebx, edi
	lea	rdi, .LC1[rip]
	call	puts@PLT
	pxor	xmm0, xmm0
	cvtsi2sd	xmm0, ebx
	pop	rbx
	.cfi_def_cfa_offset 8
	ret
	.cfi_endproc
.LFE2004:
	.size	_Z1fi, .-_Z1fi
	.section	.rodata.str1.1
.LC2:
	.string	"float f(float a)"
	.text
	.globl	_Z1ff
	.type	_Z1ff, @function
_Z1ff:
.LFB2005:
	.cfi_startproc
	sub	rsp, 24
	.cfi_def_cfa_offset 32
	movss	DWORD PTR 12[rsp], xmm0
	lea	rdi, .LC2[rip]
	call	puts@PLT
	movss	xmm0, DWORD PTR 12[rsp]
	add	rsp, 24
	.cfi_def_cfa_offset 8
	ret
	.cfi_endproc
.LFE2005:
	.size	_Z1ff, .-_Z1ff
	.section	.rodata.str1.1
.LC3:
	.string	"void f(...) / "
	.text
	.globl	_Z1fPKcz
	.type	_Z1fPKcz, @function
_Z1fPKcz:
.LFB2006:
	.cfi_startproc
	push	rbx
	.cfi_def_cfa_offset 16
	.cfi_offset 3, -16
	sub	rsp, 208
	.cfi_def_cfa_offset 224
	mov	rbx, rdi
	mov	QWORD PTR 40[rsp], rsi
	mov	QWORD PTR 48[rsp], rdx
	mov	QWORD PTR 56[rsp], rcx
	mov	QWORD PTR 64[rsp], r8
	mov	QWORD PTR 72[rsp], r9
	test	al, al
	je	.L8
	movaps	XMMWORD PTR 80[rsp], xmm0
	movaps	XMMWORD PTR 96[rsp], xmm1
	movaps	XMMWORD PTR 112[rsp], xmm2
	movaps	XMMWORD PTR 128[rsp], xmm3
	movaps	XMMWORD PTR 144[rsp], xmm4
	movaps	XMMWORD PTR 160[rsp], xmm5
	movaps	XMMWORD PTR 176[rsp], xmm6
	movaps	XMMWORD PTR 192[rsp], xmm7
.L8:
	mov	rax, QWORD PTR fs:40
	mov	QWORD PTR 24[rsp], rax
	xor	eax, eax
	lea	rdi, .LC3[rip]
	call	printf@PLT
	mov	DWORD PTR [rsp], 8
	mov	DWORD PTR 4[rsp], 48
	lea	rax, 224[rsp]
	mov	QWORD PTR 8[rsp], rax
	lea	rax, 32[rsp]
	mov	QWORD PTR 16[rsp], rax
	mov	rdx, rsp
	mov	rsi, rbx
	mov	rdi, QWORD PTR stdout[rip]
	call	vfprintf@PLT
	mov	rax, QWORD PTR 24[rsp]
	sub	rax, QWORD PTR fs:40
	jne	.L11
	add	rsp, 208
	.cfi_remember_state
	.cfi_def_cfa_offset 16
	pop	rbx
	.cfi_def_cfa_offset 8
	ret
.L11:
	.cfi_restore_state
	call	__stack_chk_fail@PLT
	.cfi_endproc
.LFE2006:
	.size	_Z1fPKcz, .-_Z1fPKcz
	.section	.rodata.str1.1
.LC4:
	.string	"void f(long)"
	.text
	.globl	_Z1fl
	.type	_Z1fl, @function
_Z1fl:
.LFB2007:
	.cfi_startproc
	sub	rsp, 8
	.cfi_def_cfa_offset 16
	lea	rdi, .LC4[rip]
	call	puts@PLT
	add	rsp, 8
	.cfi_def_cfa_offset 8
	ret
	.cfi_endproc
.LFE2007:
	.size	_Z1fl, .-_Z1fl
	.section	.rodata.str1.1
.LC5:
	.string	"void f(long&)"
	.text
	.globl	_Z1fRl
	.type	_Z1fRl, @function
_Z1fRl:
.LFB2008:
	.cfi_startproc
	sub	rsp, 8
	.cfi_def_cfa_offset 16
	lea	rdi, .LC5[rip]
	call	puts@PLT
	add	rsp, 8
	.cfi_def_cfa_offset 8
	ret
	.cfi_endproc
.LFE2008:
	.size	_Z1fRl, .-_Z1fRl
	.section	.rodata.str1.1
.LC6:
	.string	"Nums: %d %d \n"
	.text
	.globl	main
	.type	main, @function
main:
.LFB2009:
	.cfi_startproc
	sub	rsp, 8
	.cfi_def_cfa_offset 16
	lea	rdi, .LC0[rip]
	call	puts@PLT
	lea	rdi, .LC1[rip]
	call	puts@PLT
	lea	rdi, .LC2[rip]
	call	puts@PLT
	mov	edx, 4
	mov	esi, 3
	lea	rdi, .LC6[rip]
	mov	eax, 0
	call	_Z1fPKcz
	lea	rdi, .LC4[rip]
	call	puts@PLT
	mov	eax, 0
	add	rsp, 8
	.cfi_def_cfa_offset 8
	ret
	.cfi_endproc
.LFE2009:
	.size	main, .-main
	.ident	"GCC: (GNU) 13.2.1 20230801"
	.section	.note.GNU-stack,"",@progbits
