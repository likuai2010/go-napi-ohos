package js

import (
	"fmt"
	"unsafe"

	"github.com/akshayganeshen/napi-go"
)

type Func struct {
	Value
}
type TsFunc struct {
	Value  napi.ThreadsafeFunction
	Status napi.Status
}

func (e Env) CreateThreadsafeFunction(fn Value, resourceName string) TsFunc {
	asyncResourceName := e.ValueOf(resourceName)
	caller := napi.ThreadsafeFunctionsCaller{
		Cb: func(env napi.Env, jsCallbck napi.Value, ctx unsafe.Pointer, data unsafe.Pointer) {
			params := make([]napi.Value, 1)
			params[0] = e.ValueOf(1).Value
			fmt.Println("xxxx ThreadsafeFunctionsCaller")
			napi.CallFunction(env, jsCallbck, params)
		},
	}
	tsfn, stuats := napi.CreateThreadsafeFunction(e.Env, fn.Value, nil, asyncResourceName.Value, 0, 1, &caller)
	return TsFunc{
		Value:  tsfn,
		Status: stuats,
	}
}
func (t TsFunc) Call(value Value) {
	params := make([]napi.Value, 1)
	params[0] = value.Value
	napi.CallThreadsafeFunction(t.Value, unsafe.Pointer(&params[0]))
}
