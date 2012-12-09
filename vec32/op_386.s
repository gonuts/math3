// vec32/op_386.s
// vim: ft=9asm ai

// func Add(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +16(FP)
// Return: +32(FP)
TEXT ·Add(SB),7,$0-48
    JMP         ·add(SB)

// func Sub(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +16(FP)
// Return: +32(FP)
TEXT ·Sub(SB),7,$0-48
    JMP         ·sub(SB)

// func Dot(v1, v2 Vector) float32
// v1: +0(FP)
// v2: +16(FP)
// Return: +32(FP)
TEXT ·Dot(SB),7,$0-36
    JMP         ·dot(SB)

// func Cross(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +16(FP)
// Return: +32(FP)
TEXT ·Cross(SB),7,$0-48
    JMP         ·cross(SB)
