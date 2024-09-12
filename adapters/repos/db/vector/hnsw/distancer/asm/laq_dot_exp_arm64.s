//go:build !noasm && arm64
// AUTO-GENERATED BY GOAT -- DO NOT EDIT

TEXT ·laq_dot_exp_neon(SB), $0-32
	MOVD x+0(FP), R0
	MOVD y1+8(FP), R1
	MOVD y2+16(FP), R2
	MOVD a1+24(FP), R3
	MOVD a2+32(FP), R4
	MOVD res+40(FP), R5
	MOVD len+48(FP), R6
	WORD $0xa9bf7bfd    // stp	x29, x30, [sp, #-16]!
	WORD $0xb94000c8    // ldr	w8, [x6]
	WORD $0x910003fd    // mov	x29, sp
	WORD $0x6f00e400    // movi	v0.2d, #0000000000000000
	WORD $0x7100411f    // cmp	w8, #16
	WORD $0x540005ab    // b.lt	.LBB0_3
	WORD $0x4d40c861    // ld1r	{ v1.4s }, [x3]
	WORD $0x11004108    // add	w8, w8, #16
	WORD $0x6f00e400    // movi	v0.2d, #0000000000000000
	WORD $0xbd400082    // ldr	s2, [x4]

LBB0_2:
	WORD $0x3cc10423 // ldr	q3, [x1], #16
	WORD $0x3cc10444 // ldr	q4, [x2], #16
	WORD $0xad404c12 // ldp	q18, q19, [x0]
	WORD $0x2f08a465 // ushll	v5.8h, v3.8b, #0
	WORD $0x51004108 // sub	w8, w8, #16
	WORD $0x2f08a486 // ushll	v6.8h, v4.8b, #0
	WORD $0x71007d1f // cmp	w8, #31
	WORD $0x2f10a4a7 // ushll	v7.4s, v5.4h, #0
	WORD $0x6f10a4a5 // ushll2	v5.4s, v5.8h, #0
	WORD $0x6e21d8e7 // ucvtf	v7.4s, v7.4s
	WORD $0x2f10a4d1 // ushll	v17.4s, v6.4h, #0
	WORD $0x6e21d8a5 // ucvtf	v5.4s, v5.4s
	WORD $0x6f08a463 // ushll2	v3.8h, v3.16b, #0
	WORD $0x6e21da31 // ucvtf	v17.4s, v17.4s
	WORD $0x6e27dc27 // fmul	v7.4s, v1.4s, v7.4s
	WORD $0x2f10a470 // ushll	v16.4s, v3.4h, #0
	WORD $0x6f10a4c6 // ushll2	v6.4s, v6.8h, #0
	WORD $0x6e21da10 // ucvtf	v16.4s, v16.4s
	WORD $0x6e25dc25 // fmul	v5.4s, v1.4s, v5.4s
	WORD $0x6e21d8c6 // ucvtf	v6.4s, v6.4s
	WORD $0x6f08a484 // ushll2	v4.8h, v4.16b, #0
	WORD $0x4f821227 // fmla	v7.4s, v17.4s, v2.s[0]
	WORD $0x6f10a463 // ushll2	v3.4s, v3.8h, #0
	WORD $0x2f10a491 // ushll	v17.4s, v4.4h, #0
	WORD $0x6e21d863 // ucvtf	v3.4s, v3.4s
	WORD $0x4f8210c5 // fmla	v5.4s, v6.4s, v2.s[0]
	WORD $0x6e30dc30 // fmul	v16.4s, v1.4s, v16.4s
	WORD $0x6e21da31 // ucvtf	v17.4s, v17.4s
	WORD $0x4e32cce0 // fmla	v0.4s, v7.4s, v18.4s
	WORD $0x6f10a484 // ushll2	v4.4s, v4.8h, #0
	WORD $0x6e23dc23 // fmul	v3.4s, v1.4s, v3.4s
	WORD $0x6e21d884 // ucvtf	v4.4s, v4.4s
	WORD $0x4f821230 // fmla	v16.4s, v17.4s, v2.s[0]
	WORD $0x4e33cca0 // fmla	v0.4s, v5.4s, v19.4s
	WORD $0xad411805 // ldp	q5, q6, [x0, #32]
	WORD $0x91010000 // add	x0, x0, #64
	WORD $0x4f821083 // fmla	v3.4s, v4.4s, v2.s[0]
	WORD $0x4e25ce00 // fmla	v0.4s, v16.4s, v5.4s
	WORD $0x4e26cc60 // fmla	v0.4s, v3.4s, v6.4s
	WORD $0x54fffb28 // b.hi	.LBB0_2

LBB0_3:
	WORD $0x4e0c0401 // dup	v1.4s, v0.s[1]
	WORD $0x4e140402 // dup	v2.4s, v0.s[2]
	WORD $0x4e21d401 // fadd	v1.4s, v0.4s, v1.4s
	WORD $0x4e1c0400 // dup	v0.4s, v0.s[3]
	WORD $0x4e21d441 // fadd	v1.4s, v2.4s, v1.4s
	WORD $0x4e21d400 // fadd	v0.4s, v0.4s, v1.4s
	WORD $0x2f00e401 // movi	d1, #0000000000000000
	WORD $0x1e212800 // fadd	s0, s0, s1
	WORD $0xbd0000a0 // str	s0, [x5]
	WORD $0xa8c17bfd // ldp	x29, x30, [sp], #16
	WORD $0xd65f03c0 // ret