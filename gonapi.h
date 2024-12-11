//
// Created on 2024/12/10.
//
// Node APIs are not fully supported. To solve the compilation error of the interface cannot be found,
// please include "napi/native_api.h".
#include <napi/native_api.h>

typedef struct {
    napi_value Key;
    napi_value Value;
} CallbackData;

napi_threadsafe_function_call_js ThreadsafeFunctionCallback(void);

void CallThreadsafeFunctionCallback(void *caller, napi_env env, napi_value callback, void *ctx, void *data);