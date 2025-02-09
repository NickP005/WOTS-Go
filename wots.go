package wots

/*
#cgo CFLAGS: -I${SRCDIR}
#include "wots.h"
*/
import "C"
import "unsafe"

func wotsPkgen(seed [32]byte, pubSeed [32]byte, addr [32]byte) [2144]byte {
	var pk [2144]byte
	C.wots_pkgen(
		(*C.word8)(unsafe.Pointer(&pk[0])),
		(*C.word8)(unsafe.Pointer(&seed[0])),
		(*C.word8)(unsafe.Pointer(&pubSeed[0])),
		(*C.word32)(unsafe.Pointer(&addr[0])),
	)
	return pk
}

func wotsSign(msg [32]byte, seed [32]byte, pubSeed [32]byte, addr [32]byte) [2144]byte {
	var sig [2144]byte
	C.wots_sign(
		(*C.word8)(unsafe.Pointer(&sig[0])),
		(*C.word8)(unsafe.Pointer(&msg[0])),
		(*C.word8)(unsafe.Pointer(&seed[0])),
		(*C.word8)(unsafe.Pointer(&pubSeed[0])),
		(*C.word32)(unsafe.Pointer(&addr[0])),
	)
	return sig
}

func wotsPkFromSig(sig [2144]byte, msg [32]byte, pubSeed [32]byte, addr [32]byte) [2144]byte {
	var pk [2144]byte
	C.wots_pk_from_sig(
		(*C.word8)(unsafe.Pointer(&pk[0])),
		(*C.word8)(unsafe.Pointer(&sig[0])),
		(*C.word8)(unsafe.Pointer(&msg[0])),
		(*C.word8)(unsafe.Pointer(&pubSeed[0])),
		(*C.word32)(unsafe.Pointer(&addr[0])),
	)
	return pk
}
