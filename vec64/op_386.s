// vec64/op_386.s
// vim: ft=9asm ai

// func Add(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +24(FP)
// Return: +48(FP)
TEXT ·Add(SB),7,$0
    JMP         ·add(SB)

// func Sub(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +24(FP)
// Return: +48(FP)
TEXT ·Sub(SB),7,$0
    JMP         ·sub(SB)

// func Dot(v1, v2 Vector) float64
// v1: +0(FP)
// v2: +24(FP)
// Return: +48(FP)
TEXT ·Dot(SB),7,$0
    JMP         ·dot(SB)

// func Cross(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +24(FP)
// Return: +48(FP)
TEXT ·Cross(SB),7,$0
    JMP         ·cross(SB)
