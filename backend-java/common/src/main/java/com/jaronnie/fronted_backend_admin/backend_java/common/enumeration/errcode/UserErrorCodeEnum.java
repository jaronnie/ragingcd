package com.jaronnie.fronted_backend_admin.backend_java.common.enumeration.errcode;

import com.jaronnie.fronted_backend_admin.backend_java.common.component.exception.assertion.BusinessExceptionAssert;
import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public enum UserErrorCodeEnum implements BusinessExceptionAssert {
    LoginError(10001, "用户名或密码错误"),
    DuplicateUsernameError(10002, "用户名重名"),
    RegisterEmailVerificationCodeError(10003, "邮箱验证码错误或已过期"),
    RegisterError(10003, "注册失败");

    private final int code;
    private final String message;
}
