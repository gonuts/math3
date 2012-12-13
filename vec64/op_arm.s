// vec64/op_arm.s
// vim: ft=9asm ai

// func Add(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +32(FP)
// Return: +64(FP)
TEXT ·Add(SB),7,$0
    B           ·add(SB)

// func Sub(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +32(FP)
// Return: +64(FP)
TEXT ·Sub(SB),7,$0
    B           ·sub(SB)

// func Dot(v1, v2 Vector) float64
// v1: +0(FP)
// v2: +32(FP)
// Return: +64(FP)
TEXT ·Dot(SB),7,$0
    B           ·dot(SB)

// func Cross(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +32(FP)
// Return: +64(FP)
TEXT ·Cross(SB),7,$0
    B           ·cross(SB)
