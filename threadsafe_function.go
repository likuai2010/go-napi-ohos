package napi

import (
	"unsafe"
)

type ThreadsafeFunction unsafe.Pointer

type TsfnCallbackFunc func(Env, Value, unsafe.Pointer, unsafe.Pointer)

type ThreadsafeFunctionsCaller struct {
	Cb TsfnCallbackFunc
}
