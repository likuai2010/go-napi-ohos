package js

import (
	"github.com/akshayganeshen/napi-go"
)

type Func struct {
	Value
}
type TsFunc struct {
	Env    Env
	Value  napi.ThreadsafeFunction
	Status napi.Status
}

func (e Env) CreateThreadsafeFunction(fn Value, resourceName string) TsFunc {
	asyncResourceName := e.ValueOf(resourceName)
	tsfn, stuats := napi.CreateThreadsafeFunction(e.Env, fn.Value, nil, asyncResourceName.Value, 0, 1, true)
	return TsFunc{
		Env:    e,
		Value:  tsfn,
		Status: stuats,
	}
}
func (t TsFunc) Call(key, value Value) {
	napi.CallThreadsafeFunction(t.Value, key.Value, value.Value)
}
