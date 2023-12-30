package com.jaronnie.ragingcd.stdb.common.component.exception;

import com.jaronnie.ragingcd.stdb.common.component.exception.assertion.CommonExceptionAssert;
import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public enum ResponseEnum implements CommonExceptionAssert {
    SUCCESS(0, "success"),
    BODY_NOT_MATCH(400, "请求数据格式不匹配"),
    NO_AUTHENTICATION(401, "未认证"),
    NOT_FOUND(404, "未找到该资源"),
    INTERNAL_SERVER_ERROR(500, "服务器内部错误!"),
    SERVER_BUSY(503, "服务器正忙,请稍后再试!"),
    SERVER_ERROR(504, "网络异常"),
    VALID_ERROR(507, "参数校验异常"),
    NullPointer_ERROR(508, "空指针异常"),
    Method_Argument_Type_Mismatch_ERROR(509, "方法参数类型不匹配");

    private int code;
    private String message;

    @Override
    public int getCode() {
        return code;
    }

    @Override
    public String getMessage() {
        return message;
    }

}
