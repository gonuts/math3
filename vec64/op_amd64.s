// vec64/op_amd64.s
// vim: ft=9asm ai

// func Add(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +24(FP)
// Return: +48(FP)
TEXT ·Add(SB),$0-72
    MOVUPD      v1+0(FP), X0
    MOVUPD      v2+24(FP), X2
    ADDPD       X2, X0
    MOVUPD      X0, ret+48(FP)
    MOVSD       v1Z+16(FP), X0
    MOVSD       v2Z+40(FP), X2
    ADDSD       X2, X0
    MOVSD       X0, retZ+64(FP)
    RET

// func Sub(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +24(FP)
// Return: +48(FP)
TEXT ·Sub(SB),$0-72
    MOVUPD      v1+0(FP), X0
    MOVUPD      v2+24(FP), X2
    SUBPD       X2, X0
    MOVUPD      X0, ret+48(FP)
    MOVSD       v1Z+16(FP), X0
    MOVSD       v2Z+40(FP), X2
    SUBSD       X2, X0
    MOVSD       X0, retZ+64(FP)
    RET

// func Dot(v1, v2 Vector) float64
// v1: +0(FP)
// v2: +24(FP)
// Return: +48(FP)
TEXT ·Dot(SB),$0-56
    MOVUPD      v1+0(FP), X0
    MOVUPD      v2+24(FP), X2
    MULPD       X2, X0
    MOVUPD      X0, X2
    SHUFPD      $1, X2, X2
    MOVSD       v1Z+16(FP), X1
    MOVSD       v2Z+40(FP), X3
    MULSD       X3, X1
    ADDSD       X2, X0
    ADDSD       X1, X0
    MOVSD       X0, ret+48(FP)
    RET

// func Cross(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +24(FP)
// Return: +48(FP)
TEXT ·Cross(SB),$0-72
    // X
    MOVUPD      v1YZ+8(FP), X0
    SHUFPD      $1, X0, X0
    MOVUPD      v2YZ+32(FP), X1
    MULPD       X1, X0
    SHUFPD      $0, X0, X1
    SUBPD       X1, X0
    MOVHPD      X0, retX+48(FP)
    // Y
    MOVHPD      v1Z+16(FP), X0
    MOVLPD      v1X+0(FP), X0
    MOVHPD      v2X+24(FP), X1
    MOVLPD      v2Z+40(FP), X1
    MULPD       X1, X0
    SHUFPD      $0, X0, X1
    SUBPD       X1, X0
    MOVHPD      X0, retY+56(FP)
    // Z
    MOVUPD      v1XY+0(FP), X0
    SHUFPD      $1, X0, X0
    MOVUPD      v2XY+24(FP), X1
    MULPD       X1, X0
    SHUFPD      $0, X0, X1
    SUBPD       X1, X0
    MOVHPD      X0, retZ+64(FP)
    // Done!
    RET
