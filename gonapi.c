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
    napi_create_string_utf8(env, cd->Type, NAPI_AUTO_LENGTH, &params[0]);
    napi_create_string_utf8(env, cd->Content, NAPI_AUTO_LENGTH, &params[1]);
    napi_call_function(env, NULL, callback, 2, params, NULL);
}
napi_threadsafe_function_call_js ThreadsafeFunctionCallback() { return tsfn_callback_function; }
napi_value *GetParams(napi_env env, napi_value fn) {
    napi_value dd[1];
    napi_create_int32(env, 222, &dd[0]);
    return &dd;
}