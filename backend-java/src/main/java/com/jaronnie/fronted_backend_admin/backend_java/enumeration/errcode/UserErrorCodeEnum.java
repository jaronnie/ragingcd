package com.jaronnie.fronted_backend_admin.backend_java.enumeration.errcode;

import com.jaronnie.fronted_backend_admin.backend_java.component.exception.assertion.BusinessExceptionAssert;
import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public enum UserErrorCodeEnum implements BusinessExceptionAssert {
    LoginError(10001, "用户名或密码错误"),
    LogUpError(10002, "用户名重名");

    private final int code;
    private final String message;
}
