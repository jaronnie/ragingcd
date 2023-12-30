package com.jaronnie.ragingcd.stdb.common.component.exception;

import lombok.Getter;

@Getter
public class BaseException extends RuntimeException {
    private static final long serialVersionUID = 187364533553L;

    /**
     * 标准错误码
     */
    private int code;

    /**
     * 标准错误信息
     */
    private String message;

    /**
     * 异常信息枚举，所有的异常枚举类都实现了该接口
     */
    private IResponseEnum responseEnum;

    /**
     * 自定义异常消息参数
     */
    private Object[] customArgs;

    /**
     * 自定义异常消息
     */
    private String customMessage;

    public BaseException(int code, String message) {
        super(message);
        this.message = message;
        this.code = code;
    }

    /**
     * 自定义 500 系统内部错误
     *
     * @param message message
     */
    public BaseException(String message) {
        this(ResponseEnum.INTERNAL_SERVER_ERROR.getCode(), message);
    }

    public BaseException(IResponseEnum responseEnum) {
        super(responseEnum.getMessage());
        this.responseEnum = responseEnum;
        this.code = responseEnum.getCode();
        this.message = responseEnum.getMessage();
    }

    public int getCode() {
        return this.code;
    }

    public BaseException(IResponseEnum responseEnum, Object[] args, String message) {
        super(responseEnum.getMessage());
        this.responseEnum = responseEnum;
        this.code = responseEnum.getCode();
        this.message = responseEnum.getMessage();
        this.customArgs = args;
        this.customMessage = message;
    }


    public BaseException(IResponseEnum responseEnum, Object[] args, String message, Throwable cause) {
        super(responseEnum.getMessage(), cause);
        this.responseEnum = responseEnum;
        this.code = responseEnum.getCode();
        this.message = responseEnum.getMessage();
        this.customArgs = args;
        this.customMessage = message;
    }

    @Override
    public Throwable fillInStackTrace() {
        return this;
    }
}
