package com.jaronnie.ragingcd.stdb.common.component.exception;

public class CommonException extends BaseException {
    private static final long serialVersionUID = 376492087463L;

    public CommonException(IResponseEnum responseEnum, Object[] args, String message) {
        super(responseEnum, args, message);
    }

    // message 和 args 配合使用，MessageFormat.format(message,args)
    public CommonException(IResponseEnum responseEnum, Object[] args, String message, Throwable cause) {
        super(responseEnum, args, message, cause);
    }

    public CommonException(int code, String msg) {
        super(code, msg);
    }
}
