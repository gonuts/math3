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
    //     0     1     2     3
    // X0: v1.y, v1.z, v1.z, v1.x
    // X1: ----, ----, v1.x, v1.y
    // X2: v2.z, v2.x, v2.y, v2.z
    // X3: ----, ----, v2.y, v2.x
    MOVUPS      v1+0(FP), X0
    MOVUPS      X0, X1
    SHUFPS      $0x29, X0, X0
    SHUFPS      $0x40, X1, X1
    MOVUPS      v2+16(FP), X2
    MOVUPS      X2, X3
    SHUFPS      $0x92, X2, X2
    SHUFPS      $0x10, X3, X3

    // perform all multiplications
    MULPS       X2, X0
    MULPS       X3, X1

    // X0: result.x, result.y, JUNK, JUNK
    MOVHLPS     X0, X2
    SUBPS       X2, X0

    // X1: 0, 0, result.z, JUNK
    SHUFPS      $0x30, X1, X3
    MOVLHPS     X1, X4
    MOVHLPS     X4, X3
    SUBPS       X3, X1

    SHUFPS      $0x20, X1, X1
    SHUFPS      $0xe4, X1, X0
    MOVUPS      X0, ret+32(FP)
    RET
