// vec64/op_amd64.s
// vim: ft=9asm ai

// func Add(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +32(FP)
// Return: +64(FP)
TEXT ·Add(SB),7,$0-96
    MOVUPD      v1+0(FP), X0
    MOVUPD      v2+32(FP), X2
    ADDPD       X2, X0
    MOVUPD      X0, ret+64(FP)
    MOVUPD      v1ZW+16(FP), X0
    MOVUPD      v2ZW+48(FP), X2
    ADDPD       X2, X0
    MOVUPD      X0, retZW+80(FP)
    RET

// func Sub(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +32(FP)
// Return: +64(FP)
TEXT ·Sub(SB),7,$0-96
    MOVUPD      v1+0(FP), X0
    MOVUPD      v2+32(FP), X2
    SUBPD       X2, X0
    MOVUPD      X0, ret+64(FP)
    MOVUPD      v1ZW+16(FP), X0
    MOVUPD      v2ZW+48(FP), X2
    SUBPD       X2, X0
    MOVUPD      X0, retZW+80(FP)
    RET

// func Dot(v1, v2 Vector) float64
// v1: +0(FP)
// v2: +32(FP)
// Return: +64(FP)
TEXT ·Dot(SB),7,$0-72
    MOVUPD      v1XY+0(FP), X0
    MOVUPD      v2XY+32(FP), X2
    MULPD       X2, X0
    MOVUPD      X0, X2
    SHUFPD      $1, X2, X2
    ADDSD       X2, X0

    MOVUPD      v1ZW+16(FP), X1
    MOVUPD      v2ZW+48(FP), X3
    MULPD       X3, X1
    MOVUPD      X1, X3
    SHUFPD      $1, X3, X3
    ADDSD       X3, X1

    ADDSD       X1, X0
    MOVSD       X0, ret+64(FP)
    RET

// func Cross(v1, v2 Vector) Vector
// v1: +0(FP)
// v2: +32(FP)
// Return: +64(FP)
TEXT ·Cross(SB),7,$0-96
    // X
    MOVUPD      v1YZ+8(FP), X0
    SHUFPD      $1, X0, X0
    MOVUPD      v2YZ+40(FP), X1
    MULPD       X1, X0
    SHUFPD      $0, X0, X1
    SUBPD       X1, X0
    MOVHPD      X0, retX+64(FP)
    // Y
    MOVHPD      v1Z+24(FP), X0
    MOVLPD      v1X+0(FP), X0
    MOVHPD      v2X+32(FP), X1
    MOVLPD      v2Z+48(FP), X1
    MULPD       X1, X0
    SHUFPD      $0, X0, X1
    SUBPD       X1, X0
    MOVHPD      X0, retY+72(FP)
    // Z
    MOVUPD      v1XY+0(FP), X0
    SHUFPD      $1, X0, X0
    MOVUPD      v2XY+32(FP), X1
    MULPD       X1, X0
    SHUFPD      $0, X0, X1
    SUBPD       X1, X0
    MOVHPD      X0, retZ+80(FP)
    // W
    MOVSD       $0.0, X2
    MOVSD       X2, retW+88(FP)
    // Done!
    RET
