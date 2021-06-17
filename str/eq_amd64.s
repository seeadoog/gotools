

TEXT ·equal(SB),$0-$40
    MOVQ s1+8(FP),CX
    MOVQ s2+24(FP),BX
    CMPQ CX,BX
    JEQ conn
    MOVQ $0 ret+32(FP)
    RET
 conn:

 retF:
