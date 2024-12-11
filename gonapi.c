//
// Created on 2024/12/10.
//
// Node APIs are not fully supported. To solve the compilation error of the interface cannot be found,
// please include "napi/native_api.h".

#include "gonapi.h"
#include <napi/native_api.h>


void tsfn_callback_function(napi_env env, napi_value callback, void *ctx, void *data) {
    CallbackData *cd = (CallbackData *)data;
    napi_value params[2];
    params[0] = cd->Key;
    params[1] = cd->Value;
    napi_call_function(env, NULL, callback, 2, params, NULL);
}
napi_threadsafe_function_call_js ThreadsafeFunctionCallback() { return tsfn_callback_function; }
