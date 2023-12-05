package com.jaronnie.fronted_backend_admin.backend_java.component.exception.assertion;

/**
 * 只包装了 错误信息 的 {@link RuntimeException}.
 * 用于 {@link Assert} 中用于包装自定义异常信息
 */
public class WrapMessageException extends RuntimeException {
    public WrapMessageException(String message) {
        super(message);
    }

    public WrapMessageException(String message, Throwable cause) {
        super(message, cause);
    }
}
