// vec32/op_amd64.s
// vim: ft=9asm ai

// func Add(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +16(FP)
// Return: +32(FP)
TEXT ·Add(SB),7,$0-48
    MOVUPS      v1+0(FP), X0
    MOVUPS      v2+16(FP), X1
    ADDPS       X1, X0
    MOVUPS      X0, ret+32(FP)
    RET

// func Sub(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +16(FP)
// Return: +32(FP)
TEXT ·Sub(SB),7,$0-48
    MOVUPS      v1+0(FP), X0
    MOVUPS      v2+16(FP), X1
    SUBPS       X1, X0
    MOVUPS      X0, ret+32(FP)
    RET

// func Dot(v1, v2 Vector) float32
// v1: +0(FP)
// v2: +16(FP)
// Return: +32(FP)
TEXT ·Dot(SB),7,$0-36
    MOVUPS      v1+0(FP), X0
    MOVUPS      v2+16(FP), X1
    MULPS       X1, X0
    MOVLHPS     X0, X2
    MOVHLPS     X0, X2
    ADDPS       X2, X0
    SHUFPS      $0x03, X0, X2
    ADDSS       X2, X0
    MOVSS       X0, ret+32(FP)
    RET

// func Cross(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +16(FP)
// Return: +32(FP)
TEXT ·Cross(SB),7,$0-48
    JMP         ·cross(SB)
