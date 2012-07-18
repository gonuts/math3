// vec64/op_arm.s
// vim: ft=9asm ai

// func Add(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +24(FP)
// Return: +48(FP)
TEXT ·Add(SB),$0-72
    B           ·add(SB)

// func Sub(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +24(FP)
// Return: +48(FP)
TEXT ·Sub(SB),$0-72
    B           ·sub(SB)

// func Dot(v1, v2 Vector) float64
// v1: +0(FP)
// v2: +24(FP)
// Return: +48(FP)
TEXT ·Dot(SB),$0-56
    B           ·dot(SB)

// func Cross(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +24(FP)
// Return: +48(FP)
TEXT ·Cross(SB),$0-72
    B           ·cross(SB)
