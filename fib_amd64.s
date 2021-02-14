#include "textflag.h"

TEXT ·Fib(SB), NOSPLIT, $0
    MOVL $0, AX

    MOVQ n+0(FP), CX
    CMPQ CX, $0
    JZ exit

    MOVL $1, BX

calc:
    MOVQ BX, DX
    ADDQ AX, BX
    MOVQ DX, AX
    LOOP calc

exit:
    MOVQ AX, ret+8(FP)
    RET
    