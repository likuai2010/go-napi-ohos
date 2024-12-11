package napi

/*

#cgo CFLAGS: -I/usr/include/napi
#include "gonapi.h"
#include <napi/native_api.h>
*/
import "C"

import (
	"unsafe"
)

func CreateAsyncWork(
	env Env,
	asyncResource, asyncResourceName Value,
	execute AsyncExecuteCallback,
	complete AsyncCompleteCallback,
) (AsyncWork, Status) {
	provider, status := getInstanceData(env)
	if status != StatusOK || provider == nil {
		return AsyncWork{}, status
	}

	return provider.GetAsyncWorkData().CreateAsyncWork(
		env,
		asyncResource, asyncResourceName,
		execute,
		complete,
	)
}

func DeleteAsyncWork(env Env, work AsyncWork) Status {
	provider, status := getInstanceData(env)
	if status != StatusOK || provider == nil {
		return status
	}

	defer provider.GetAsyncWorkData().DeleteAsyncWork(work.ID)
	return Status(C.napi_delete_async_work(
		C.napi_env(env),
		C.napi_async_work(work.Handle),
	))
}

func QueueAsyncWork(env Env, work AsyncWork) Status {
	return Status(C.napi_queue_async_work(
		C.napi_env(env),
		C.napi_async_work(work.Handle),
	))
}

func CancelAsyncWork(env Env, work AsyncWork) Status {
	return Status(C.napi_cancel_async_work(
		C.napi_env(env),
		C.napi_async_work(work.Handle),
	))
}

func GetNodeVersion(env Env) (NodeVersion, Status) {
	var cresult *C.napi_node_version
	status := Status(C.napi_get_node_version(
		C.napi_env(env),
		(**C.napi_node_version)(&cresult),
	))

	if status != StatusOK {
		return NodeVersion{}, status
	}

	return NodeVersion{
		Major:   uint(cresult.major),
		Minor:   uint(cresult.minor),
		Patch:   uint(cresult.patch),
		Release: C.GoString(cresult.release),
	}, status
}

func GetModuleFileName(env Env) (string, Status) {
	var cresult *C.char
	status := Status(C.node_api_get_module_file_name(
		C.napi_env(env),
		(**C.char)(&cresult),
	))

	if status != StatusOK {
		return "", status
	}

	return C.GoString(cresult), status
}

func CreateThreadsafeFunction(
	env Env,
	fn Value,
	asyncResource, asyncResourceName Value,
	maxQueueSize, initialThreadCount int,
	callback bool,
) (ThreadsafeFunction, Status) {
	var result ThreadsafeFunction
	status := Status(C.napi_create_threadsafe_function(
		C.napi_env(env),
		C.napi_value(fn),
		C.napi_value(asyncResource),
		C.napi_value(asyncResourceName),
		C.size_t(maxQueueSize),
		C.size_t(initialThreadCount),
		nil,
		nil,
		nil,
		hasCallback(callback),
		(*C.napi_threadsafe_function)(unsafe.Pointer(&result)),
	))
	return result, status
}
func hasCallback(callback bool) C.napi_threadsafe_function_call_js {
	if callback {
		return C.ThreadsafeFunctionCallback()
	} else {
		return nil
	}
}

//export CallThreadsafeFunctionCallback
func CallThreadsafeFunctionCallback(wrap unsafe.Pointer, env C.napi_env, fn C.napi_value, ctx unsafe.Pointer, data unsafe.Pointer) {
	caller := (*ThreadsafeFunctionsCaller)(wrap)
	caller.Cb(Env(env), Value(fn), ctx, data)
}

func CallThreadsafeFunction(
	fn ThreadsafeFunction,
	key, value Value,
) Status {
	params := C.CallbackData{
		Key:   C.napi_value(key),
		Value: C.napi_value(value),
	}
	return Status(C.napi_call_threadsafe_function(
		C.napi_threadsafe_function(fn),
		unsafe.Pointer(&params),
		C.napi_tsfn_blocking,
	))
}

func AcquireThreadsafeFunction(fn ThreadsafeFunction) Status {
	return Status(C.napi_acquire_threadsafe_function(
		C.napi_threadsafe_function(fn),
	))
}

func ReleaseThreadsafeFunction(
	fn ThreadsafeFunction,
) Status {
	return Status(C.napi_release_threadsafe_function(
		C.napi_threadsafe_function(fn),
		C.napi_tsfn_release,
	))
}
