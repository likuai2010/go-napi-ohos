package js

import (
	"unsafe"

	"github.com/akshayganeshen/napi-go"
)

type Func struct {
	Value
}
type TsFunc struct {
	Env    napi.Env
	Value  napi.ThreadsafeFunction
	Status napi.Status
}

func (e Env) CreateThreadsafeFunction(fn Value, resourceName string) TsFunc {
	asyncResourceName := e.ValueOf(resourceName)
	caller := napi.ThreadsafeFunctionsCaller{
		Cb: func(env napi.Env, jsCallbck napi.Value, ctx unsafe.Pointer, data unsafe.Pointer) {
			napi.CallFunction(env, jsCallbck, nil)
		},
	}
	tsfn, stuats := napi.CreateThreadsafeFunction(e.Env, fn.Value, nil, asyncResourceName.Value, 0, 1, &caller)
	return TsFunc{
		Env:    e.Env,
		Value:  tsfn,
		Status: stuats,
	}
}
func (t TsFunc) Call(data unsafe.Pointer) {
	napi.CallThreadsafeFunction(t.Value, t.Env)
}
